package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// WriteProtobufToBinaryFile writes a proto message to file in binary
// format
func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message to binary %w", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to file %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read from binary file %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("unable to marshal from binary file %w", err)
	}
	return nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) (err error) {
	data, err := protojson.Marshal(message)
	if err != nil {
		return fmt.Errorf("unable to marshal message to json %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to json file %w", err)
	}
	return
}
