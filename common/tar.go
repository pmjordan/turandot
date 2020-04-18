package common

import (
	"archive/tar"
	"io"
	"sync"
)

// TODO: unused?

//
// TarEncoder
//
// Encodes a single tar entry, named "portable"
//

type TarEncoder struct {
	reader     io.Reader
	size       int64
	pipeReader *io.PipeReader
	pipeWriter *io.PipeWriter
	waitGroup  sync.WaitGroup
}

func NewTarEncoder(reader io.Reader, size int64) *TarEncoder {
	pipeReader, pipeWriter := io.Pipe()
	return &TarEncoder{
		reader:     reader,
		size:       size,
		pipeReader: pipeReader,
		pipeWriter: pipeWriter,
	}
}

func (self *TarEncoder) Encode() io.Reader {
	self.waitGroup.Add(1)
	go self.copy()
	return self.pipeReader
}

func (self *TarEncoder) Drain() {
	self.waitGroup.Wait()
}

func (self *TarEncoder) copy() {
	defer self.waitGroup.Done()

	tarWriter := tar.NewWriter(self.pipeWriter)

	tarWriter.WriteHeader(&tar.Header{
		Name: "portable",
		Size: self.size,
	})

	if _, err := io.Copy(tarWriter, self.reader); err == nil {
		tarWriter.Close()
		self.pipeWriter.Close()
	} else {
		self.pipeWriter.CloseWithError(err)
	}
}
