package whatever

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func lastPart(splitter, input interface{}) interface{} {
	splitterString := realString(splitter)
	if splitterString == "" {
		return input
	}
	inputString := realString(input)
	if inputString == "" {
		return input
	}

	parts := strings.Split(inputString, splitterString)
	if len(parts) == 0 {
		return inputString
	}

	return parts[len(parts)-1]
}

func splitCammelJoin(splitter, input interface{}) interface{} {
	inputString := realString(input)
	if inputString == "" {
		return inputString
	}
	splitterString := realString(splitter)
	if splitterString == "" {
		return inputString
	}

	parts := strings.Split(inputString, splitterString)
	var result bytes.Buffer
	for _, part := range parts {
		result.WriteString(strings.Title(part))
	}
	return result.String()
}

func realString(input interface{}) string {
	stringValue, ok := input.(string)
	if ok {
		return stringValue
	}
	pointerValue, ok := input.(*string)
	if ok {
		return *pointerValue
	}
	return fmt.Sprintf("%s", input)
}

func hasServices(input interface{}) interface{} {
	fileDescriptorProto, ok := input.(*descriptor.FileDescriptorProto)
	if !ok {
		return false
	}
	return len(fileDescriptorProto.Service) > 0
}
