{{ if .Versions -}}
# Changelog

Todos los cambios notables de este proyecto serán documentados en este archivo de acuerdo a los estándares de [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/) y [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

{{ range .Versions }}
## {{ if .Tag.Previous }}[{{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}){{ else }}{{ .Tag.Name }}{{ end }} — *({{ if .Tag.Date }}{{ datetime "2006-01-02" .Tag.Date }}{{ else }}En desarrollo / Unreleased{{ end }})*

{{ range .CommitGroups -}}
### {{ .Title }}

{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }} *({{ datetime "2006-01-02" .Author.Date }})* ([{{ .Hash.Short }}]({{ $.Info.RepositoryURL }}/commit/{{ .Hash.Long }}))
{{ end }}
{{ end -}}

{{- if .RevertCommits -}}
### ⏪ Reverts
{{ range .RevertCommits -}}
- {{ .Revert.Header }} ([{{ .Hash.Short }}]({{ $.Info.RepositoryURL }}/commit/{{ .Hash.Long }}))
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### ⚠️ {{ .Title }}
{{ range .Notes }}
- {{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ else -}}
No se han registrado versiones en el repositorio todavía.
{{ end -}}