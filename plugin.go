package whatever

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

type PluginOptions struct {
	TemplateFile string
	OutputFile   string
}

func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		TemplateFile: "whatever.tpl",
		OutputFile:   "output",
	}

	params := req.GetParameter()
	if params == "" {
		return options, nil
	}

	parts := strings.Split(params, ",")

	options.TemplateFile = parts[0]
	if len(parts) > 1 {
		options.OutputFile = path.Base(parts[1])
	}

	return options, nil
}

func RunPlugin(request *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	options, err := ParseOptions(request)
	if err != nil {
		return nil, err
	}

	customTemplate := ""

	if options.TemplateFile != "" {
		data, err := ioutil.ReadFile(options.TemplateFile)
		if err != nil {
			return nil, err
		}

		customTemplate = string(data)
	}

	t := template.Must(template.New("stuff").Parse(customTemplate))
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, request); err != nil {
		return nil, err
	}

	result := tpl.String()

	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(options.OutputFile),
		Content: proto.String(string(result)),
	})

	return resp, nil
}
