
package main

import (
    "fmt"
    browser "syscall/js"
    DOM "github.com/craigfay/golang_reactive_dom/dom"
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
    root := DOM.Query("#root")
    div := DOM.Element("div")

    text := fmt.Sprintf("I am clicked %d time", count)
    DOM.SetAttr(div, "textContent", text)

    onClick := func(this browser.Value, args []browser.Value) interface {} {
        count++
        text := fmt.Sprintf("I am clicked %d time", count)
        DOM.SetAttr(div, "textContent", text)
        return browser.Value {}
    }

    DOM.Event(div, "click", onClick)
    DOM.Append(root, div)

    keepAlive()
}
