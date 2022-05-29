package bufferpool

import (
	"bytes"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// Get returns a buffer from the pool.
func Get() (buf *bytes.Buffer) {
	return bufferPool.Get().(*bytes.Buffer)
}

// Put returns a buffer to the pool.
// The buffer is reset before it is put back into circulation.
func Put(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
