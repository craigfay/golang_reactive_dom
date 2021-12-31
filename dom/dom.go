
package lib

import(
    "syscall/js"
)

func Element(tag string) js.Value {
    return js.Global().Get("document").Call("createElement", tag)
}

func SetAttr(el js.Value, key string, val string) {
    el.Set(key, val)
}

type EventHandler = func(this js.Value, args []js.Value) interface {}


func Event(el js.Value, eventName string, fn EventHandler) {
    el.Call("addEventListener", eventName, js.FuncOf(fn))
}

func Query(selector string) js.Value {
    return js.Global().Get("document").Call("querySelector", selector)
}

func QueryAll(selector string) js.Value {
    return js.Global().Get("document").Call("querySelectorAll", selector)
}

func Append(parent js.Value, child js.Value) {
    parent.Call("appendChild", child)
}