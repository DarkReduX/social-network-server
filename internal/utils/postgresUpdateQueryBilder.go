package utils

import (
	"errors"
	"fmt"
	"github.com/DarkReduX/social-network-server/internal/utils/updateRules"
	"strings"
)

func ValidateUpdateRequestWithRules(fields map[string]interface{}) error {
	for field := range fields {
		if available, ok := updateRules.ProfileChangeRule[field]; !ok || !available {
			return errors.New("Update request doesn't match upd rules ")
		}
	}
	return nil
}

func BuildProfileUpdateQuery(fields map[string]interface{}, id string) (string, []interface{}) {
	setFieldsQuery := ""
	args := make([]interface{}, 0, len(fields))
	argNum := 1
	for field, value := range fields {
		setFieldsQuery += field + fmt.Sprintf(" = $%d,", argNum)
		args = append(args, value)
		argNum++
	}
	args = append(args, id)
	resultQuery := "update profile set " + strings.TrimSuffix(setFieldsQuery, ",") + fmt.Sprintf(" where username = $%d", argNum)
	return resultQuery, args
}
