package model

type CoverageAnalyzer interface {
	// AnalyzeDiff 增量分析
	AnalyzeDiff() (CiOutput, error)
	// AnalyzeRepo 全量分析
	AnalyzeRepo() (CiOutput, error)
}
