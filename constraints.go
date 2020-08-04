package main

import (
	"time"
)

type Constraint struct {
	Name        string      `json:"name"`
	DisplayName string      `json:"displayName"`
	Value       interface{} `json:"value"`
}

func NewAlphaNumericConstraint(bool2 bool) *Constraint {
	c := Constraint{Name: "alphaNumeric", DisplayName: "Alpha Numeric?", Value: bool2}
	return &c
}

func NewMixedCaseConstraint(b bool) *Constraint {
	c := Constraint{Name: "mixedCase", DisplayName: "Mixed Case?", Value: b}
	return &c
}

func NewExpressionConstraint(expression string) *Constraint {
	c := Constraint{Name: "expression", DisplayName: "Expression", Value: expression}
	return &c
}

func NewDecimalPlacesConstraint(decimalPlaces int) *Constraint {
	c := Constraint{Name: "decimalPlaces", DisplayName: "Decimal Places", Value: decimalPlaces}
	return &c
}

func NewAfterDateConstraint(date time.Time) *Constraint {
	c := Constraint{Name: "afterDate", DisplayName: "After", Value: date}
	return &c
}

func NewBeforeDateConstraint(date time.Time) *Constraint {
	c := Constraint{Name: "beforeDate", DisplayName: "Before", Value: date}
	return &c
}

func NewMinValueConstraint(date time.Time) *Constraint {
	c := Constraint{Name: "minValue", DisplayName: "Min Value", Value: date}
	return &c
}

func NewMaxValueConstraint(date time.Time) *Constraint {
	c := Constraint{Name: "maxValue", DisplayName: "Max Value", Value: date}
	return &c
}

func NewMaxLengthConstraint(date time.Time) *Constraint {
	c := Constraint{Name: "maxLength", DisplayName: "Max Length", Value: date}
	return &c
}

func NewMinLengthConstraint(date time.Time) *Constraint {
	c := Constraint{Name: "minLength", DisplayName: "Min Length", Value: date}
	return &c
}

func NewPrefixConstraint(s string) *Constraint {
	c := Constraint{Name: "prefix", DisplayName: "Prefix", Value: s}
	return &c
}

func NewSuffixConstraint(s string) *Constraint {
	c := Constraint{Name: "suffix", DisplayName: "Suffix", Value: s}
	return &c
}

func NewInListConstraint(list []string) *Constraint {
	c := Constraint{Name: "inList", DisplayName: "Choices", Value: list}
	return &c
}

func NewDaysOfTheWeekConstraint(list []int) *Constraint {
	c := Constraint{Name: "daysOfTheWeek", DisplayName: "Days of the week", Value: list}
	return &c
}

type AutoIncrementValue struct {
	StartValue string
	increment  int
}

func NewAutoIncrementConstraint(v AutoIncrementValue) *Constraint {
	c := Constraint{Name: "autoIncrement", DisplayName: "Auto Increment", Value: v}
	return &c
}

// NumberRange ...
type NumberRange struct {
	Start float64
	End   float64
}

// DateRange ...
type DateRange struct {
	Start time.Time
	End   time.Time
}

func NewNumberRangeConstraint(r NumberRange) *Constraint {
	c := Constraint{Name: "numberRange", DisplayName: "Range", Value: r}
	return &c
}

func NewDateRangeConstraint(r DateRange) *Constraint {
	c := Constraint{Name: "dateRange", DisplayName: "Date Range", Value: r}
	return &c
}

func NewTimeRangeConstraint(r DateRange) *Constraint {
	c := Constraint{Name: "timeRange", DisplayName: "Time Range", Value: r}
	return &c
}

func NewObjectTypeNameConstraint(t string) *Constraint {
	c := Constraint{Name: "objectTypeName", DisplayName: "Object Type Name", Value: t}
	return &c
}

func NewObjectFieldConstraints(fc Constraint) *Constraint {
	c := Constraint{Name: "objectFieldConstraints", DisplayName: "Object Field Constraints", Value: fc}
	return &c
}
