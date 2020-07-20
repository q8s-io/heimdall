package model

type TrivyScanResult struct {
	Target          string          `json:"Target"`
	Type            string          `json:"Type"`
	Vulnerabilities []Vulnerability `json:"Vulnerabilities"`
}

type Vulnerability struct {
	VulnerabilityID  string        `json:"VulnerabilityID"`
	PkgName          string        `json:"PkgName"`
	InstalledVersion string        `json:"InstalledVersion"`
	Layer            Layer         `json:"Layer"`
	SeveritySource   string        `json:"SeveritySource"`
	Title            string        `json:"Title"`
	Description      string        `json:"Description"`
	Severity         string        `json:"Severity"`
	VendorVectors    VendorVectors `json:"VendorVectors"`
	References       []string      `json:"References"`
}

type Layer struct {
	DiffID string `json:"DiffID"`
}

type VendorVectors struct {
	Nvd    Level `json:"nvd"`
	Redhat Level `json:"redhat"`
}

type Level struct {
	V2 string `json:"v2"`
	V3 string `json:"v3"`
}
