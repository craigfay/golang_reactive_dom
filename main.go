
package main

import (
    "fmt"
    "syscall/js"
    . "github.com/craigfay/golang_reactive_dom/dom"
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

func main() {
    count := 0
    root := Query("#root")
    div := Element("div")

    text := fmt.Sprintf("I am clicked %d time", count)
    SetAttr(div, "textContent", text)

    onClick := func(this js.Value, args []js.Value) interface {} {
        count++
        text := fmt.Sprintf("I am clicked %d time", count)
        SetAttr(div, "textContent", text)
        return js.Value {}
    }

    Event(div, "click", onClick)
    Append(root, div)

    keepAlive()
}
