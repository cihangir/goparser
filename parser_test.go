package goparser

import (
	"fmt"
	// "reflect"
	"testing"
)

func TestBatchCreation(t *testing.T) {
	fiel, err := ParseFile("test.go")
	if err != nil {
		t.Error(err)
	}
	ef := fiel.GetExportedFunctions()
	for _, f := range ef {
		fmt.Println(f.ConvertToJSFunc())
	}
}
