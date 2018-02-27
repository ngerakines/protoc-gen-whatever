# protoc-gen-whatever

This is a plugin that allows for protocol buffer definitions to be used as inputs to Golang's text template library to generate files.

# Example Usage

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

# Credit

This project was inspired by the work of [David Muto](https://github.com/pseudomuto) on [https://github.com/pseudomuto/protoc-gen-doc/](https://github.com/pseudomuto/protoc-gen-doc/).

# License

Open source under the MIT license.

Copyright (c) 2018 Nick Gerakines (ngerakines)
