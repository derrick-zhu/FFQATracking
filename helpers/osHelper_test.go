package helpers

import "testing"

func TestAbosolutePath(t *testing.T) {
	type args struct {
		relateFilePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				relateFilePath: "views/demo.tpl",
			},
			want: GetCurrentDirectory() + "/helpers.test/views/demo.tpl",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbosolutePath(tt.args.relateFilePath); got != tt.want {
				t.Errorf("AbosolutePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
