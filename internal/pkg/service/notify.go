package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/talbx/weam/internal/pkg/util"
)

type Notifier interface {
	Notify(payload util.Webex) string
}

type WebexNotifier struct{}

func (notifier WebexNotifier) Notify(payload util.Webex, conf *util.AppConfig) string {

	parser := TemplateParser{}
	var parsed bytes.Buffer
	if payload.Markdown.Status == "firing" {
		parsed = parser.parseError(payload)
	} else {
		parsed = parser.parseResolved(payload)
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", "https://webexapis.com/v1/messages", &parsed)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", conf.ApiKey))
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Panicln(err)
	}

	res, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(res))

	response, err := client.Do(request)

	if err != nil {
		log.Panicln(err)
	}

	log.Println(response.Body)
	log.Println(response.StatusCode)

	log.Println(response.Status)
	return response.Status
}
