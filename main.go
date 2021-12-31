
package main

import (
    "fmt"
    "syscall/js"
    . "github.com/siongui/godom/wasm"
)

var signal = make(chan int)

func keepAlive() {
    for {
        m := <- signal
        if m == 0 {
            println("quit signal received")
            break
        }
    }
}

func element(tag string) js.Value {
    return js.Global().Get("document").Call("createElement", tag)
}

func setAttr(el js.Value, key string, val string) {
    el.Set(key, val)
}

type EventHandler = func(this js.Value, args []js.Value) interface {}
func event(el js.Value, eventName string, fn EventHandler) {
    el.Call("addEventListener", eventName, js.FuncOf(fn))
}

func main() {
    count := 0

    // div := Document.CreateElement("div")
    div := element("div")
    setAttr(div, "textContent", fmt.Sprintf("I am clicked %d time", count))

    onClick := func(this js.Value, args []js.Value) interface {} {
        count++
        setAttr(div, "textContent", fmt.Sprintf("I am clicked %d time", count))
        return js.Value {}
    }

    event(div, "click", onClick)

    root := Document.GetElementById("root")
    root.Call("appendChild", div)

    keepAlive()
}
