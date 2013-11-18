package goparser

import (
	"fmt"
	"go/ast"
)

type ParsedStruct struct {
	Name   string
	Fields []*StructField
}

type StructField struct {
	Name          string
	TypeOf        string
	Documentation string
}

// type StructType struct {
//     Struct     token.Pos  // position of "struct" keyword
//     Fields     *FieldList // list of field declarations
//     Incomplete bool       // true if (source) fields are missing in the Fields list
// }

// type Field struct {
//     Doc     *CommentGroup // associated documentation; or nil
//     Names   []*Ident      // field/method/parameter names; or nil if anonymous field
//     Type    Expr          // field/method/parameter type
//     Tag     *BasicLit     // field tag; or nil
//     Comment *CommentGroup // line comments; or nil
// }

func parseStructs(specs []ast.Spec) *ParsedStruct {
	// type TypeSpec struct {
	//  Doc     *CommentGroup // associated documentation; or nil
	//  Name    *Ident        // type name
	//  Type    Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
	//  Comment *CommentGroup // line comments; or nil
	// }
	parsedStruct := &ParsedStruct{}
	if len(specs) > 1 {
		panic("noliii")
	}
	for _, s := range specs {
		tSpec := s.(*ast.TypeSpec)
		// todo check type names against javascript reserved words
		parsedStruct.Name = tSpec.Name.String()
		switch typ := tSpec.Type.(type) {

		case *ast.StructType:
			parsedStruct.Fields = parseStructFields(typ)
			// type StructType struct {
			//     Struct     token.Pos  // position of "struct" keyword
			//     Fields     *FieldList // list of field declarations
			//     Incomplete bool       // true if (source) fields are missing in the Fields list
			// }

		case *ast.ArrayType:
			// type ArrayType struct {
			//     Lbrack token.Pos // position of "["
			//     Len    Expr      // Ellipsis node for [...]T array types, nil for slice types
			//     Elt    Expr      // element type
			// }

		case *ast.MapType:
			// type MapType struct {
			//     Map   token.Pos // position of "map" keyword
			//     Key   Expr
			//     Value Expr
			// }

		case *ast.Ident:
			// type Ident struct {
			//     NamePos token.Pos // identifier position
			//     Name    string    // identifier name
			//     Obj     *Object   // denoted object; or nil
			// }
		case *ast.SelectorExpr:
			// type SelectorExpr struct {
			//    X   Expr   // expression
			//    Sel *Ident // field selector
			// }
		case *ast.FuncType, *ast.InterfaceType:

		default:
			panic(fmt.Sprintf("unimplemented: %T", typ))
		}

	}
	return parsedStruct
}

func parseStructFields(structType *ast.StructType) []*StructField {

	structFields := make([]*StructField, 0)
	if structType.Fields == nil {
		return structFields
	}

	if structType.Fields.NumFields() == 0 {
		return structFields
	}

	for _, field := range structType.Fields.List {
		// normalize the parameters
		var typeOf string
		switch field.Type.(type) {
		case *ast.SelectorExpr:
			typeOf = field.Type.(*ast.SelectorExpr).Sel.String()
		case *ast.Ident:
			typeOf = field.Type.(*ast.Ident).String()
		}

		for _, name := range field.Names {

			// instead of using Doc.Text() we can handle them seperately! wtf
			structFields = append(structFields,
				&StructField{
					Name:          name.String(),
					TypeOf:        typeOf,
					Documentation: field.Doc.Text(),
				})
		}
	}
	return structFields
}
