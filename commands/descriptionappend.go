package commands

import (
	"gofreelance/model"
	"time"
)

func AppendDescription(storage EntriesStorage, description string) error {
	record, err := storage.GetCurrentRecord()
	if err != nil {
		return err
	}
	if record == nil {
		return ErrNoActiveRecord
	}
	record.Descriptions = append(record.Descriptions, model.RecordDescriptionEntry{
		Text: description,
		Time: time.Now(),
	})
	return storage.Update(*record)
}
