package commands

import "slices"

const (
	TaskAdd                 = "add"
	DescriptionActionDelete = "delete"
	DescriptionActionList   = "list"
)

var (
	actions = []string{
		TaskAdd,
		DescriptionActionDelete,
		DescriptionActionList,
	}
)

func ValidateDescriptionAction(action string) bool {
	return slices.IndexFunc(actions, func(s string) bool {
		return s == action
	}) != -1
}
