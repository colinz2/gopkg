package bufferpool

import (
	"github.com/valyala/bytebufferpool"
)

var pool = new(bytebufferpool.Pool)

func GetBuffer() (buf *bytebufferpool.ByteBuffer) {
	return pool.Get()
}

func PutBuffer(buf *bytebufferpool.ByteBuffer) {
	pool.Put(buf)
}
