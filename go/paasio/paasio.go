package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	reader   io.Reader
	bytes    int64
	count    int64
	readLock sync.Mutex
}
type writeCounter struct {
	writer     io.Writer
	writeBytes int64
	writeCount int64
	writeLock  sync.Mutex
}
type readWriteCounter struct {
	readCounter
	writeCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer, 0, 0, sync.Mutex{}}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader, 0, 0, sync.Mutex{}}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	t := new(readWriteCounter)
	t.reader = NewReadCounter(readwriter)
	t.writer = NewWriteCounter(readwriter)
	return t
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	rc.readLock.Lock()
	rc.bytes += int64(n)
	rc.count++
	rc.readLock.Unlock()
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.readLock.Lock()
	defer func() {
		rc.readLock.Unlock()
	}()
	return rc.bytes, int(rc.count)
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	wc.writeLock.Lock()
	wc.writeBytes += int64(n)
	wc.writeCount++
	wc.writeLock.Unlock()
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.writeLock.Lock()
	defer func() {
		wc.writeLock.Unlock()
	}()
	return wc.writeBytes, int(wc.writeCount)
}
