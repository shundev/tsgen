package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var (
	typeMap = map[string]TSType{
		"string": &TSString{},
		"int":    &TSNumber{},
		"uint":   &TSNumber{},
	}
)

type (
	TSType interface {
		String() string
	}
	TSString   struct{}
	TSNumber   struct{}
	TSNullable struct {
		Base TSType
	}
	TSArray struct {
		Base TSType
	}
)

func (t *TSString) String() string {
	return "string"
}

func (t *TSNumber) String() string {
	return "number"
}

func (t *TSNullable) String() string {
	return t.Base.String() + " | null"
}

func (t *TSArray) String() string {
	return "(" + t.Base.String() + ")[]"
}

type (
	TSField struct {
		Name string `json:"name"`
		Type TSType `json:"type"`
	}

	TSStruct struct {
		Name     string     `json:"name"`
		Fields   []*TSField `json:"fields"`
		Exported bool       `json:"exported"`
	}
)

func (f *TSField) String() string {
	return f.Name + ": " + f.Type.String() + ";"
}

func (s *TSStruct) String() string {
	var sb strings.Builder
	if s.Exported {
		sb.WriteString("export ")
	}
	sb.WriteString("type ")
	sb.WriteString(s.Name)
	sb.WriteString(" = {\n")
	for _, f := range s.Fields {
		sb.WriteString("  ")
		sb.WriteString(f.String())
		sb.WriteRune('\n')
	}
	sb.WriteString("};\n")
	return sb.String()
}

func NewTSField(name string) *TSField {
	return &TSField{Name: name, Type: &TSString{}}
}

func NewTSTypeFromExp(exp ast.Expr) TSType {
	switch t := exp.(type) {
	case *ast.Ident:
		if t.Obj == nil {
			v, ok := typeMap[t.Name]
			if !ok {
				panic("Invalid ident for Ident:" + t.Name)
			}
			return v
		} else {
			v, ok := t.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				fmt.Printf("Invalid type for ident obj: %T\n", t.Obj.Decl)
				panic("")
			}
			return NewTSTypeFromExp(v.Type)
		}
	case *ast.StarExpr:
		return &TSNullable{Base: NewTSTypeFromExp(t.X)}
	case *ast.ArrayType:
		return &TSArray{Base: NewTSTypeFromExp(t.Elt)}
	case *ast.SelectorExpr:
		// uuid.UUID, time.Time and so on...
		return &TSString{}
	}

	return nil
}

func NewTSSTruct(name string) *TSStruct {
	return &TSStruct{Name: name, Fields: []*TSField{}, Exported: true}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Please specify a file path")
		os.Exit(1)
	}

	fpath := os.Args[1]

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, fpath, nil, 0)
	if err != nil {
		panic(err)
	}

	structs := []*TSStruct{}
	var cur *TSStruct
	var numFields int

	for _, decl := range parsed.Decls {
		ast.Inspect(decl, func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.TypeSpec:
				cur = NewTSSTruct(t.Name.Name)
			case *ast.StructType:
			case *ast.FieldList:
				numFields = t.NumFields()
			case *ast.Field:
				s := strings.ReplaceAll(strings.ReplaceAll(t.Tag.Value, "`", ""), "\"", "")
				ss := strings.Split(s, ":")
				name := ss[1]
				field := NewTSField(name)

				var typ TSType
				switch t := t.Type.(type) {
				case *ast.Ident:
					typ = NewTSTypeFromExp(t)
				case *ast.StarExpr:
					typ = NewTSTypeFromExp(t)
				case *ast.ArrayType:
					typ = NewTSTypeFromExp(t)
				case *ast.SelectorExpr:
					typ = NewTSTypeFromExp(t)
				case *ast.StructType:
					// TODO: recursive definition
				default:
					fmt.Printf("Unexpected node for type: %T\n", t)
					panic("")
				}

				field.Type = typ
				cur.Fields = append(cur.Fields, field)
				numFields--
				if numFields <= 0 {
					structs = append(structs, cur)
					cur = nil
				}
			case ast.Decl:
			default:
				return false
			}

			return true
		})
	}

	for _, st := range structs {
		fmt.Println(st.String())
	}
}
