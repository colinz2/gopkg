package utils

import (
	"bytes"
	"errors"
)

type FileBuffer struct {
	Buffer bytes.Buffer
	Index  int64
}

func NewFileBuffer() *FileBuffer {
	return &FileBuffer{}
}

func (F *FileBuffer) Bytes() []byte {
	return F.Buffer.Bytes()
}

func (F *FileBuffer) Read(p []byte) (int, error) {
	n, err := bytes.NewBuffer(F.Buffer.Bytes()[F.Index:]).Read(p)

	if err == nil {
		if F.Index+int64(len(p)) < int64(F.Buffer.Len()) {
			F.Index += int64(len(p))
		} else {
			F.Index = int64(F.Buffer.Len())
		}
	}

	return n, err
}

func (F *FileBuffer) Write(p []byte) (int, error) {
	n, err := F.Buffer.Write(p)

	if err == nil {
		F.Index = int64(F.Buffer.Len())
	}

	return n, err
}

func (F *FileBuffer) Seek(offset int64, whence int) (int64, error) {
	var err error
	var Index int64 = 0

	switch whence {
	case 0:
		if offset >= int64(F.Buffer.Len()) || offset < 0 {
			err = errors.New("Invalid Offset.")
		} else {
			F.Index = offset
			Index = offset
		}
	default:
		err = errors.New("Unsupported Seek Method.")
	}

	return Index, err
}
