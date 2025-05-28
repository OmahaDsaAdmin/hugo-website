# Hugo Website

This is a simple static site built with [Hugo](https://gohugo.io/) and a theme called [Poison](https://github.com/lukeorth/poison?tab=readme-ov-file#readme) for the Omaha Democratic Socialists of America chapter.

## How to update the website

Here are some key concepts to understand:

1. Learn [git](https://git-scm.com/doc) (`push`, `pull`, and `branch`) to update the website repo.
2. Learn [markdown](https://www.markdownguide.org/cheat-sheet/) to edit posts and pages of the site under `./content` folder.
3. Learn [css](https://www.w3schools.com/css/default.asp) to edit the styling of the website in `./assets/css/custom.css` these will modify all styles sitewide.
4. To run the site locally, download Go, download Hugo, `git clone` the repo, and use the following terminal command `hugo server`. To stop the local server, use `ctrl+c` in the terminal.
5. You can also directly modify files through Github, but you won't be able to preview your changes.

### Publish Dates

You may notice that running `hugo server` by itself doesn't render posts or events in the future. To get around this default behavior, use `hugo server --buildFuture` instead.

## Useful tools and info

1. [Visual Studio Code](https://code.visualstudio.com/)
2. [Markdown Editor](https://stackedit.io/app#)
3. [Poison docs](https://github.com/lukeorth/poison)
4. [DSA Styleguide:](https://design.dsausa.org/resources/)

## Recomendations

Run the following commands to sign commits as the admin user:

`git config user.name OmahaDsaAdmin`
`git config user.email omahadsa@gmail.com`
