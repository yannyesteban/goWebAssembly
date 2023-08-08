package main

import (
	"fmt"
	"syscall/js"
)

var html = `Hello <b> Muuuuuundo</b> !!!`

func GetAlgo() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return html + args[0].String()
	})
}

func main() {

	ch := make(chan struct{}, 0)

	fmt.Println("hello Mundo 2024 !")

	js.Global().Set("getHtml", GetAlgo())

	<-ch
}
