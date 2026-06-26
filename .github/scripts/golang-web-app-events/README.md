# Golang Web App Events

Generate markdown files from remote custom API

1. Export Path
2. API Url
3. API Key

## How to run

1. `go build`
2. `./golang-web-app-events './content-path' 'http://localhost:8090/api/export-unsafe-posts' 'apikey'`

Depending on your permissioning, you may need to use `sudo`

This will create a .md file for every post. Those files are added to either the `/events` or `/posts` folder, and the site will automatically include them when a new build is scheduled.

NOTE: Blog posts should no longer be directly created in this directory, instead they should be created in the remote custom API
