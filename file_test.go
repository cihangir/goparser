package goparser_test

import (
	"fmt"
	// "testing"
	"time"
)

const CONSTANT string = "this is a constant"
const CONSTANT2 string = "this is a constant"

var Variable string
var Variable2 string

// Goset is a thread safe SET data structure implementation
// The thread safety encompasses all operations on one set.
// Operations on multiple sets are consistent in that the elements
// of each set used was valid at exactly one point in time between the
// start and thepackage test
// type Hede testing.T

// gede
type Filter struct {
	//@Data.Validate.Min(1)
	//@Data.Validate.Min(1)
	minAge int

	//@Data.Validate.Max(20)
	maxAge int

	//@Data.Validate.Min(1980)
	//@Data.Validate.Max(2100)
	birthYear int

	//@Data.Validate.Len(20)
	max20CharStr string

	//@Data.Validate.Range(4,20)
	username string //username len should be between 4-20

	//@Data.Validate.Required
	//@Data.Validate.Email
	email string

	//@Data.Validate.Match("regex")
	str string

	//@Data.Validate.Format("format")
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

// gede
type Filter2 struct {
}

// this is documentation for MethodWithNoInputNoOutput
func (f *Filter) TestMethodWithNoInputNoOutput() {

}

func (f *Filter) MethodWithInputNoOutput(input string) {
}

func (f *Filter) MethodWithInputOutput(input string) string {
	return fmt.Sprintf("%s", "hede")
}

func (f *Filter) MethodWithMultipleInputOutput(input1, input2 string) string {
	return ""
}

func (f *Filter) MethodWithMultiple2InputOutput(input1 string, input2 string) string {
	return ""
}

// not shared ones

func (f *Filter) methodWithNoInputNoOutputNotShared() {

}

func (f *Filter) methodWithInputNoOutputNotShared(input string) {
}

func (f *Filter) methodWithInputOutputNotShared(input string) string {
	return ""
}

func (f *Filter) methodWithMultipleInputOutputNotShared(input1, input2 string) string {
	return ""
}

func (f *Filter) methodWithMultiple2InputOutputNotShared(input1 string, input2 string) string {
	return ""
}

// this is documentation for MethodWith2Input2Output
func (f *Filter) MethodWith2Input2Output(firstParam, secondParam string) (string, string) {
	return "first", "second"
}

// @Data.Validate.Permission
func Name() {
	// bu bi doc
}

// @Data.Validate.Permission
func Name2(gel, gt string, hede int) string {
	// reguslar gile
	return gel + ""
}
