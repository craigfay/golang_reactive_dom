
package main

import (
    "fmt"
    "syscall/js"
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

func query(selector string) js.Value {
    return js.Global().Get("document").Call("querySelector", selector)
}

func queryAll(selector string) js.Value {
    return js.Global().Get("document").Call("querySelectorAll", selector)
}

func append(parent js.Value, child js.Value) {
    parent.Call("appendChild", child)
}

func main() {
    count := 0

    div := element("div")

    text := fmt.Sprintf("I am clicked %d time", count)
    setAttr(div, "textContent", text)

    onClick := func(this js.Value, args []js.Value) interface {} {
        count++
        text := fmt.Sprintf("I am clicked %d time", count)
        setAttr(div, "textContent", text)
        return js.Value {}
    }

    event(div, "click", onClick)

    root := query("#root")
    append(root, div)

    keepAlive()
}
