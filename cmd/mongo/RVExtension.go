package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"regexp"
	"unsafe"

	"github.com/fairytale5571/mongo/internal/app"
)

func main() {}

var funks = map[string]func(a []string) string{
	"version": app.Version,
	"setup":   app.Setup,
	"write":   app.Write,
}

//export goRVExtensionVersion
func goRVExtensionVersion(output *C.char, outputSize C.size_t) {
	printInArma(output, outputSize, "mongo_log v1")
}

func cleanInput(argv **C.char, argc int) []string {
	newArgs := make([]string, argc)
	offset := unsafe.Sizeof(uintptr(0))
	i := 0
	for i < argc {
		_arg := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + offset*uintptr(i)))
		arg := C.GoString(*_arg)
		arg = arg[1 : len(arg)-1]

		reArg := regexp.MustCompile(`""`)
		arg = reArg.ReplaceAllString(arg, `"`)

		newArgs[i] = arg
		i++
	}
	return newArgs
}

func printInArma(output *C.char, outputSize C.size_t, input string) {
	result := C.CString(input)
	defer C.free(unsafe.Pointer(result))
	size := C.strlen(result) + 1
	if size > outputSize {
		size = outputSize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export goRVExtension
func goRVExtension(output *C.char, outputSize C.size_t, input *C.char) {
	goRVExtensionArgs(output, outputSize, input, nil, C.int(0))
}

//export goRVExtensionArgs
func goRVExtensionArgs(output *C.char, outputSize C.size_t, input *C.char, argv **C.char, argc C.int) {
	str := C.GoString(input)
	f, ok := funks[str]
	if !ok {
		printInArma(output, outputSize, fmt.Sprintf("Unkown '%v' method", str))
		return
	}
	id := f(cleanInput(argv, int(argc)))
	printInArma(output, outputSize, id)
}
