package main

import (
	"bytes"
	"errors"
	"io"
)

// Buffer is an bytes.buffer that also implements io.WriteSeeker.
type Buffer struct {
	buf bytes.Buffer
	pos int
}

// compile-time interface checks:
var _ io.WriteSeeker = &Buffer{}

// Write writes to the buffer of this WriterSeeker instance
func (ws *Buffer) Write(p []byte) (n int, err error) {
	// If the offset is past the end of the buffer, grow the buffer with null bytes.
	if extra := ws.pos - ws.buf.Len(); extra > 0 {
		if _, err := ws.buf.Write(make([]byte, extra)); err != nil {
			return n, err
		}
	}

	// If the offset isn't at the end of the buffer, write as much as we can.
	if ws.pos < ws.buf.Len() {
		n = copy(ws.buf.Bytes()[ws.pos:], p)
		p = p[n:]
	}

	// If there are remaining bytes, append them to the buffer.
	if len(p) > 0 {
		var bn int
		bn, err = ws.buf.Write(p)
		n += bn
	}

	ws.pos += n
	return n, err
}

// Seek seeks in the buffer of this WriterSeeker instance
func (ws *Buffer) Seek(offset int64, whence int) (int64, error) {
	newPos, offs := 0, int(offset)
	switch whence {
	case io.SeekStart:
		newPos = offs
	case io.SeekCurrent:
		newPos = ws.pos + offs
	case io.SeekEnd:
		newPos = ws.buf.Len() + offs
	}
	if newPos < 0 {
		return 0, errors.New("negative result pos")
	}
	ws.pos = newPos
	return int64(newPos), nil
}

// Reader returns a bytes.Reader which can be used to copy data from this
// buffer to a writer.
func (ws *Buffer) Reader() *bytes.Reader {
	return bytes.NewReader(ws.buf.Bytes())
}
