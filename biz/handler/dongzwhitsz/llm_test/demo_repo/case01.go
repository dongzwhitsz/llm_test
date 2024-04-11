package demo_repo

import (
	"fmt"
	"github.com/dongzwhitsz/llm_test/biz/model"
	"github.com/dongzwhitsz/llm_test/biz/service/golang"
)

// 指定当前的单测函数的输入config.Language是 model.GoLang或者model.JavaLang,config.Type="condition", 请根据这个条件输出对应的单元测试；
func InitAnalyzers(config model.Config) (model.CoverageAnalyzer, error) {
	var choseAnalyzer model.CoverageAnalyzer
	switch config.Language {
	case string(model.GoLang):
		if string(config.Type) == "branch" {
			choseAnalyzer = &golang.BranchAnalyzer{Config: config}
		} else if string(config.Type) == "condition" {
			choseAnalyzer = &golang.ConditionAnalyzer{Config: config}
		}
	default:
		choseAnalyzer = nil
	}
	if choseAnalyzer == nil {
		return choseAnalyzer, fmt.Errorf("did not find supported analyzers for indicator %s, language %s",
			string(config.Type), config.Language)
	}
	return choseAnalyzer, nil
}
