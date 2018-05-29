# {{ .Name }}

{{.Description}}

Using custom dimension allow one to extrapolate web analytics beyond it usual dimensions and metrics. Custom dimensions can be used to tie certain business level data to web analytics data, providing such data even more context, thereby making it much more useful.

{{if gt .CustomDimensionCount 0}}We observe that you have {{.CustomDimensionCount}} Custom Dimensions in use
{{else}}We observe that you have no custom dimensions in use.{{end}}

<div style="page-break-after: always;"></div>
