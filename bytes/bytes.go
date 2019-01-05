package bytes

import (
	"bytes"
	"encoding/binary"
)

// 数字转换字节数组
// 支持uint16/uint32/uint64/int16/int32/int64/bool及其指针
// int和uint不可以
func ToByte(order binary.ByteOrder, n interface{}) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, order, n)
	return buf.Bytes()
}

// 字节数据转换数字
func ByteTo(order binary.ByteOrder, b []byte, n interface{}) {
	buf := bytes.NewBuffer(b)
	binary.Read(buf, order, n)
}

// ----------------

func ToByteForLE(n interface{}) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, n)
	return buf.Bytes()
}

func ByteToForLE(b []byte, n interface{}) {
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, n)
}

func Extend(dest, src []byte) []byte {
	for i := 0; i < len(src); i++ {
		dest = append(dest, src[i])
	}
	return dest
}

// 只是为了保存
func BytesToUint16(array []byte) uint16 {
	var data uint16 = 0
	for i := 0; i < len(array); i++ {
		data = data + uint16(uint(array[i])<<uint(8*i))
	}

	return data
}
