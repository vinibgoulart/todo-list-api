package db

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/vinibgoulart/todo-list/utils"
)

func HelperDBUpdateTableById(tableName string, data interface{}, id string) (string, []interface{}) {
	var fields []string
	var args []interface{}

	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		fieldName := dataType.Field(i).Name

		fieldName = utils.ToSnakeCase(fieldName)

		fieldValue := dataValue.Field(i).Interface()

		if !reflect.DeepEqual(fieldValue, reflect.Zero(dataType.Field(i).Type).Interface()) {
			fields = append(fields, fmt.Sprintf("%s=$%d", fieldName, len(args)+1))
			args = append(args, fieldValue)
		}
	}

	updateFieldsSQL := strings.Join(fields, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING *", tableName, updateFieldsSQL, len(args)+1)
	args = append(args, id)

	return query, args
}
