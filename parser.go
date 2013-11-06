package goparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type ParsedFile struct {
	Name string
	Path string
	// tests implemented
	PackageName string
	// tests implemented
	Imports []*ParsedImports
	// tests are completed
	Functions []*ParsedFunc

	Structs []*ParsedStruct
}

func (pf *ParsedFile) GetExportedFunctions() []*ParsedFunc {
	ef := make([]*ParsedFunc, 0)
	for _, fun := range pf.Functions {
		if ast.IsExported(fun.Name) {
			ef = append(ef, fun)
		}
	}
	return ef
}

func ParseFile(fileName string) (*ParsedFile, error) {
	fset := token.NewFileSet() // positions are relative to fset

	file := &ParsedFile{}
	file.Path = fileName
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// set package name
	file.PackageName = node.Name.String()

	// parse imports
	file.Imports = parseImports(node.Imports)

	for _, decl := range node.Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			fun := decl.(*ast.FuncDecl)
			file.Functions = append(file.Functions, parseFunctions(fun))
			// type GenDecl struct {
			//     Doc    *CommentGroup // associated documentation; or nil
			//     TokPos token.Pos     // position of Tok
			//     Tok    token.Token   // IMPORT, CONST, TYPE, VAR
			//     Lparen token.Pos     // position of '(', if any
			//     Specs  []Spec
			//     Rparen token.Pos // position of ')', if any
			// }
		case *ast.GenDecl:
			genDecl := decl.(*ast.GenDecl)

			switch genDecl.Tok {
			case token.IMPORT:
				fmt.Println("import not implemented    ", genDecl.Specs)
			case token.CONST:
				fmt.Println("constant not implemented  ", genDecl.Specs)
			case token.VAR:
				fmt.Println("var not implemented       ", genDecl.Specs)
			case token.TYPE:
				file.Structs = append(file.Structs, parseStructs(genDecl.Specs))
			}

		default:
			panic(fmt.Sprintf("unimplemented: %T", decl))
		}
	}
	return file, nil
}
