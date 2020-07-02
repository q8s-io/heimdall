package entity

type JobImageAnalyzer struct {
	TaskID      string
	JobID       string
	JobStatus   string
	ImageName   string
	ImageDigest string
	ImageLayers string
	CreateTime  string
	Active      int
}
