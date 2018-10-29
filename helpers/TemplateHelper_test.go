package helpers

import (
	"testing"
)

func Test_templateToHTML(t *testing.T) {
	type args struct {
		tplFile  string
		tmplName string
		obj      interface{}
	}

	type CFoo struct {
		Name string
		Age  int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				tplFile:  "/Users/derrick.zhu/Documents/go/src/FFQATracking/tests/templateToHTML.tpl",
				tmplName: "templateHelper",
				obj:      CFoo{Name: "John", Age: 16},
			},
			want: "\nHi, John welcome!\n",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TemplateToHTML(tt.args.tplFile, tt.args.tmplName, tt.args.obj); got != tt.want {
				t.Errorf("templateToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
