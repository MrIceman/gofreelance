package commands

import "slices"

const (
	DescriptionActionAdd    = "add"
	DescriptionActionDelete = "delete"
	DescriptionActionList   = "list"
)

var (
	actions = []string{
		DescriptionActionAdd,
		DescriptionActionDelete,
		DescriptionActionList,
	}
)

func ValidateDescriptionAction(action string) bool {
	return slices.IndexFunc(actions, func(s string) bool {
		return s == action
	}) != -1
}
