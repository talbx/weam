package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/talbx/weam/internal/pkg/util"
)

type Notifier interface {
	Notify(payload util.Webex, conf *util.AppConfig) string
}

type WebexNotifier struct{}

func (notifier WebexNotifier) Notify(payload util.Webex, conf *util.AppConfig) string {
	payload.Markdown.AnnotationsString = util.MapToString(payload.Markdown.Annotations)
	payload.Markdown.LabelsString = util.MapToString(payload.Markdown.Labels)

	var parsed bytes.Buffer
	parsed = findParser(payload, parsed)

	request := prepareRequest(parsed, conf)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Panicln(err)
	}
	log.Println(response.Status)
	return response.Status
}

func findParser(payload util.Webex, parsed bytes.Buffer) bytes.Buffer {
	parser := TemplateParser{}
	if payload.Markdown.Status == "firing" {
		parsed = parser.parseError(payload)
	} else {
		parsed = parser.parseResolved(payload)
	}
	return parsed
}

func prepareRequest(parsed bytes.Buffer, conf *util.AppConfig) *http.Request {
	request, _ := http.NewRequest("POST", "https://webexapis.com/v1/messages", &parsed)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", conf.ApiKey))
	request.Header.Add("Content-Type", "application/json")
	return request
}
