package demo_repo

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/dongzwhitsz/llm_test/biz/model"
	"github.com/dongzwhitsz/llm_test/biz/service/golang"
	"reflect"
	"testing"
)

func TestInitAnalyzers1(t *testing.T) {
	type args struct {
		config model.Config
	}
	tests := []struct {
		name    string
		args    args
		want    model.CoverageAnalyzer
		wantErr bool
	}{
		{
			name: "测试Java类型的ConditionCoverage",
			args: args{
				config: model.Config{
					Scene:      "A Scene",
					Increment:  false,
					Language:   string(model.JavaLang),
					AddLines:   nil,
					Region:     "boe",
					Type:       model.ConditionCoverage,
					RepoPath:   "",
					RawCovFile: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "测试Golang类型的ConditionCoverage",
			args: args{
				config: model.Config{
					Scene:      "Another Scene",
					Increment:  false,
					Language:   string(model.GoLang),
					AddLines:   nil,
					Region:     "boe",
					Type:       model.ConditionCoverage,
					RepoPath:   "",
					RawCovFile: "",
				},
			},
			want: &golang.ConditionAnalyzer{Config: model.Config{
				Scene:      "Another Scene",
				Increment:  false,
				Language:   string(model.GoLang),
				AddLines:   nil,
				Region:     "boe",
				Type:       model.ConditionCoverage,
				RepoPath:   "",
				RawCovFile: "",
			}},
			wantErr: false,
		},
		{
			name: "测试Golang类型的BranchCoverage",
			args: args{
				config: model.Config{
					Scene:      "Third Scene",
					Increment:  false,
					Language:   string(model.GoLang),
					AddLines:   nil,
					Region:     "boe",
					Type:       model.BranchCoverage,
					RepoPath:   "",
					RawCovFile: "",
				},
			},
			want: &golang.BranchAnalyzer{Config: model.Config{
				Scene:      "Third Scene",
				Increment:  false,
				Language:   string(model.GoLang),
				AddLines:   nil,
				Region:     "boe",
				Type:       model.BranchCoverage,
				RepoPath:   "",
				RawCovFile: "",
			}},
			wantErr: false,
		},
	}
	indent, _ := json.MarshalIndent(tests, "", "  ")
	t.Log(string(indent))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitAnalyzers(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitAnalyzers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitAnalyzers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
