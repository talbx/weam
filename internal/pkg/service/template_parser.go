package service

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/talbx/weam/internal/pkg/util"
)

type TemplateParser struct{}

func (parser TemplateParser) parseError(toBeParsed util.Webex) bytes.Buffer {
	return parser.parse(toBeParsed, "alert-firing.tpl")
}

func (parser TemplateParser) parseResolved(toBeParsed util.Webex) bytes.Buffer {
	return parser.parse(toBeParsed, "alert-resolved.tpl")
}

func (parser TemplateParser) parse(toBeParsed util.Webex, templateFile string) bytes.Buffer {
	var parsed bytes.Buffer
	tpl, err := template.New("webex-template").ParseFiles(templateFile)

	if err != nil {
		fmt.Println(err)
	}
	tpl.ExecuteTemplate(&parsed, templateFile, toBeParsed)
	return parsed
}
