// +build js,wasm

package main

import (
	"fmt"
	"net/http"
	"syscall/js"
)

func doFetch() {
	res, err := http.Get("http://localhost:8181")
	fmt.Println(res, err)
}

func goFetch(this js.Value, inputs []js.Value) interface{} {
	doFetch()
	return nil
}

func main() {
	fmt.Println("Running doFetch() from wasm directly...")
	doFetch()
	fmt.Println("Now run goFetch() from javascript (this will crash):")
	js.Global().Set("goFetch", js.FuncOf(goFetch))
	select {}
}
