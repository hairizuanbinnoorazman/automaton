# {{ .Name }}

{{ .Description }}

Using custom metrics allow one to extrapolate web analytics beyond it usual dimensions and metrics. Custom metrics can be used to tie certain business level data to web analytics data, providing such data even more context, thereby making it much more useful.

{{ if .UsedCustomMetrics }}We observe that you have {{ .CustomMetricCount }} Custom Metrics in use
{{ else }}We observe that you have no custom metrics in use.
{{ end }}

<div style="page-break-after: always;"></div>
