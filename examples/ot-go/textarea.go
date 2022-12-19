//go:build js && wasm
// +build js,wasm

package ot

import (
	"encoding/json"
	"fmt"
	"log"
	"syscall/js"
)

// TODO: handle ctrl+z by looking up state in operation history

func NewTextarea(document, parentElement js.Value, authorityURL, name string) {
	textarea := document.Call("createElement", "textarea")
	parentElement.Call("appendChild", textarea)
	textarea.Call("setAttribute", "autofocus", "")
	ws := js.Global().Get("WebSocket").New(authorityURL)
	messageData := make(chan string)
	ws.Call("addEventListener", "message", js.FuncOf(func(target js.Value, args []js.Value) interface{} {
		var (
			event = args[0]
		)
		messageData <- event.Get("data").String()
		return nil
	}))
	go func() {
		for buf := range messageData {
			var message Message
			if err := json.Unmarshal([]byte(buf), &message); err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(message)
		}
	}()
	insertOps := handleInput(textarea)
	keydownOps := handleKeydown(textarea)
	go func() {
		for {
			var opperation []Applier
			select {
			case op := <-insertOps:
				opperation = op
			case op := <-keydownOps:
				opperation = op
			}
			buf, err := json.Marshal(Message{Operation: ApplierList(opperation)})
			if err != nil {
				fmt.Println(err)
				return
			}
			ws.Call("send", string(buf))
			fmt.Printf("event: %+v\n", opperation)
		}
	}()
}

type ApplierList []Applier

func (list *ApplierList) UnmarshalJSON(data []byte) error {
	var marshaled []interface{}
	if err := json.Unmarshal(data, &marshaled); err != nil {
		return err
	}
	li := ApplierList(make([]Applier, len(marshaled)))
	list = &li
	for i, op := range marshaled {
		switch o := op.(type) {
		case string:
			(*list)[i] = Insert(o)
		case float64:
			if o < 0 {
				(*list)[i] = Delete(int(o))
			} else {
				(*list)[i] = Retain(int(o))
			}
		default:
			return fmt.Errorf("unknown op type %v %t", o, o)
		}
	}
	return nil
}

func getCaretPosition(textarea js.Value) (int, int) {
	start := textarea.Get("selectionStart").Int()
	end := textarea.Get("selectionEnd").Int()
	return start, end
}

func handleInput(textarea js.Value) chan []Applier {
	selections := make(chan Selection)
	input := make(chan string)
	cut := make(chan struct{})
	ops := make(chan []Applier)
	textarea.Call("addEventListener", "select", onSelect(textarea, selections))
	textarea.Call("addEventListener", "input", js.FuncOf(func(target js.Value, args []js.Value) interface{} {
		if val := args[0].Get("data"); val.Truthy() {
			input <- val.String()
		}
		return nil
	}))
	textarea.Call("addEventListener", "cut", js.FuncOf(func(target js.Value, args []js.Value) interface{} {
		go func() { cut <- struct{}{} }()
		return nil
	}))
	go func() {
		var selection Selection
		for {
			select {
			case sel, ok := <-selections:
				if !ok {
					continue
				}
				selection = sel
			case <-cut:
				val := textarea.Get("value")
				if !val.Truthy() {
					continue
				}
				value := val.String()
				var op []Applier
				if selection.Start == 0 && selection.End == 0 {
					op = append(op, Delete(-len(value)))
				}
				if selection.Start > 0 {
					op = append(op, Retain(selection.Start))
				}
				if rm := -len(value[selection.Start:selection.End]); rm < 0 {
					op = append(op, Delete(rm))
				}
				if retainEnd := len(value) - selection.End; retainEnd > 0 && (selection.Start != 0 || selection.End != 0) {
					op = append(op, Retain(retainEnd))
				}
				if len(op) > 0 {
					go func() { ops <- op }()
				}
				selection.End = selection.Start
			case txt, ok := <-input:
				if !ok {
					continue
				}
				var op []Applier
				if selection.Start != selection.End {
					op = append(op, Retain(selection.Start))
					op = append(op, Delete(selection.Start-selection.End))
					selection.End = selection.Start
				} else {
					start, _ := getCaretPosition(textarea)
					if startRetain := start - len(txt); startRetain > 0 {
						op = append(op, Retain(startRetain))
					}
				}
				if len(txt) > 0 {
					op = append(op, Insert(txt))
				}
				var value string
				val := textarea.Get("value")
				if !val.Truthy() {
					continue
				}
				value = val.String()
				_, end := getCaretPosition(textarea)
				if endRetain := len(value) - end; endRetain != 0 {
					op = append(op, Retain(endRetain))
				}
				if len(op) > 0 {
					go func() { ops <- op }()
				}
			}
		}
	}()
	return ops /*, func() { // close func
		textarea.Call("removeEventListener", "select")
		textarea.Call("removeEventListener", "input")
		close(selections)
		close(input)
	}*/
}

type Selection struct {
	Start, End int
}

func onSelect(textarea js.Value, selections chan<- Selection) js.Func {
	return js.FuncOf(func(target js.Value, args []js.Value) interface{} {
		start, end := getCaretPosition(textarea)
		val := textarea.Get("value").String()
		if start < len(val) && start >= 0 && end < len(val) && end >= 0 {
			selections <- Selection{start, end}
		}
		return nil
	})
}

func handleKeydown(textarea js.Value) chan []Applier {
	ops := make(chan []Applier)
	textarea.Call("addEventListener", "keydown", js.FuncOf(func(target js.Value, args []js.Value) interface{} {
		event := args[0]
		if key := event.Get("key").String(); key == "Delete" || key == "Backspace" {
			start, end := getCaretPosition(textarea)
			if start == end {
				start--
			}
			val := textarea.Get("value")
			if !val.Truthy() {
				return nil
			}
			var op []Applier
			if start > 0 {
				op = append(op, Retain(start))
			}
			if rm := start - end; rm < 0 {
				op = append(op, Delete(rm))
			}
			if retainEnd := len(val.String()) - end; retainEnd > 0 {
				op = append(op, Retain(retainEnd))
			}
			if len(op) > 0 {
				go func() { ops <- op }()
			}
		}
		return nil
	}))
	return ops
}
