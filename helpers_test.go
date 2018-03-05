package whatever

import (
	"bytes"
	"testing"
	"text/template"

	_ "github.com/ngerakines/protoc-gen-whatever/test"
)

func TestSplitHelper(t *testing.T) {

	tests := map[string]string{
		`{{ "protocol.buffers" | lastPart "." }}`:        "buffers",
		`{{ lastPart "." ".mcpp.update.RotateKeyReq" }}`: "RotateKeyReq",
	}

	for templateData, expected := range tests {
		funcMap := template.FuncMap{
			"lastPart": lastPart,
		}

		tpl := template.Must(template.New("test").Funcs(funcMap).Parse(templateData))
		var output bytes.Buffer
		if err := tpl.Execute(&output, nil); err != nil {
			t.Fatalf("%v", err)
		}
		if output.String() != expected {
			t.Fatalf("Expected '%s' but got '%s'", expected, output.String())
		}
	}
}

func TestSplitCammelJoin(t *testing.T) {
	tests := map[string]string{
		`{{ "nick.gerakines" | splitCammelJoin "." }}`:          "NickGerakines",
		`{{ splitCammelJoin "." ".mcpp.update.RotateKeyReq" }}`: "McppUpdateRotateKeyReq",
	}

	for templateData, expected := range tests {
		funcMap := template.FuncMap{
			"splitCammelJoin": splitCammelJoin,
		}

		tpl := template.Must(template.New("test").Funcs(funcMap).Parse(templateData))
		var output bytes.Buffer
		if err := tpl.Execute(&output, nil); err != nil {
			t.Fatalf("%v", err)
		}
		if output.String() != expected {
			t.Fatalf("Expected '%s' but got '%s'", expected, output.String())
		}
	}
}
