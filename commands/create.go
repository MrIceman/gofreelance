package commands

import (
	"gofreelance/model"
)

func Create(storage EntriesStorage, record model.Record) (int, error) {
	r, err := storage.GetCurrentRecord()
	if err != nil {
		return 0, err
	}
	if r != nil {
		return 0, ErrRecordAlreadyStarted
	}

	return storage.Create(record)
}
