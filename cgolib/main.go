package cgolib

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L./src -lmylib -Wl,-rpath=./src
#include "a.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main_test() {
	fmt.Println("-------------------------------")

	str := C.CString("Hello C library")
	C.testCFunc(str)
	C.free(unsafe.Pointer(str))

	fmt.Println("-------------------------------")
}
