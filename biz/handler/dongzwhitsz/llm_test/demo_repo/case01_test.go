package demo_repo

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/dongzwhitsz/llm_test/biz/model"
	"reflect"
	"testing"
)

func TestInitAnalyzers(t *testing.T) {
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
			name: "测试branch配置",
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
