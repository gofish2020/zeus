package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//4字节存储空间
	var i int32 = 0x01020304
	size := unsafe.Sizeof(i)
	for j := 0; j < int(size); j++ {

		fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + uintptr(j))))
		//fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)))))
		// fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(byte(0)) + unsafe.Sizeof(byte(0)))))
		// fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(byte(0)) + unsafe.Sizeof(byte(0)) + unsafe.Sizeof(byte(0)))))

	}

	if *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)))) == 0x04 {
		fmt.Println("小端存储")
	} else {
		fmt.Println("大端存储")
	}
}
