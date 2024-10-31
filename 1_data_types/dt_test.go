package __data_types

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v 1_data_types/dt_test.go

// PrintAsBinary - принт в бинарном виде.
func PrintAsBinary(a any) {
	type iface struct {
		t, v unsafe.Pointer
	}
	p := uintptr((*(*iface)(unsafe.Pointer(&a))).v)

	t := reflect.TypeOf(a)

	for i := 0; i < int(t.Size()); i++ {
		n := *(*byte)(unsafe.Pointer(p))
		fmt.Printf("%08b ", n)
		p += unsafe.Sizeof(n)
	}

	fmt.Print("\n")
}

// Bound - ограниея дженериков.
type Bound interface {
	~uint16 | ~uint32 | ~uint64
}

func ToLittleEndian[T Bound](number T) T {
	var (
		// расчет количества байт в типе `number` для счетчика
		size = int(T(unsafe.Sizeof(number)))
		// указаетль на первый байт number
		ptr = unsafe.Pointer(&number)
		// итог
		res T
	)

	for i := 0; i < size; i++ {
		// отщипываем байт от куска памяти, который занимает `number`
		byteVal := *(*byte)(unsafe.Add(ptr, i))
		// сдвигаем все биты байта от меньшего к большему
		res <<= 8
		res += T(byteVal)
		PrintAsBinary(res)
		// на примере: 0x0000FFFF
		// 11111111 00000000 00000000 00000000
		// 11111111 11111111 00000000 00000000
		// 00000000 11111111 11111111 00000000
		// 00000000 00000000 11111111 11111111
	}

	//fmt.Printf("%X\n", result)
	return res
}

func TestSerializationProperties(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestToLittleEndianUint16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: 0x0000,
			result: 0x0000,
		},
		"test case #2": {
			number: 0xFFFF,
			result: 0xFFFF,
		},
		"test case #3": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #4": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #5": {
			number: 0x0102,
			result: 0x0201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestToLittleEndianUint64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x0000000000000000,
			result: 0x0000000000000000,
		},
		"test case #2": {
			number: 0xFFFFFFFFFFFFFFFF,
			result: 0xFFFFFFFFFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF00FF00FF,
			result: 0xFF00FF00FF00FF00,
		},
		"test case #4": {
			number: 0x00000000FFFFFFFF,
			result: 0xFFFFFFFF00000000,
		},
		"test case #5": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
