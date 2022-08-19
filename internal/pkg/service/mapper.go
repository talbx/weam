package service

import (
	"github.com/talbx/weam/internal/pkg/util"
)

type PayloadMapper struct{}

func (mapper PayloadMapper) Apply(payload util.Payload) []util.Markdown {
	len := len(payload.Alerts)
	var sl = make([]util.Markdown, len)
	for index := range payload.Alerts {
		alert := payload.Alerts[index]
		var md = util.Markdown{
			Status:      payload.Status,
			Severity:    alert.Labels["severity"],
			Alertname:   alert.Labels["alertname"],
			Guideline:   alert.Annotations["guideline"],
			Timestamp:   alert.StartsAt,
			Annotations: alert.Annotations,
			Labels:      alert.Labels,
		}
		sl = append(sl, md)
	}

	return sl
}
