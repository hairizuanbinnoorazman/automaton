# {{.Metadata.Name}}

{{.Metadata.Description}}

It is vital to ensure that there is such a profile available.

There are {{.Result.ProfileCount}} profile/s that is available for this account.
{{if .Result.UnfilteredProfileAvailable}}
There is at least 1 profile that has no filters attached to it
{{else}}
There is no profile that has no filter attached to it. It is highly recommended for you to have at least 1 profile view with no filters attached to it.
{{end}}