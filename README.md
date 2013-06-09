# Ocotokit [![Build Status](https://drone.io/github.com/jingweno/octokit/status.png)](https://drone.io/github.com/jingweno/octokit/latest)

Simple Go wrapper for the GitHub API.

# GoDoc

[http://godoc.org/github.com/jingweno/octokit](http://godoc.org/github.com/jingweno/octokit)

# Example

## Show a user

```go
package main

import "github.com/jingweno/octokit"

func main() {
    client := octokit.NewClient()
    user, err := client.User("jingweno")
    // Do something with user
}
```

## List authorizations

```go
package main

import "github.com/jingweno/octokit"

func main() {
    client := octokit.NewClientWithPassword("LOGIN", "PASSWORD")
    authorizations, err := client.Authorizations()
    // Do something with authorizations
}
```

## Create a pull request

```go
package main

import "github.com/jingweno/octokit"

func main() {
    client := octokit.NewClientWithToken("OAUTH_TOKEN")
    repository := octokit.Repository{Name: "octokit", UserName: "jingweno"}
    params := octokit.PullRequestParams{Base: "master", Head: "feature", Title: "A pull request", Body: "A body"}
    pullRequest, err := client.CreatePullRequest(repository, params)
    // Do something with pullRequest
}
```

## Release Notes

* **0.1.0** June 8, 2013
  * Extract `octokit` from [`gh`](https://github.com/jingweno/gh)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

octokit is released under the MIT license. See LICENSE.md.
