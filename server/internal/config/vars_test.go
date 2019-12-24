package config

import (
	"reflect"
	"testing"
)

func Test_filesFromFolder(t *testing.T) {
	tests := []struct {
		name      string
		root      string
		wantFiles []string
		wantErr   bool
	}{
		{"test config folder", ".", []string{"configuration.go", "vars.go", "vars_test.go"}, false},
		{"test error", "./trustwallet/error/test", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := filesFromFolder(tt.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("filesFromFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("filesFromFolder() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

// REPLACE TEST
func Test_replaceVars(t *testing.T) {
	type args struct {
		path string
		old  string
		new  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test replace", args{"vars_test.go", "REPLACE TEST", "REPLACE TEST"}, false},
		{"test error", args{"./trustwallet/error/test", "", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := replaceVars(tt.args.path, tt.args.old, tt.args.new); (err != nil) != tt.wantErr {
				t.Errorf("replaceVars() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
