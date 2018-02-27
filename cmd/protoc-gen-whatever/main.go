// protoc-gen-whatever is used to generate stuff from proto files.
//
// It is a protoc plugin, and can be invoked by passing `--whatever_out` and `--whatever_opt` arguments to protoc.
package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/ngerakines/protoc-gen-whatever"
)

var (
	commitHash string
	timestamp  string
	gitTag     = "0.1.0"
)

func main() {
	flags := whatever.ParseFlags(os.Stdout, os.Args)

	if flags.HasMatch() {
		if flags.ShowHelp() {
			flags.PrintHelp()
		}

		if flags.ShowVersion() {
			flags.PrintVersion()
		}

		os.Exit(flags.Code())
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read contents from stdin")
	}

	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		log.Fatal(err)
	}

	resp, err := whatever.RunPlugin(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(data); err != nil {
		log.Fatal(err)
	}
}
