package debug

import(
	"github.com/kr/pretty"
)

func DumpVars(args ...interface{}) {
	pretty.Println("%# v", args)
}

func FDumpVars(args ...interface{}) string {
	return pretty.Sprintf("%# v", args)
}