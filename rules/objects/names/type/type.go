package typepkg

import (
	"go/ast"
	"go/token"

	"github.com/lovego/gospec/rules/name"
)

var Rule = name.Rule{
	MaxLen: 30,
	Style:  "camelCase",
}
var LocalRule = name.Rule{
	MaxLen: 20,
	Style:  "lowerCamelCase",
}

func Check(local bool, node ast.Node, fileSet *token.FileSet) {
	switch typ := node.(type) {
	case *ast.TypeSpec:
		checkType(local, typ, fileSet)
	}
}

func checkType(local bool, typ *ast.TypeSpec, fileSet *token.FileSet) {
	ident := typ.Name
	if local {
		LocalRule.Exec(ident.Name, `local type`, `localType`, fileSet.Position(ident.Pos()))
	} else {
		Rule.Exec(ident.Name, `type`, `type`, fileSet.Position(ident.Pos()))
	}
}
