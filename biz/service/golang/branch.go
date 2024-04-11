package golang

import "github.com/dongzwhitsz/llm_test/biz/model"

type BranchAnalyzer struct {
	Config model.Config
}

func (c *BranchAnalyzer) AnalyzeDiff() (model.CiOutput, error) {
	var result model.CiOutput
	return result, nil
}

func (c *BranchAnalyzer) AnalyzeRepo() (model.CiOutput, error) {
	var result model.CiOutput
	return result, nil
}
