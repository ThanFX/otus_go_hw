package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotCreated        = errors.New("dest file not created")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fs, err := os.Stat(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > fs.Size() {
		return ErrOffsetExceedsFileSize
	}

	fi, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fi.Close()

	fo, err := os.Create(toPath)
	if err != nil {
		return ErrFileNotCreated
	}
	defer fo.Close()

	fi.Seek(offset, 0)

	if limit == 0 || limit > fs.Size() {
		limit = fs.Size()
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(io.LimitReader(fi, limit))
	_, err = io.Copy(fo, barReader)
	bar.Finish()
	if err != nil {
		return err
	}

	return nil
}
