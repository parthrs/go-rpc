package serializer

import (
	"github.com/parthrs/go-rpc/gen"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "laptop.bin"
	laptop1 := gen.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}
