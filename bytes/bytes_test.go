package bytes

import (
	"encoding/binary"
	"fmt"
	"testing"
)

// func TestToByte(t *testing.T) {
// 	i := 2

// 	fmt.Println("uint16", ToByte(binary.LittleEndian, uint16(i)))
// 	fmt.Println("uint32", ToByte(binary.LittleEndian, uint32(i)))
// 	fmt.Println("uint64", ToByte(binary.LittleEndian, uint64(i)))
// 	fmt.Println("int64 ", ToByte(binary.LittleEndian, int64(i)))
// 	fmt.Println("uint ", ToByte(binary.LittleEndian, uint(i)))
// }

// func TestByteTo(t *testing.T) {
// 	b := ToByte(binary.LittleEndian, uint16(2))

// 	var u16 uint16
// 	ByteTo(binary.LittleEndian, b, &u16)
// 	fmt.Println("uint16", u16)
// }

func TestMerge(t *testing.T) {
	b1 := ToByte(binary.LittleEndian, uint16(2))
	b2 := ToByte(binary.LittleEndian, uint16(4))
	b3 := ToByte(binary.LittleEndian, uint16(6))

	fmt.Println(b1, b2, b3)

	fmt.Println(Merge(b1, b2, b3))
}
