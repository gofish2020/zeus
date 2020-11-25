package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	a bool
	b int8
	c uint16
	d uint32
	e int64
	f bool
}

type R struct {
	a bool
	b int8
	d uint32
	c uint16
	e int64
	f bool
}

func main() {

	var i bool = true
	var j int = 40
	fmt.Println("bool-size", unsafe.Sizeof(i))
	fmt.Println("int-size", unsafe.Sizeof(j))

	fmt.Println("bool-align", unsafe.Alignof(i))
	fmt.Println("int-align", unsafe.Alignof(j))

	type temp struct {
		a bool
		b int
	}
	var tmp = temp{}

	fmt.Println("a", unsafe.Alignof(tmp.a))
	fmt.Println("b", unsafe.Alignof(tmp.b))

	var t = T{}

	fmt.Println("t占用的实际内存大小:", unsafe.Sizeof(t), "字节,结构体对齐保证:", unsafe.Alignof(t))
	fmt.Println("a:", unsafe.Sizeof(t.a), "字节,字段对齐保证:", unsafe.Alignof(t.a), ",偏移地址:", unsafe.Offsetof(t.a))
	fmt.Println("b:", unsafe.Sizeof(t.b), "字节,字段对齐保证:", unsafe.Alignof(t.b), ",偏移地址:", unsafe.Offsetof(t.b))
	fmt.Println("c:", unsafe.Sizeof(t.c), "字节,字段对齐保证:", unsafe.Alignof(t.c), ",偏移地址:", unsafe.Offsetof(t.c))
	fmt.Println("d:", unsafe.Sizeof(t.d), "字节,字段对齐保证:", unsafe.Alignof(t.d), ",偏移地址:", unsafe.Offsetof(t.d))
	fmt.Println("e:", unsafe.Sizeof(t.e), "字节,字段对齐保证:", unsafe.Alignof(t.e), ",偏移地址:", unsafe.Offsetof(t.e))
	fmt.Println("f:", unsafe.Sizeof(t.f), "字节,字段对齐保证:", unsafe.Alignof(t.f), ",偏移地址:", unsafe.Offsetof(t.f))
	fmt.Println(uintptr(unsafe.Pointer(&t)))
	/*
		如果按照每个字段大小进行计算:17字节,但是这里为什么是24字节?因为go编译器内部做了内存对齐,对齐的原则就是基于对齐保证
		并且类型大小一定是对齐保证的整数倍,24=8*3,3倍,不管是结构体还是内置的类型,都满足
		结构体的对齐保证数值:是求各个字段的对齐保证中最大的对齐保证,作为结构体体的对其保证

		输出结果:
		t占用的实际内存大小: 24 字节,结构体对齐保证: 8
		a: 1 字节,字段对齐保证: 1 ,偏移地址: 0
		b: 1 字节,字段对齐保证: 1 ,偏移地址: 1
		c: 2 字节,字段对齐保证: 2 ,偏移地址: 2
		d: 4 字节,字段对齐保证: 4 ,偏移地址: 4
		e: 8 字节,字段对齐保证: 8 ,偏移地址: 8
		f: 1 字节,字段对齐保证: 1 ,偏移地址: 16
		824634378592
	*/

	var r = R{}
	fmt.Println("r占用的实际内存大小:", unsafe.Sizeof(r), "字节,结构体对齐保证:", unsafe.Alignof(r))
	fmt.Println("a:", unsafe.Sizeof(r.a), "字节,字段对齐保证:", unsafe.Alignof(r.a), ",偏移地址:", unsafe.Offsetof(r.a))
	fmt.Println("b:", unsafe.Sizeof(r.b), "字节,字段对齐保证:", unsafe.Alignof(r.b), ",偏移地址:", unsafe.Offsetof(r.b))
	fmt.Println("d:", unsafe.Sizeof(r.d), "字节,字段对齐保证:", unsafe.Alignof(r.d), ",偏移地址:", unsafe.Offsetof(r.d))
	fmt.Println("c:", unsafe.Sizeof(r.c), "字节,字段对齐保证:", unsafe.Alignof(r.c), ",偏移地址:", unsafe.Offsetof(r.c))
	fmt.Println("e:", unsafe.Sizeof(r.e), "字节,字段对齐保证:", unsafe.Alignof(t.e), ",偏移地址:", unsafe.Offsetof(r.e))
	fmt.Println("f:", unsafe.Sizeof(r.f), "字节,字段对齐保证:", unsafe.Alignof(r.f), ",偏移地址:", unsafe.Offsetof(r.f))
	fmt.Println(uintptr(unsafe.Pointer(&r)))

}
