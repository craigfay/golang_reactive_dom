
package main

import (
	"fmt"
	"syscall/js"
	. "github.com/siongui/godom/wasm"
)

var signal = make(chan int)

func keepAlive() {
	for {
		m := <-signal
		if m == 0 {
			println("quit signal received")
			break
		}
	}
}

func main() {
	count := 0

	div := Document.CreateElement("div")
	div.Set("textContent", fmt.Sprintf("I am clicked %d time", count))

	onClick := js.FuncOf(func(this js.Value, args []js.Value) interface {} {
		count++
		div.Set("textContent", fmt.Sprintf("I am clicked %d time", count))
        return js.Value {}
	})

	div.Call("addEventListener", "click", onClick)

    root := Document.GetElementById("root")
    root.Call("appendChild", div)

	keepAlive()
}
