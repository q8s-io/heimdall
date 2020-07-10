package model

type ClairScanResult struct {
	LayerCount      int             `json:"LayerCount"`
	Vulnerabilities Vulnerabilities `json:"Vulnerabilities"`
}

type Vulnerabilities struct {
	Unknown    []Grade `json:"Unknown"`
	Negligible []Grade `json:"Negligible"`
	Low        []Grade `json:"Low"`
	Medium     []Grade `json:"Medium"`
	High       []Grade `json:"High"`
	Critical   []Grade `json:"Critical"`
	Defcon1    []Grade `json:"Defcon1"`
}

type Grade struct {
	Name           string `json:"Name"`
	NamespaceName  string `json:"NamespaceName"`
	Description    string `json:"Description"`
	Link           string `json:"Link"`
	Severity       string `json:"Severity"`
	FeatureName    string `json:"FeatureName"`
	FeatureVersion string `json:"FeatureVersion"`
}
