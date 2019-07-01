# Github_Fetch
Github_Fetch is a library for fetching the contributions of a user of Github in the last year. <a href="https://help.github.com/en/articles/why-are-my-contributions-not-showing-up-on-my-profile">Contributions</a> that are counted include issues opened by a user, commits made on a master repository, and pull requests opened on a master repository. 

## Installation
Github_Fetch requires Go version 1.9 or later

## Usage
import "https://github.com/meganabyte/Github_Fetch/github" ...

### Authentication
Create an OAuth2 Access token (for example, a personal API token) and store it as an environment variable. The Github_Fetch library will directly handle authentication for you.

## Acknowledgments
https://github.com/google/go-github

