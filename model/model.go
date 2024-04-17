package model

import "time"

type Record struct {
	ID       int        `json:"ID,omitempty"`
	Commited bool       `json:"commited,omitempty"`
	Started  *time.Time `json:"started,omitempty"`
	Ended    *time.Time `json:"ended,omitempty"`
	Tasks    []Task     `json:"descriptions,omitempty"`
	Index    string     `json:"index"`
}

type Task struct {
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}

type RecordDataSet struct {
	// Index is month+year in string format
	Entries         map[string][]*Record `json:"_list"`
	CurrentRecordID int                  `json:"_currentRecordID"`
}
