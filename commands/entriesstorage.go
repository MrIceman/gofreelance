package commands

import "gofreelance/model"

type EntriesStorage interface {
	Create(record model.Record) (int, error)
	Update(record model.Record) error
	Get(id int) (*model.Record, error)
	Delete(id int) error
	GetCurrentRecord() (*model.Record, error)
	GetRecordSet() (*model.RecordDataSet, error)

	OverrideRecordSet(r *model.RecordDataSet) error
}
