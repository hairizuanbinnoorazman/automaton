# {{ .Name }}

{{ .Description }}

{{if gt .CustomMetricCount 0}}We observe that you have {{.CustomMetricCount}} Custom Metrics in use
{{else}}We observe that you have no custom metrics in use.
{{end}}
