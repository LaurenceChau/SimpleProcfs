package procfs

import (
	"io"
	"os"
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

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func toInt[T Integer](buf []byte) (n T) {
	for _, v := range buf {
		n = n*10 + T(v-'0')
	}
	return
}
