package debug

import(
	"github.com/kr/pretty"
)

func DumpVars(args ...interface{}) {
	pretty.Println("%# v", args)
}