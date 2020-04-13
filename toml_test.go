package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_parseMetadata(t *testing.T) {

	f, _ := ioutil.TempFile("", "")
	f.WriteString(`<!-- 
					title="this is a title"
					date="2020-04-11 16:22"
					thumbnail=["xxxx"]
					summary="this a summary"
					category="test"
					tags=["xxxii"]
					-->
					# Title
					> this is markdown package tips
					
					
					In this package, U will see how to use gopkg.in/russross/blackfriday.v2 parse markdown and generate html code.`)
	f.Close()
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantMd  metadata
		wantErr bool
	}{
		{
			name: "parse metadata",
			args: struct{ path string }{path: f.Name()},
			wantMd: metadata{
				Title:     "this is a title",
				Date:      "2020-04-11 16:22",
				Thumbnail: []string{"xxxx"},
				Summary:   "this a summary",
				Category:  "test",
				Tags:      []string{"xxxii"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMd, err := parseMetadata(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMd, tt.wantMd) {
				t.Errorf("parseMetadata() gotMd = %v, want %v", gotMd, tt.wantMd)
			}
		})
	}
}
