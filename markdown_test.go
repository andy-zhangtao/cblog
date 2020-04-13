package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test_parseMarkdown(t *testing.T) {

	f, _ := ioutil.TempFile("", "")
	f.WriteString(`# cblog`)
	f.Close()

	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantHtml string
		wantErr  bool
	}{
		{
			name:     "Parse sample markdown",
			args:     struct{ path string }{path: f.Name()},
			wantHtml: `<h1>cblog</h1>`,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHtml, err := parseMarkdown(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if strings.TrimSpace(gotHtml) != tt.wantHtml {
				t.Errorf("parseMarkdown() gotHtml = %v, want %v", gotHtml, tt.wantHtml)
			}
		})
	}
}
