package commands

import (
	"errors"
	"fmt"
	"gofreelance/internals/index"
	"gofreelance/model"
)

func Continue(storage EntriesStorage) (*model.Record, error) {
	rs, err := storage.GetRecordSet()
	if err != nil {
		return nil, fmt.Errorf("failed to get current record: %s", err.Error())
	}
	if rs.CurrentRecordID != 0 {
		return nil, errors.New("there is already an active record")
	}
	idx := index.Current()
	records := rs.Entries[idx]
	if len(records) == 0 {
		return nil, errors.New("no record exists for the current month")
	}

	lastEntry := records[len(records)-1]
	rs.CurrentRecordID = lastEntry.ID
	lastEntry.Ended = nil

	if err := storage.OverrideRecordSet(rs); err != nil {
		return nil, err
	}

	return lastEntry, nil

}
