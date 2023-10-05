package code

import (
	"fmt"
	"strings"
)

/* TODO change type to string */
type AuditEventAction int

const (
	AuditEventActionCreate AuditEventAction = iota
	AuditEventActionRead
	AuditEventActionUpdate
	AuditEventActionDelete
	AuditEventActionExecute
)

func AllAuditEventAction() []AuditEventAction {
	return []AuditEventAction{
		AuditEventActionCreate,
		AuditEventActionRead,
		AuditEventActionUpdate,
		AuditEventActionDelete,
		AuditEventActionExecute,
	}
}

func FindAddressUses(filter string) []AuditEventAction {
	ret := make([]AuditEventAction, 0)
	for _, at := range AllAuditEventAction() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au AuditEventAction) ToString() {
	fmt.Println(au.String())
}
func (au AuditEventAction) String() string {
	switch au {
	case AuditEventActionCreate:
		return "Create"
	case AuditEventActionRead:
		return "Read"
	case AuditEventActionUpdate:
		return "Update"
	case AuditEventActionDelete:
		return "Delete"
	case AuditEventActionExecute:
		return "Execute"
	default:
		return "Unknown AuditEvent Action"
	}
}

/*
Create
Read
Update
Delete
Execute
*/
