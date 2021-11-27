package main

/*
int cUpdateGoSlice(void *buf,int len){
	int count=0;
	for (int i=3;i<len;i++){
		((char *)buf)[i]=i;
		count++;
	}
	return count;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]byte, 10)
	// b := make([]byte, 10)
	ret := cgoUpdateGoSlice(a, len(a))
	fmt.Printf("ret:%d, cgoupdate result Slice:%+v\n", ret, a)
	// ret = goUpdateGoSlice(b, len(b))
	// fmt.Printf("ret:%d, goupdate result Slice:%+v\n", ret, b)

}

func cgoUpdateGoSlice(b []byte, l int) unsafe.Pointer {
	var i int = 10
	cint := C.int(i)
	ret := C.cUpdateGoSlice((unsafe.Pointer(&b[0])), C.int(l))
	if int(ret) < len(b) {
		b = b[:len(b)]
	}
	// cb := C.CBytes(b)
	// ret := C.cUpdateGoSlice(cb, C.int(len))
	// b = append(b, C.GoBytes(cb, C.int(len))...)
	fmt.Printf("i : %p , cint : %p\n", &i, &cint)

	return unsafe.Pointer(&cint)
}

func goUpdateGoSlice(b []byte, len int) int {
	b[9] = byte('2')
	return 1
}
