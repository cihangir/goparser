package goparser

import (
	"fmt"
	"go/ast"
)

type ParsedImports struct {
	Name string
	Path string
}

func parseImports(imports []*ast.ImportSpec) []*ParsedImports {
	result := make([]*ParsedImports, 0)
	for _, imp := range imports {
		impStruct := &ParsedImports{
			Name: imp.Name.String(),
			Path: imp.Path.Value,
		}
		result = append(result, impStruct)
	}
	return result
}
