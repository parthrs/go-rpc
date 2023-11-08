package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "laptop.bin"
	laptop1 := NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}
