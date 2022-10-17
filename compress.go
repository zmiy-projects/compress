package compress

import (
	"compressions/counter"
)

func Compress(value string) string {
	compressor := counter.CreateCounter(value)
	return compressor.Compress()
}
