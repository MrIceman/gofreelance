package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"gofreelance/internals/index"
	"gofreelance/model"
	"os"
	"path"
	"slices"
)

var (
	readRecordSet *model.RecordDataSet
)

type Storage struct {
}

func NewDefaultEntriesHandler() *Storage {
	return &Storage{}
}

func (d Storage) Create(record model.Record) (int, error) {
	records, err := d.GetRecordSet()
	if err != nil {
		return 0, err
	}
	if records.Entries == nil {
		return 0, errors.New("records list is nil")
	}
	dateIndex := index.Current()
	record.Index = dateIndex
	currentMonthRecords := records.Entries[dateIndex]
	if currentMonthRecords == nil || len(currentMonthRecords) == 0 {
		currentMonthRecords = make([]*model.Record, 0)
	}
	id := d.getAvailabeRandomID(currentMonthRecords)
	record.ID = id
	currentMonthRecords = append(currentMonthRecords, &record)
	records.Entries[dateIndex] = currentMonthRecords
	records.CurrentRecordID = id

	err = d.overwriteRecords(records)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d Storage) getAvailabeRandomID(set []*model.Record) int {
	id := getRandID()
	contains := slices.ContainsFunc(set, func(record *model.Record) bool {
		return record.ID == id
	})
	if contains {
		return d.getAvailabeRandomID(set)
	}

	return id
}

func (d Storage) Update(record model.Record) error {
	records, err := d.GetRecordSet()
	if err != nil {
		return fmt.Errorf("failed to get current records: %s", err.Error())
	}
	entries := records.Entries[record.Index]
	updated := false
	for i := range entries {
		item := entries[i]
		if item.ID == record.ID {
			updated = true
			entries[i] = &record
			break
		}
	}
	if !updated {
		return fmt.Errorf("record %d could not be updated because it does not exist", record.ID)
	}
	if record.Ended != nil && records.CurrentRecordID == record.ID {
		// current record ended
		records.CurrentRecordID = 0
	}

	return d.overwriteRecords(records)
}

func (d Storage) Get(id int) (*model.Record, error) {
	records, err := d.GetRecordSet()
	if err != nil {
		return nil, err
	}

	return d.getByID(records.Entries, id), nil
}

func (d Storage) GetCurrentRecord() (*model.Record, error) {
	records, err := d.GetRecordSet()
	if err != nil {
		return nil, err
	}

	return d.getByID(records.Entries, records.CurrentRecordID), nil
}

func (d Storage) getByID(data map[string][]*model.Record, id int) *model.Record {
	for _, entries := range data {
		for idx := range entries {
			r := entries[idx]
			if r.ID == id {
				return r
			}
		}
	}

	return nil
}

func (d Storage) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (d Storage) GetRecordSet() (*model.RecordDataSet, error) {
	if readRecordSet != nil {
		return readRecordSet, nil
	}
	file, err := os.Open(path.Join("data", "records.json"))
	if err != nil {
		return nil, fmt.Errorf("could not open records file: %w", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	records := &model.RecordDataSet{}
	err = json.NewDecoder(file).Decode(records)
	if err != nil {
		return nil, fmt.Errorf("could not decode records file: %w", err)
	}
	readRecordSet = records
	return records, nil
}

func (d Storage) OverrideRecordSet(r *model.RecordDataSet) error {
	return d.overwriteRecords(r)
}

func (d Storage) overwriteRecords(set *model.RecordDataSet) error {
	file, err := os.OpenFile(path.Join("data", "records.json"), os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	return json.NewEncoder(file).Encode(set)
}
