package procfs

import (
	"io"
	"os"
	"unsafe"
)

func ReadFileNoStat(filename string) ([]byte, error) {
	const maxBufferSize = 1024 * 1024
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := io.LimitReader(f, maxBufferSize)
	return io.ReadAll(reader)
}

func IsDigitsOnly(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func toInt[T Integer](arr []byte) (n T) {
	val := T(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}

// func toInt[T Integer](buf []byte) (n T) {
// 	for _, v := range buf {
// 		n = n*10 + T(v-'0')
// 	}
// 	return
// }
