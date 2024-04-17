package commands

import (
	"gofreelance/model"
	"time"
)

func AppendTask(storage EntriesStorage, description string) error {
	record, err := storage.GetCurrentRecord()
	if err != nil {
		return err
	}
	if record == nil {
		return ErrNoActiveRecord
	}
	record.Tasks = append(record.Tasks, model.Task{
		Text: description,
		Time: time.Now(),
	})
	return storage.Update(*record)
}
