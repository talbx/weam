package service

import (
	"fmt"

	"github.com/talbx/weam/internal/pkg/util"
)

type PayloadMapper struct{}

func (mapper PayloadMapper) Apply(payload util.Payload) []util.Markdown {
	len := len(payload.Alerts)
	var sl = make([]util.Markdown, len)
	for index := range payload.Alerts {
		fmt.Println(payload.Alerts[index])
		fmt.Println(payload.Alerts[index].Labels["alertname"])

		alert := payload.Alerts[index]

		var md = util.Markdown{
			Status:    payload.Status,
			Severity:  alert.Labels["severity"],
			Alertname: alert.Labels["alertname"],
			Guideline: alert.Annotations["guideline"],
			Namespace: alert.Labels["namespace"],
			Cluster:   alert.Labels["cluster_id"],
			Impact:    alert.Labels["impact"],
			Message:   alert.Annotations["message"],
			Occurence: alert.StartsAt,
			Labels:    alert.Labels,
		}
		sl = append(sl, md)
	}

	return sl
}
