package names

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
)

func Check(dir string, astFiles []*ast.File, fileSet *token.FileSet) {
	checkDir(dir)

	w := walker{fileSet: fileSet}
	for i, astFile := range astFiles {
		if i == 0 {
			checkIdent(astFile.Name, false, fileSet)
		}
		checkFile(fileSet.Position(astFile.Pos()).Filename)
		ast.Walk(w, astFile)
	}
}

func checkDir(path string) {
	if path == `.` || path == `..` || path == `/` {
		return
	}
	Rules.Dir.exec(filepath.Base(path), token.Position{Filename: path})
}

func checkFile(path string) {
	Rules.File.exec(
		strings.TrimSuffix(strings.TrimSuffix(filepath.Base(path), `.go`), `_test`),
		token.Position{Filename: path},
	)
}

func checkIdent(ident *ast.Ident, local bool, fileSet *token.FileSet) {
	if ident == nil || ident.Obj == nil {
		return
	}
	if rule := getRuleForIdent(ident, local, fileSet); rule.valid() {
		rule.exec(ident.Name, fileSet.Position(ident.Pos()))
	}
}