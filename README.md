# goru

goru is a small Go command-line tool that fetches a GitHub user's public repositories and displays the repository name, star count, and description in a tabular format.

## Requirements

- Go 1.26+

## Build

```bash
go build -o goru
```

## Run

```bash
go run .
```

or with the built binary:

```bash
./goru
```

When prompted, enter a GitHub username.

## Example

```bash
Enter GitHub username: github
```

## Notes

- The tool uses the GitHub public API and does not require authentication for public repositories.
- If the user has no public repositories, the tool prints a message and exits.
" 
