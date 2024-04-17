package main

import (
	"errors"
	"fmt"
	"gofreelance/commands"
	"gofreelance/data"
	"gofreelance/export"
	"gofreelance/model"
	"log"
	"strings"
	"time"
)

type Handler struct {
	storage *data.Storage
}

func NewHandler(storage *data.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Handle(arg ...string) error {
	cmd := arg[0]
	if cmd == "help" {
		log.Println(`
## Here's how to use Clitt.
## Available commands for records:
- start: start a new record (if there is no active record)
- stop: stop the current record
- continue: set the last record as active again
## Available commands for tasks:
- task add <value>: add a task to the current record
## Available commands for export:
- export: export the records to an excel file
		


	`)
		return nil

	}
	if cmd == "start" {
		log.Println("starting record")
		id, err := h.create()
		if err != nil {
			return err
		}
		log.Printf("record started with id %d\n", id)
		return nil
	}
	if cmd == "stop" {
		log.Println("stopping record")
		return h.stop()
	}
	if cmd == "continue" {
		log.Println("setting last record to active again")
		return h.continueLastRecord()
	}
	if cmd == "task" {
		if len(arg) < 2 {
			return errors.New("missing action argument, see task help")
		}
		action := arg[1]
		if ok := commands.ValidateDescriptionAction(action); !ok {
			return fmt.Errorf("%s cmd is not a valid task action", action)
		}
		if action == commands.TaskAdd {
			v := strings.Join(arg[2:], " ")
			if v == "" {
				return errors.New("missing task - the syntax is gof task add <value>")
			}
			err := commands.AppendTask(h.storage, v)
			if err != nil {
				return err
			}
			log.Println("description was added")
		}
		return nil
	}
	if cmd == "export" {
		rs, err := h.storage.GetRecordSet()
		if err != nil {
			return err
		}
		return export.SaveAsExcel(rs)
	}
	return fmt.Errorf("%cmd: %cmd", commands.ErrUnknownCommand, cmd)

}

func (h *Handler) create() (int, error) {
	now := time.Now()
	record := model.Record{
		Commited: false,
		Started:  &now,
	}
	return commands.Create(h.storage, record)
}

func (h *Handler) stop() error {
	_, err := commands.Stop(h.storage)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) continueLastRecord() error {
	if cmd, err := commands.Continue(h.storage); err != nil {
		return err
	} else {
		log.Printf("record %d is set as active again", cmd.ID)
		return nil
	}
}
