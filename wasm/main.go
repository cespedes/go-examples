package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("stdout goes to the console.")

	document := js.Global().Get("document")

	button := document.Call("createElement", "button")
	button.Set("textContent", "Click me")
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("button pressed")
		return nil
	}))

	document.Get("body").Call("appendChild", button)

	select {} // make Go app running
}
