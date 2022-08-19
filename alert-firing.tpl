{
  "toPersonId": "{{ .RoomId }}",
  "markdown": "# 🚨 Alert _{{ .Markdown.Alertname }}_ is {{ .Markdown.Status }}\n Alert: {{ .Markdown.Alertname }}\n Severity: {{ .Markdown.Severity}}\n Timestamp: {{ .Markdown.Timestamp}}\n#### Labels\n {{ .Markdown.LabelsString }}\n ##### Annotations\n {{ .Markdown.AnnotationsString }} ##### Handlungsanweisung \n - {{ .Markdown.Guideline }}\n ---" 
}
