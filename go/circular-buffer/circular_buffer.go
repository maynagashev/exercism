package circular

import (
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Buffer is a circular buffer of bytes.
type Buffer struct {
	data  []byte
	write int // index of the next byte to write (head)
	read  int // index of the next byte to read (tail)
	count int // number of bytes in the buffer
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size)}
}

// ReadByte reads and returns a byte from the buffer. ]
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("buffer is empty")
	}
	b.count--
	readByte := b.data[b.read]
	circularMove(&b.read, len(b.data))
	return readByte, nil
}

func circularMove(ptr *int, l int) {
	*ptr = (*ptr + 1) % l
}

// WriteByte writes a byte to the buffer.
func (b *Buffer) WriteByte(c byte) error {
	if b.count == len(b.data) {
		return errors.New("buffer is full")
	}
	b.data[b.write] = c
	b.count++
	circularMove(&b.write, len(b.data))
	return nil
}

// Overwrite writes a byte to the buffer, possibly overwriting the oldest byte.
// There is no error return, as this method never fails.
// The method always writes the byte, overwriting the oldest byte if the buffer is full.
// [1, 2] overwrite A => [A, 2]
func (b *Buffer) Overwrite(c byte) {
	b.data[b.write] = c

	// if buffer is full, move the read pointer, because we are overwriting the oldest byte
	if b.count >= len(b.data) {
		circularMove(&b.read, len(b.data))
	} else {
		b.count++
	}
	circularMove(&b.write, len(b.data))

}

// Reset resets the buffer to be empty, but retains its current size.
func (b *Buffer) Reset() {
	b.write = 0
	b.read = 0
	b.count = 0
}
