# protoc-gen-whatever

This is a plugin that allows for protocol buffer definitions to be used as inputs to Golang's text template library to generate files.

# Installation

Instlal with `go get`.

    $ go get -u github.com/ngerakines/protoc-gen-whatever/cmd/protoc-gen-whatever

# Plugin Configuration

This plugin is used when the `--whatever_out` and `--whatever_opt` arguments are provided to the base `protoc` command.

The `--whatever_out` argument is used to set the location of generated output files.

The `--whatever_opt` argument is used to set the input template and optional output file name. This argument is optional and defaults to `--whatever_opt=whatever.tpl`.

For example, with `--whatever_out=. --whatever_opt=input.tpl` the plugin will look for a file named `input.tpl` in the current working directory and generate a file namee `output` in the current working directory.

With `--whatever_out=cmd --whatever_opt=templates/cobra.tpl,commands.go`, the plugin will look for the `templates/cobra.tpl` file as the template and render the file `cmd/commands.go`.

# Templates

The input to the template is the `CodeGeneratorRequest` structure as defined in `google/protobuf/compiler/plugin.proto`.

# Usage

With the proto file `simple.proto`:

```protobuf
syntax = "proto3";

package test;

message Example {
    string label = 1;
}

service Foo {
    rpc GetFoo(GetFooReq) returns (GetFooRes);
}

message GetFooReq {

}

message GetFooRes {

}
```

And the template file `simple.tpl`:

```
{{range .ProtoFile}}{{.Name}}{{range .MessageType}}
*{{.Name}}{{end}}{{end}}
```

Running the command:

    $ protoc --plugin=protoc-gen-whatever --whatever_out=. --whatever_opt=simple.tpl,output.txt simple.proto

Yields the file `output.txt`:

```
simple.proto
*Example
*GetFooReq
*GetFooRes
```

# Release

To release new versions, use [mage](https://github.com/magefile/mage).

    $ go get github.com/goreleaser/goreleaser
    $ go get -u -d github.com/magefile/mage
    $ mage clean
    $ TAG=0.2.0 mage release

# Credit

This project was inspired by the work of [David Muto](https://github.com/pseudomuto) on [https://github.com/pseudomuto/protoc-gen-doc/](https://github.com/pseudomuto/protoc-gen-doc/).

# License

Open source under the MIT license.

Copyright (c) 2018 Nick Gerakines (ngerakines)
