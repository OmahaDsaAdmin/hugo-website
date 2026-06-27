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

NOTE: Blog posts should no longer be directly created in this repo (though they are committed / cached here!), instead they should be created in the remote custom API.

## How this works

1. Chapter members edit and publish posts in the web app.
2. Every night at 11:00 PM CST, the web app pulls any new action network events in too.
3. Every morning at 6:00 AM CST, this Hugo Site triggers a pull events GH action that stores all the posts as .md files in this repo, under `/content/events` and `/content/posts` depending on if it came from action network.
4. After that GH action is triggered, another deploy GH action is triggered which publishes the latest version of the hugo site.

Therefore, if a chapter member publishes a post today, the earliest they should expect to see it on the website is tomorrow.
