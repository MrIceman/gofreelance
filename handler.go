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
	if cmd == "description" {
		if len(arg) < 2 {
			return errors.New("missing action argument, see description help")
		}
		action := arg[1]
		if ok := commands.ValidateDescriptionAction(action); !ok {
			return fmt.Errorf("%cmd is not a valid description action", action)
		}
		if action == commands.DescriptionActionAdd {
			v := strings.Join(arg[2:], " ")
			if v == "" {
				return errors.New("missing description - the syntax is gof description add <value>")
			}
			err := commands.AppendDescription(h.storage, v)
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
