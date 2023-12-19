package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/parser"
)

func main() {
	ctx := cuecontext.New()
	f, err := parser.ParseFile("test.cue", nil)
	if err != nil {
		fmt.Println(err)
	}

	value := ctx.BuildFile(f)
	fmt.Println(value.LookupPath(cue.ParsePath("name")))
	// ctx := cuecontext.New()

	// v := ctx.CompileString(config)
	// msg := v.LookupPath(cue.ParsePath("msg"))
	// fmt.Println(msg)
}
