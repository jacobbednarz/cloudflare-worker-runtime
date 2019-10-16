package main

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"rogchap.com/v8go"
)

func main() {
	js, _ := ioutil.ReadFile("file.js")

	vm, _ := v8go.NewIsolate()
	ctx1, _ := v8go.NewContext(vm)
	val, _ := ctx1.RunScript(string(js), "math.js")

	spew.Dump(val.String())
}
