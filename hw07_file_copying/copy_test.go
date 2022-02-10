package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	defer os.RemoveAll("/temp")

	t.Run("Source file not found", func(t *testing.T) {
		err := Copy("temp/input1.txt", "temp/out.txt", 0, 0)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("Dest file not created", func(t *testing.T) {
		err := Copy("testdata/input.txt", "bin/bash/out.txt", 0, 0)
		require.ErrorIs(t, err, ErrFileNotCreated)
	})

	t.Run("Offset gt filesize", func(t *testing.T) {
		err := Copy("testdata/input.txt", "tmp/out.txt", 100000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})
}
