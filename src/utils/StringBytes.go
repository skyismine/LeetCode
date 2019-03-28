package utils

import (
	"reflect"
	"unsafe"
)

// 在不拷贝字符的情况下将 go string 转换为 []byte
// 字符串是只读的，但转换后的 []byte 可写
func StringBytes(s string) []byte {
	var bh reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// 在不拷贝字符的情况下将 []byte 转换为 string
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 返回 string 引用数据的指针
func StringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// 返回 []byte 引用数据的指针
func BytesPointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}


