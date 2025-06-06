package goslog

import "sync"

type buffer []byte

var bufPool = sync.Pool{
	New: func() interface{} {
		b := make(buffer, 0, 1024)
		return (*buffer)(&b)
	},
}

func newBuffer() *buffer {
	return bufPool.Get().(*buffer)
}

func (b *buffer) Free() {
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}

func (b *buffer) Write(bytes []byte) (int, error) {
	*b = append(*b, bytes...)
	return len(bytes), nil
}

func (b *buffer) WriteByte(char byte) error {
	*b = append(*b, char)
	return nil
}

func (b *buffer) WriteString(str string) (int, error) {
	*b = append(*b, str...)
	return len(str), nil
}

func (b *buffer) Bytes() []byte {
	return *b
}
