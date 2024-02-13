package commands

import (
	"errors"
	"fmt"
	"gofreelance/model"
)

func Stop(storage EntriesStorage) (*model.Record, error) {
	s, err := storage.GetCurrentRecord()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch current record: %s", err.Error())
	}
	if s == nil {
		return nil, errors.New("there is no current active record")
	}
	s.Ended = nowWithPtr()

	err = storage.Update(*s)
	if err != nil {
		return nil, fmt.Errorf("failed to update record: %s", err.Error())
	}

	return s, nil
}
