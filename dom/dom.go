
package lib

import(
    browser "syscall/js"
)

type EventHandler = func(this browser.Value, args []browser.Value) interface {}

func Element(tag string) browser.Value {
    return browser.Global().Get("document").Call("createElement", tag)
}

func SetAttr(el browser.Value, key string, val string) {
    el.Set(key, val)
}

func Event(el browser.Value, eventName string, fn EventHandler) {
    el.Call("addEventListener", eventName, browser.FuncOf(fn))
}

func Query(selector string) browser.Value {
    return browser.Global().Get("document").Call("querySelector", selector)
}

func QueryAll(selector string) browser.Value {
    return browser.Global().Get("document").Call("querySelectorAll", selector)
}

func Append(parent browser.Value, child browser.Value) {
    parent.Call("appendChild", child)
}
