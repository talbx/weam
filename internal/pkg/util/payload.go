package util

type Webex struct {
	RoomId   string
	Markdown Markdown
}

type Markdown struct {
	Status      string
	Alertname   string
	Cluster     string
	Namespace   string
	Guideline   string
	Caption     string
	Message     string
	Occurence   string
	Severity    string
	Impact      string
	Labels      map[string]string
	Annotations map[string]string
}

type Payload struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}

type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     string            `json:"startsAt"`
	EndsAt       string            `json:"endsAt"`
	GeneratorUrl string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}
