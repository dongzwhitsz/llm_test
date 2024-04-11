package model

type CoverType string

const (
	LineCoverage      CoverType = "line"
	BranchCoverage    CoverType = "branch"
	ConditionCoverage CoverType = "condition"
	// TODO: other
)

// CoverUnit 对应一个基本的覆盖代码单元，如分支或条件
type CoverUnit struct {
	Type       CoverType `json:"Type"` // condition or branch or ...
	SourceCode string    `json:"SourceCode"`
	TrueCount  int       `json:"TrueCount"`
	FalseCount int       `json:"FalseCount"`
	FilePath   string    `json:"FilePath"`
	StartRow   int       `json:"StartRow"`
	EndRow     int       `json:"EndRow"`
	StartCol   int       `json:"StartCol"`
	EndCol     int       `json:"EndCol"`
	CreateTime int64     `json:"CreateTime"`
	Owner      string    `json:"owner"`
}

// Metric 取值: line、branch、condition、statement
type Metric struct {
	TotalCount int64   `json:"TotalCount"`
	HitCount   int64   `json:"HitCount"`
	HitRatio   float64 `json:"HitRatio"`
}

type FileCovUnit struct {
	Metrics map[string]Metric      `json:"Metrics"`
	RawCov  map[string][]CoverUnit `json:"RawCov"`
}

type CiOutput struct {
	Metrics      map[string]Metric      `json:"Metrics"`
	FileCov      map[string]FileCovUnit `json:"FileCov"`
	ErrorMsgList []string               `json:"ErrorMsgList"`
	HTMLPath     string                 `json:"HTMLPath"`
}
