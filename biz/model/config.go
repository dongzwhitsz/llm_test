package model

type Language string

const (
	GoLang   Language = "go"
	JavaLang Language = "java"
	FELang   Language = "fe"
)

type Config struct {
	Scene     string                  `json:"Scene"`
	Increment bool                    `json:"Increment"`
	Language  string                  `json:"Language"`
	AddLines  map[string]map[int]bool `json:"AddLines"` // 增量行
	Region    string                  `json:"Region"`   // 网络区域
	Type      CoverType               `json:"Type"`
	RepoPath  string                  `json:"RepoPath"`

	// RawCovFile 接受codecov或地层工具的原始覆盖率数据文件。
	// 设置空字符串，Analyzer行为依赖具体的实现
	RawCovFile string `json:"RawCovFile"`
}
