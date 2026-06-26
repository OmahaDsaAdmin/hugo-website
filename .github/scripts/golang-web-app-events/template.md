---
title: "{{ .Title }}"
{{ if .ActionNetworkGuid }}relPermalink: "{{ .ActionNetworkGuid }}"{{ end }}
date: "{{ .EventStartDate.Format `2006-01-02` }}"
draft: {{ not .IsPublished }}
{{ if .AuthorPublicName }}author: "{{ .AuthorPublicName }}"{{ end }}
hideToc: true
---

{{ if .EventImageLink }} {{`{{`}}< smallImg src="{{ .EventImageLink}}" alt="event image">{{`}}`}} {{ end }}
**EVENT START: {{ .EventStartDate.Format "Mon, 02 Jan 2006 3:04PM" }}**

{{ .Description }}

{{ if .EventLink }}[View Event and RSVP Here]({{ .EventLink }}){{ end }}
