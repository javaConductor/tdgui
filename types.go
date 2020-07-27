package main

// GeneratedType ...
type GeneratedType struct {
	Name string `json:"name"`
}

const (
	INTEGER    = "integer"
	STRING     = "string"
	CALCULATED = "calculated"
	DECIMAL    = "decimal"
	TIME       = "time"
	DATE       = "date"
	DATETIME   = "dateTime"
	OBJECT     = "object"
)

var Names = []string{INTEGER, DECIMAL, STRING, CALCULATED, TIME, DATE, DATETIME, OBJECT}
