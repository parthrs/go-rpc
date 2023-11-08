package main

import (
	gen "go-rpc/gen/go-rpc"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "laptop.bin"
	laptop1 := NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &gen.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}

func TestWriteProtobufToJSONFile(t *testing.T) {
	jsonFile := "laptop.json"
	laptop1 := NewLaptop()
	err := WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
