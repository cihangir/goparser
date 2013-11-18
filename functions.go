package goparser

import (
	"fmt"
	"go/ast"
)

type ParsedFunc struct {
	Name           string
	Documentation  string
	Receiver       string
	IncomingParams []*ParsedFuncParam
	OutgoingParams []*ParsedFuncParam
}

type ParsedFuncParam struct {
	Name   string
	TypeOf string
}

// type FuncDecl struct {
// 	Doc  *CommentGroup // associated documentation; or nil
// 	Recv *FieldList    // receiver (methods); or nil (functions)
// 	Name *Ident        // function/method name
// 	Type *FuncType     // position of Func keyword, parameters and results
// 	Body *BlockStmt    // function body; or nil (forward declaration)
// }

func parseFunctions(fun *ast.FuncDecl) *ParsedFunc {
	parsedFunc := &ParsedFunc{}

	if fun.Doc != nil {
		parsedFunc.Documentation = fun.Doc.Text()
	}

	parsedFunc.Receiver = getReceiver(fun.Recv)
	// fmt.Println("fun.Recv", fun.Recv)
	parsedFunc.Name = fun.Name.String()

	// parse incoming params
	incomingParams := make([]*ParsedFuncParam, 0)
	if fun.Type.Params.NumFields() > 0 {
		// iterate over params list
		for _, v := range fun.Type.Params.List {
			// normalize the parameters
			// in golang we can define parameters in a function like fun(x, y string)
			// this is valid and equals to fun(x string, y string)
			for _, name := range v.Names {
				incomingParams = append(incomingParams, &ParsedFuncParam{
					Name:   name.String(),
					TypeOf: getType(v.Type),
				})
			}
		}
	}

	outgoingParameters := make([]*ParsedFuncParam, 0)
	// parse outgoing params
	if fun.Type.Results.NumFields() > 0 {
		for _, v := range fun.Type.Results.List {
			// normalize the parameters
			if len(v.Names) > 0 {
				for _, name := range v.Names {
					outgoingParameters = append(outgoingParameters, &ParsedFuncParam{
						Name:   name.String(),
						TypeOf: getType(v.Type),
					})
				}
			} else {
				pd := &ParsedFuncParam{
					Name:   "",
					TypeOf: getType(v.Type),
				}
				outgoingParameters = append(outgoingParameters, pd)

			}

		}
	}

	parsedFunc.IncomingParams = incomingParams
	parsedFunc.OutgoingParams = outgoingParameters

	return parsedFunc
}

func getType(v interface{}) string {
	var typeOf string
	switch v.(type) {
	case *ast.Ident:
		typeOf = v.(*ast.Ident).String()
	case *ast.SelectorExpr:
		typeOf = fmt.Sprintf("%v", v.(*ast.SelectorExpr).X)
	case *ast.StarExpr:
		typeOf = fmt.Sprintf("%v", v.(*ast.StarExpr).X)
	case *ast.ArrayType, *ast.MapType:
		typeOf = "[]"
	case *ast.InterfaceType:
		typeOf = "Object"
	case *ast.FuncType:
		typeOf = "function"
	case *ast.Ellipsis:
		typeOf = "ellipsis"
	case *ast.ChanType:
		typeOf = "chan"
	default:
		panic(fmt.Sprintf("%T uncompleted incoming type", v))
	}
	return typeOf
}

// type FieldList struct {
//         Opening token.Pos // position of opening parenthesis/brace, if any
//         List    []*Field  // field list; or nil
//         Closing token.Pos // position of closing parenthesis/brace, if any
// }

func getReceiver(rec *ast.FieldList) string {
	if rec != nil {
		list := rec.List
		return getType(list[0].Type)
	} else {
		return ""
	}

}
