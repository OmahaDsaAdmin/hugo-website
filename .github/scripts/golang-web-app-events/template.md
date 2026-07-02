---
title: "{{ .Title }}"
{{ if .ActionNetworkGuid }}relPermalink: "{{ .ActionNetworkGuid }}"{{ end }}
{{ if .ActionNetworkGuid }}date: "{{ .EventStartDate.Format `2006-01-02` }}"{{ end }}
{{ if not .ActionNetworkGuid }}date: "{{ .CreatedAt.Format `2006-01-02` }}"{{ end }}
draft: {{ not .IsPublished }}
{{ if .AuthorPublicName }}author: "{{ .AuthorPublicName }}"{{ end }}
hideToc: true
---

{{ if .EventImageLink }} {{`{{`}}< smallImg src="{{ .EventImageLink}}" alt="event image">{{`}}`}} {{ end }}
{{ if .ActionNetworkGuid }}**EVENT START: {{ .EventStartDate.Format "Mon, 02 Jan 2006 3:04PM" }}**{{ end }}

{{ .Description }}

{{ if .EventLink }}[View Event and RSVP Here]({{ .EventLink }}){{ end }}
