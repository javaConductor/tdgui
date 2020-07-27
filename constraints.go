package main

import (
	"fmt"
	"time"
)

// Constraint ...
type Constraint interface {
	Name() string
	DisplayName() string
	Value() interface{}
}

// ValueConstraint ...
type ValueConstraint struct {
	Name            string      `json:"name"`
	DisplayName     string      `json:"displayName"`
	ConstraintValue interface{} `json:"value"`
}

// AlphaNumericConstraint ...
type AlphaNumericConstraint struct {
	ValueConstraint
}

// Name ...
func (c *AlphaNumericConstraint) Name() string {
	return "alphaNumeric"
}

// DisplayName ...
func (c *AlphaNumericConstraint) DisplayName() string {
	return "Alpha Numeric?"
}

// Value ...
func (c *AlphaNumericConstraint) Value() interface{} {

	return true //c.Value
}

//Set ...
func (c *AlphaNumericConstraint) Set(v bool) *AlphaNumericConstraint {
	c.ConstraintValue = v
	return c
}

func (c *AlphaNumericConstraint) String() string {
	return fmt.Sprintf("[%s, %s, %v]", c.Name(), c.DisplayName(), c.Value())
}

//MixedCaseConstraint ...
type MixedCaseConstraint struct {
	ValueConstraint
}

// Name ...
func (c *MixedCaseConstraint) Name() string {
	return "mixedCase"
}

// DisplayName ...
func (c *MixedCaseConstraint) DisplayName() string {
	return "Mixed Case?"
}

// Value ...
func (c *MixedCaseConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *MixedCaseConstraint) Set(v bool) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

//ExpressionConstraint ...
type ExpressionConstraint struct {
	ValueConstraint
}

// Name ...
func (c *ExpressionConstraint) Name() string {
	return "expression"
}

// DisplayName ...
func (c *ExpressionConstraint) DisplayName() string {
	return "Expression"
}

// Value ...
func (c *ExpressionConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *ExpressionConstraint) Set(v string) *ExpressionConstraint {
	c.ConstraintValue = v
	// cc := Constraint(c)
	return c
}

func (c *ExpressionConstraint) String() string {
	return fmt.Sprintf("[%s, %s, %v]", c.Name(), c.DisplayName(), c.Value())
}

//DecimalPlacesConstraint ...
type DecimalPlacesConstraint struct {
	ValueConstraint
}

// Name ...
func (c *DecimalPlacesConstraint) Name() string {
	return "decimalPlaces"
}

// DisplayName ...
func (c *DecimalPlacesConstraint) DisplayName() string {
	return "Decimal Places"
}

// Value ...
func (c *DecimalPlacesConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *DecimalPlacesConstraint) Set(v int) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// AfterDateConstraint ...
type AfterDateConstraint struct {
	ValueConstraint
}

// Name ...
func (c *AfterDateConstraint) Name() string {
	return "afterDate"
}

// DisplayName ...
func (c *AfterDateConstraint) DisplayName() string {
	return "After"
}

// Value ...
func (c *AfterDateConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *AfterDateConstraint) Set(v time.Time) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// BeforeDateConstraint ...
type BeforeDateConstraint struct {
	ValueConstraint
}

// Name ...
func (c *BeforeDateConstraint) Name() string {
	return "beforeDate"
}

// DisplayName ...
func (c *BeforeDateConstraint) DisplayName() string {
	return "Before"
}

// Value ...
func (c *BeforeDateConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *BeforeDateConstraint) Set(v time.Time) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// MinValueConstraint ...
type MinValueConstraint struct {
	ValueConstraint
}

// Name ...
func (c *MinValueConstraint) Name() string {
	return "minValue"
}

// DisplayName ...
func (c *MinValueConstraint) DisplayName() string {
	return "Min Value"
}

// Value ...
func (c *MinValueConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *MinValueConstraint) Set(v int) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// MaxValueConstraint ...
type MaxValueConstraint struct {
	ValueConstraint
}

// Name ...
func (c *MaxValueConstraint) Name() string {
	return "maxValue"
}

// DisplayName ...
func (c *MaxValueConstraint) DisplayName() string {
	return "Max Value"
}

// Value ...
func (c *MaxValueConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *MaxValueConstraint) Set(v int) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// MaxLengthConstraint ...
type MaxLengthConstraint struct {
	ValueConstraint
}

// Name ...
func (c *MaxLengthConstraint) Name() string {
	return "maxLength"
}

// DisplayName ...
func (c *MaxLengthConstraint) DisplayName() string {
	return "Max Length"
}

// Value ...
func (c *MaxLengthConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *MaxLengthConstraint) Set(v int) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// MinLengthConstraint ...
type MinLengthConstraint struct {
	ValueConstraint
}

// Name ...
func (c *MinLengthConstraint) Name() string {
	return "minLength"
}

// DisplayName ...
func (c *MinLengthConstraint) DisplayName() string {
	return "Min Length"
}

// Value ...
func (c *MinLengthConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *MinLengthConstraint) Set(v int) *Constraint {
	c.ConstraintValue = v
	cc := Constraint(c)
	return &cc
}

// PrefixConstraint ...
type PrefixConstraint struct {
	ValueConstraint
}

// Name ...
func (c *PrefixConstraint) Name() string {
	return "prefix"
}

// DisplayName ...
func (c *PrefixConstraint) DisplayName() string {
	return "Prefix"
}

// Value ...
func (c *PrefixConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *PrefixConstraint) Set(p string) *Constraint {
	c.ConstraintValue = p
	cc := Constraint(c)
	return &cc
}

// SuffixConstraint ...
type SuffixConstraint struct {
	ValueConstraint
}

// Name ...
func (c *SuffixConstraint) Name() string {
	return "suffix"
}

// DisplayName ...
func (c *SuffixConstraint) DisplayName() string {
	return "Suffix"
}

// Value ...
func (c *SuffixConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *SuffixConstraint) Set(s string) *Constraint {
	c.ConstraintValue = s
	cc := Constraint(c)
	return &cc
}

// InListConstraint ...
type InListConstraint struct {
	ValueConstraint
}

// Name ...
func (c *InListConstraint) Name() string {
	return "inList"
}

// DisplayName ...
func (c *InListConstraint) DisplayName() string {
	return "Choices"
}

// Value ...
func (c *InListConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *InListConstraint) Set(s string) *Constraint {
	c.ConstraintValue = s
	cc := Constraint(c)
	return &cc
}

// DaysOfTheWeekConstraint ...
type DaysOfTheWeekConstraint struct {
	ValueConstraint
}

// Name ...
func (c *DaysOfTheWeekConstraint) Name() string {
	return "daysOfTheWeek"
}

// DisplayName ...
func (c *DaysOfTheWeekConstraint) DisplayName() string {
	return "Days of the week"
}

// Value ...
func (c *DaysOfTheWeekConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *DaysOfTheWeekConstraint) Set(daysOfTheWeek []int) *Constraint {
	c.ConstraintValue = daysOfTheWeek
	cc := Constraint(c)
	return &cc
}

// AutoIncrementConstraint ...
type AutoIncrementConstraint struct {
	ValueConstraint
}

// Name ...
func (c *AutoIncrementConstraint) Name() string {
	return "autoIncrement"
}

// DisplayName ...
func (c *AutoIncrementConstraint) DisplayName() string {
	return "Auto Increment"
}

// Value ...
func (c *AutoIncrementConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *AutoIncrementConstraint) Set(n int) *Constraint {
	c.ConstraintValue = n
	cc := Constraint(c)
	return &cc
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

// NumberRangeConstraint ...
type NumberRangeConstraint struct {
	ValueConstraint
}

// Name ...
func (c *NumberRangeConstraint) Name() string {
	return "Date Range"
}

// DisplayName ...
func (c *NumberRangeConstraint) DisplayName() string {
	return "Date Range"
}

// Value ...
func (c *NumberRangeConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *NumberRangeConstraint) Set(r NumberRange) *Constraint {
	c.ConstraintValue = r
	cc := Constraint(c)
	return &cc
}

// DateRangeConstraint ...
type DateRangeConstraint struct {
	ValueConstraint
}

// Name ...
func (c *DateRangeConstraint) Name() string {
	return "Date Range"
}

// DisplayName ...
func (c *DateRangeConstraint) DisplayName() string {
	return "Date Range"
}

// Value ...
func (c *DateRangeConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *DateRangeConstraint) Set(r DateRange) *Constraint {
	c.ConstraintValue = r
	cc := Constraint(c)
	return &cc
}

// TimeRangeConstraint ...
type TimeRangeConstraint struct {
	ValueConstraint
}

// Name ...
func (c *TimeRangeConstraint) Name() string {
	return "timeRange"
}

// DisplayName ...
func (c *TimeRangeConstraint) DisplayName() string {
	return "Time Range"
}

// Value ...
func (c *TimeRangeConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *TimeRangeConstraint) Set(r DateRange) *Constraint {
	c.ConstraintValue = r
	cc := Constraint(c)
	return &cc
}

// ObjectTypeNameConstraint ...
type ObjectTypeNameConstraint struct {
	ValueConstraint
}

// Name ...
func (c *ObjectTypeNameConstraint) Name() string {
	return "objectTypeName"
}

// DisplayName ...
func (c *ObjectTypeNameConstraint) DisplayName() string {
	return "Object Type Name"
}

// Value ...
func (c *ObjectTypeNameConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *ObjectTypeNameConstraint) Set(name string) *Constraint {
	c.ConstraintValue = name
	cc := Constraint(c)
	return &cc
}

// ObjectFieldConstraintsConstraint ...
type ObjectFieldConstraintsConstraint struct {
	ValueConstraint
}

// Name ...
func (c *ObjectFieldConstraintsConstraint) Name() string {
	return "objectTypeName"
}

// DisplayName ...
func (c *ObjectFieldConstraintsConstraint) DisplayName() string {
	return "Object Type Name"
}

// Value ...
func (c *ObjectFieldConstraintsConstraint) Value() interface{} {
	return c.ConstraintValue
}

//Set ...
func (c *ObjectFieldConstraintsConstraint) Set(constraints map[string]*Constraint) *ObjectFieldConstraintsConstraint {
	c.ConstraintValue = constraints
	// cc := Constraint(c)
	return c
}

//Pc ...
func Pc() ExpressionConstraint {
	c := new(ExpressionConstraint).Set("x")
	fmt.Println("ExpressionConstraint:", c)
	return *c
}
