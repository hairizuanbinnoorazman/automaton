# {{.Metadata.Name}}

{{.Metadata.Description}}

Goals may be of a more optional thing but it can help your analyst team to track metrics that may be of concern to you or your business.

{{if gt .Result.GoalCount 0}} We observe that that are {{.Results.GoalCount}} goal/s in your profile. 
{{else}}
We observe that there are no goals in your profile. You should consider adding them to be able to monitor your key metrics on your digital platform more closely.
{{end}}