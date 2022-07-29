{
  "toPersonId": "{{ .RoomId }}",
  "markdown": "# ✅ Prometheus Alert `{{ .Markdown.Alertname }}` is resolved!\n Alert: {{ .Markdown.Alertname }}\n Severity: {{ .Markdown.Severity}}\n Cluster: {{ .Markdown.Cluster }}\n Namespace: {{ .Markdown.Namespace }}\n Impact: {{ .Markdown.Impact }}\n* cluster: {{ .Markdown.Cluster }}\n* namespace: {{ .Markdown.Namespace }}\n* message: {{ .Markdown.Message }}\n* handlungsanweisung: {{ .Markdown.Guideline }}" 
}
