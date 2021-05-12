package main

import (
	"encoding/base64"
	"syscall/js"
)

//https://tutorialedge.net/golang/go-webassembly-tutorial/
//cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
//GOARCH=wasm GOOS=js go build -o lib.wasm main.go

func main() {

	c := make(chan struct{}, 0)

	println("WASN Go Initializing...")

	addFunc := js.FuncOf(add)
	js.Global().Set("add", addFunc)
	defer addFunc.Release()

	getFileFunc := js.FuncOf(getFile)
	js.Global().Set("getFile", getFileFunc)
	defer getFileFunc.Release()

	//convertFunc := js.FuncOf(converter)
	//js.Global().Set("convert", convertFunc)
	//defer convertFunc.Release()

	ba64Func := js.FuncOf(ba64)
	js.Global().Set("mybase64", ba64Func)
	defer ba64Func.Release()

	//1. Adding an <h1> element in the HTML document
	document := js.Global().Get("document")
	p := document.Call("createElement", "h1")
	p.Set("innerHTML", "Hello from Golang!")

	document.Get("body").Call("appendChild", p) //2. Exposing go functions/values in javascript variables.
	js.Global().Set("goVar", "I am a variable set from Go")

	<-c
}

//
//func registerCallbacks() {
//
//	helloText := js.Global().Get("document").Call("getElementById", "file-selector")
//	helloText.Set("innerHTML", "Go Was Here :))")
//}

func add(this js.Value, args []js.Value) interface{} {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i].Int()
	}
	return js.ValueOf(sum)
}

func getFile(this js.Value, args []js.Value) interface{} {
	for i := 0; i < len(args); i++ {
		println(args[i].String())
	}
	return nil
}

//func converter(this js.Value, inputs []js.Value) interface{} {
//	imageArr := inputs[0]
//	options := inputs[1].String()
//	inBuf := make([]uint8, imageArr.Get("byteLength").Int())
//	js.CopyBytesToGo(inBuf, imageArr)
//	convertOptions := convert.Options{}
//	err := json.Unmarshal([]byte(options), &convertOptions)
//	if err != nil {
//		convertOptions = convert.DefaultOptions
//	}
//	converter := convert.NewImageConverter()
//
//	return converter.ImageFile2ASCIIString(inBuf, &convertOptions)
//
//}

func ba64(this js.Value, inputs []js.Value) interface{} {
	imageArr := inputs[0]
	inBuf := make([]uint8, imageArr.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, imageArr)
	return base64.StdEncoding.EncodeToString(inBuf)
}
