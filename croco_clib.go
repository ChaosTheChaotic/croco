package main

// #include <stdlib.h>
import (
	"C"

	"croco/croco"
)

//export SendText
func SendText(msg *C.char, code *C.char) *C.char {
	goMsg := C.GoString(msg)
	goCode := C.GoString(code)
	e := croco.Send(goMsg, goCode)
	if e != nil {
		return C.CString(e.Error())
	}
	return nil
}

//export RecvText
func RecvText(code *C.char) *C.char {
	goCode := C.GoString(code)
	r, e := croco.Recv(goCode)
	if e != nil {
		return C.CString(e.Error())
	}
	return C.CString(r)
}

func main() {/* To make the compiler happy */}
