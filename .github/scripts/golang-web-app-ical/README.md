# Notes

- Scheduled to run at 6:10 every morning.
- This program pulls all the events from our backend, checks if they have an Action Network Guid, and if they do, includes them in the template.ics
- We could improve this by adding a location field in the backend, storing end times from Action Network, and only including them if they have an IsEvent field set to true.

## How to run

1. `go build`
2. `./golang-web-app-ical './export-dir' 'http://localhost:8090/api/export-unsafe-posts' 'apikey'`
