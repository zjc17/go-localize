package main

import (
	"reflect"
	"testing"
)

func Test_generateLocalizations(t *testing.T) {
	type args struct {
		files []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				files: []string{
					"mock/dir/sub/valid_json.json",
					"mock/dir/valid_json.json",
					"mock/dir/valid_yaml.yaml",
				},
			},
			want: map[string]string{
				"mock.dir.sub.valid_json.test": "test",
				"mock.dir.valid_json.test":     "test",
				"mock.dir.valid_yaml.test":     "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateLocalizations(tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateLocalizations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateLocalizations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateFile(t *testing.T) {
	type args struct {
		output       string
		translations map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				output:       "test_files",
				translations: map[string]string{"hello": "one"},
			},
		},
		{
			name: "invalid dir",
			args: args{
				output:       "",
				translations: map[string]string{"hello": "one"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateFile(tt.args.output, tt.args.translations); (err != nil) != tt.wantErr {
				t.Errorf("generateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getLocalizationsFromFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "valid json",
			args: args{"mock/valid.json"},
			want: map[string]string{"mock.valid.test1": "test2"},
		},
		{
			name: "valid yaml",
			args: args{"mock/valid.yaml"},
			want: map[string]string{"mock.valid.test1": "test2"},
		},
		{
			name:    "file not exist",
			args:    args{"mock/non_exist.json"},
			wantErr: true,
		},
		{
			name:    "invalid json",
			args:    args{"mock/invalid.json"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocalizationsFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalizationsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocalizationsFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSlicePath(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "valid",
			args: args{"mock/valid.json"},
			want: []string{"mock", "valid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSlicePath(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSlicePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFlags(t *testing.T) {
	type args struct {
		input  *string
		output *string
	}

	dirBlank := ""
	dirOk := "input"

	tests := []struct {
		name      string
		args      args
		inputDir  string
		outputDir string
		wantErr   error
	}{
		{
			name: "valid",
			args: args{
				input:  &dirOk,
				output: &dirOk,
			},
			inputDir:  dirOk,
			outputDir: dirOk,
		},
		{
			name: "default output dir",
			args: args{
				input:  &dirOk,
				output: &dirBlank,
			},
			inputDir:  dirOk,
			outputDir: defaultOutputDir,
		},
		{
			name: "invalid input",
			args: args{
				input:  &dirBlank,
				output: &dirBlank,
			},
			wantErr: errFlagInputNotSet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputDir, outputDir, err := parseFlags(tt.args.input, tt.args.output)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("parseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if inputDir != tt.inputDir {
				t.Errorf("parseFlags() got = %v, want %v", inputDir, tt.inputDir)
			}
			if outputDir != tt.outputDir {
				t.Errorf("parseFlags() got1 = %v, want %v", outputDir, tt.outputDir)
			}
		})
	}
}

func Test_getLocalizationFiles(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "valid",
			args: args{"mock/dir"},
			want: []string{
				"mock/dir/sub/valid_json.json",
				"mock/dir/valid_json.json",
				"mock/dir/valid_yaml.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocalizationFiles(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalizationFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocalizationFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}