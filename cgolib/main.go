package cgolib

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L./src -lmylib
#include "a.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------")

	//str := C.CString("Hello C library")
	C.testCFunc()
	//C.free(unsafe.Pointer(str))

	fmt.Println("-------------------------------")
}
