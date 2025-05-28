# Hugo Website Events

Generate markdown files from Action Network events given the following:

1. Action Network API Key
2. Export Path

## How to run

1. `go build`
2. `./hugo-website-events apikey ./exports`

Depending on your permissioning, you may need to use `sudo`

This will create a .md file for every Action Network event. Those files can be added to the hugo website `/events` folder, and the site will automatically include them.

NOTE: If a file exists in the export path with the same AN identifier name, this program will not overwrite it.

NOTE: This was copied from a standalone repo here: [link](https://github.com/OmahaDsaAdmin/hugo-website-events)
