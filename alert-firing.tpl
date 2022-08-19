{
  "toPersonId": "{{ .RoomId }}",
  "markdown": "# 🚨 Alert _{{ .Markdown.Alertname }}_ is {{ .Markdown.Status }}\n Alert: {{ .Markdown.Alertname }}\n Severity: {{ .Markdown.Severity}}\n Cluster: {{ .Markdown.Cluster }}\n Namespace: {{ .Markdown.Namespace }}\n Impact: {{ .Markdown.Impact }}\n##### Annotations\n* cluster: {{ .Markdown.Cluster }}\n* namespace: {{ .Markdown.Namespace }}\n* message: {{ .Markdown.Message }}\n ##### Handlungsanweisung \n - {{ .Markdown.Guideline }}\n ---" 
}
