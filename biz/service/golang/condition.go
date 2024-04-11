package golang

import "github.com/dongzwhitsz/llm_test/biz/model"

type ConditionAnalyzer struct {
	Config model.Config
}

func (c *ConditionAnalyzer) AnalyzeDiff() (model.CiOutput, error) {
	result := model.CiOutput{
		Metrics:      make(map[string]model.Metric),
		FileCov:      make(map[string]model.FileCovUnit),
		ErrorMsgList: []string{},
		HTMLPath:     "",
	}
	return result, nil
}

func (c *ConditionAnalyzer) AnalyzeRepo() (model.CiOutput, error) {
	var result model.CiOutput
	return result, nil
}
