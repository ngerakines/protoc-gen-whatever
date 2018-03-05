package whatever

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

type PluginOptions struct {
	TemplateFile string
	Suffix       string
}

func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		TemplateFile: "whatever.tpl",
	}

	params := req.GetParameter()
	if params == "" {
		return options, nil
	}

	parts := strings.Split(params, ",")

	options.TemplateFile = parts[0]
	if len(parts) > 1 {
		options.Suffix = parts[1]
	}

	return options, nil
}

func RunPlugin(request *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	options, err := ParseOptions(request)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(options.TemplateFile)
	if err != nil {
		return nil, err
	}

	templateData := string(data)

	funcMap := template.FuncMap{
		"title":           strings.Title,
		"lastPart":        lastPart,
		"splitCammelJoin": splitCammelJoin,
		"hasServices":     hasServices,
	}

	resp := new(plugin_go.CodeGeneratorResponse)

	for _, protoFile := range request.ProtoFile {

		outputFile := generatedFileName(protoFile, options)

		t := template.Must(template.New(options.TemplateFile).Funcs(funcMap).Parse(templateData))
		var tpl bytes.Buffer
		if err := t.Execute(&tpl, protoFile); err != nil {
			return nil, err
		}

		result := tpl.String()

		trimmedResult := strings.TrimSpace(result)
		if trimmedResult != "" {
			resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
				Name:    proto.String(outputFile),
				Content: proto.String(string(result)),
			})
		}

	}

	return resp, nil
}

func generatedFileName(protoFile *descriptor.FileDescriptorProto, options *PluginOptions) string {
	fileName := protoFile.GetName()
	fileName = strings.Replace(fileName, "/", "_", -1)
	if strings.HasSuffix(fileName, ".proto") {
		fileName = fileName[:len(fileName)-len(".proto")]
	}
	outputFile := fileName + options.Suffix + ".go"

	return outputFile
}
