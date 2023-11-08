package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

// WriteProtobufToBinaryFile writes a proto message to file in binary
// format
func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Unable to marshal proto message to binary %w", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("Unable to write to file %w", err)
	}

	return nil
}
