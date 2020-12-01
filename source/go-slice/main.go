package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//https://www.cnblogs.com/sunsky303/p/11820500.html

type pointer struct {
	low    [2]byte
	middle uint16
	high   uint64
}

func slice2struct() {

	fmt.Println("切片转结构体")
	var i = [...]byte{1, 2, 3, 4, 5}
	//var i uint = 0x0012
	var obj []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&obj))
	sliceHeader.Cap = 2
	sliceHeader.Len = 2
	sliceHeader.Data = uintptr(unsafe.Pointer(&i))
	fmt.Println(obj)
}

func struct2Slice() {
	//初始化结构体
	var arr = [2]byte{0, 0}
	var s1 = struct {
		array unsafe.Pointer
		len   int
		cap   int
	}{unsafe.Pointer(&arr), 2, 2}
	//将结构指针转化为切片指针*[]byte
	s2 := (*[]byte)(unsafe.Pointer(&s1))
	fmt.Println("结构体转切片")
	fmt.Println(*s2)
	fmt.Println(arr)
	//通过切片修改底层数值
	(*s2)[0] = 0
	(*s2)[1] = 1
	fmt.Println("切片赋值")
	fmt.Println(*s2)
	fmt.Println(arr)
	//通过数组修改
	fmt.Println("数组赋值")
	arr[1] = 100
	fmt.Println(*s2) //切片赋值后
	fmt.Println(arr)

	//结构体指针和切片指针指向的对象是同一个地址
	fmt.Println("切片和结构体存储地址")
	fmt.Println(uintptr(unsafe.Pointer(&s1)))
	fmt.Println(uintptr(unsafe.Pointer(s2)))
}

func copyArray() {
	// 数组拷贝
	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 4
	fmt.Println("原始数组:", arr1)
	fmt.Println("拷贝数组:", arr2)
}

func resizeArray() {
	// fmt.Println("append未发生扩容,老数组")
	// arr := [4]int{10, 20, 30, 40}
	// slice := arr[0:2]             //构造切片len=2,cap=4,array指针指向arr
	// newSlice := append(slice, 50) //新切片len=3,cap=4,array指针指向arr,新增的元素并没有超过容量限制,所以没有发生扩容,底层的数组一样
	// fmt.Println(slice)            // [10 20]
	// fmt.Println(newSlice)         // [10 20 50]
	// newSlice[1] += 10
	// fmt.Println(slice)    //[10 30]
	// fmt.Println(newSlice) // [10 30 50]
	// fmt.Println(arr)       // [10 30 50 40]

	fmt.Println("append发生扩容,新数组")
	arr := [4]int{10, 20, 30, 40}
	slice := arr[0:2:2]           //构造切片len=2,cap=2,array指针指向arr
	newSlice := append(slice, 50) //新切片len=3,cap=4,array指针指向内部新分配的连续存储空间
	fmt.Println(slice)            // [10 20]
	fmt.Println(newSlice)         // [10 20 50]
	newSlice[1] += 10
	fmt.Println(slice)    //[10 20]
	fmt.Println(newSlice) // [10 30 50]  从这里可以看出加10只修改了newSlice,因为扩容导致底层的数组已经不是一个
	fmt.Println(arr)      // [10 20 30 40]
}

func growslice() {
	a := []byte{1, 2, 3} // len=3 ,cap = 3
	a = append(a, 4, 5)
	fmt.Println("a", len(a), cap(a)) //按照规则应该是2*old.cap=2*3=6,roundupsize(6*1)向上取整为8, 8/1=8,最终的cap=8,len=5

	b := []int64{1, 2, 3} // len = 3,cap = 3
	b = append(b, 4, 5)
	fmt.Println("b", len(b), cap(b)) //按照规格 2*old.cap=2*3=6, roundupsize(6*8) = 48,满足取整,48/8=6,最终的cap=6,len=5
}

func copySlice() {
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice) //4 [10 20 30 40 0 0]
}

func main() {

	copyArray()    //拷贝数组
	slice2struct() //切片转结构体
	struct2Slice() //结构体转切片
	growslice()    // 切片扩容规则验证
	resizeArray()  //切片底层数组重新分配验证
	copySlice()    // 拷贝切片
	return

	// //使用指针的方式填充结构体
	// p := &pointer{}

	// lowPointer := (*[2]byte)(unsafe.Pointer(p))
	// *lowPointer = [2]byte{1, 2}

	// middlePointer := (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.middle)))
	// *middlePointer = 12

	// highPointer := (*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.high)))
	// *highPointer = 12212
	// fmt.Println(p)

}
