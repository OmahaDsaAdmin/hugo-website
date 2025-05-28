---
title: "{{ .Title }}"
relPermalink: "{{ index .Identifiers 0 }}"
date: "{{ .StartDate.Format `2006-01-02` }}"
draft: false
hideToc: true
---

{{ if .ImageLink }} {{`{{`}}< smallImg src="{{ .ImageLink}}" alt="event image">{{`}}`}} {{ end }}
**EVENT START: {{ .StartDate.Format "Mon, 02 Jan 2006 3:04PM" }}**

{{ .Description }}

{{ if .EventLink }}[View Event and RSVP Here]({{ .EventLink }}){{ end }}
