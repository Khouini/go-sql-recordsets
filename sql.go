package main

import (
	"database/sql"
	"reflect"
	"strings"
)

// Generic function to scan any struct slice from recordset
func scanRecordset(rows *sql.Rows, slicePtr interface{}) error {
	sliceValue := reflect.ValueOf(slicePtr).Elem()
	elementType := sliceValue.Type().Elem()

	// Get column names from the recordset
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		// Create new instance of the struct
		elementPtr := reflect.New(elementType)
		element := elementPtr.Elem()

		// Create slice to hold scan destinations
		scanDests := make([]interface{}, len(columns))
		fieldMap := make(map[string]reflect.Value)

		// Map struct fields by db tag
		for i := 0; i < element.NumField(); i++ {
			field := element.Type().Field(i)
			dbTag := field.Tag.Get("db")
			if dbTag != "" {
				fieldMap[strings.ToLower(dbTag)] = element.Field(i)
			}
		}

		// Prepare scan destinations
		for i, colName := range columns {
			if field, exists := fieldMap[strings.ToLower(colName)]; exists && field.CanSet() {
				scanDests[i] = field.Addr().Interface()
			} else {
				// If field doesn't exist in struct, scan to dummy variable
				var dummy interface{}
				scanDests[i] = &dummy
			}
		}

		if err := rows.Scan(scanDests...); err != nil {
			return err
		}

		// Append to slice
		sliceValue.Set(reflect.Append(sliceValue, element))
	}

	return rows.Err()
}

// scanMultipleRecordsets - One line to scan all recordsets
func scanMultipleRecordsets(rows *sql.Rows, result interface{}) error {
	resultValue := reflect.ValueOf(result).Elem()
	resultType := resultValue.Type()

	for i := 0; i < resultType.NumField(); i++ {
		if i > 0 && !rows.NextResultSet() {
			break // No more recordsets
		}

		field := resultValue.Field(i)
		if field.Kind() == reflect.Slice && field.CanSet() {
			if err := scanRecordset(rows, field.Addr().Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}
