# Octokat [![Build Status](https://drone.io/github.com/jingweno/octokat/status.png)](https://drone.io/github.com/jingweno/octokat/latest)

Simple Go wrapper for the GitHub API. It's extracted from [my other project](https://github.com/jingweno/gh). The API is inspired by [octokit.rb](https://github.com/octokit/octokit.rb).

# GoDoc

[http://godoc.org/github.com/jingweno/octokat](http://godoc.org/github.com/jingweno/octokat)

# Example

## Show a user

```go
package main

import "github.com/jingweno/octokat"

func main() {
    client := octokat.NewClient()
    user, err := client.User("jingweno")
    // Do something with user
}
```

## List authorizations

```go
package main

import "github.com/jingweno/octokat"

func main() {
    client := octokat.NewClient().WithLogin("LOGIN", "PASSWORD")
    authorizations, err := client.Authorizations()
    // Do something with authorizations
}
```

## Create a pull request

```go
package main

import "github.com/jingweno/octokat"

func main() {
    client := octokat.NewClient().WithToken("OAUTH_TOKEN")
    repo := octokat.Repo{Name: "octokat", UserName: "jingweno"}
    params := octokat.PullRequestParams{Base: "master", Head: "feature", Title: "A pull request", Body: "A body"}
    pullRequest, err := client.CreatePullRequest(repo, params)
    // Do something with pullRequest
}
```

## Release Notes

* **0.2.0** (in progress)
  * Implement [GET repos](http://developer.github.com/v3/repos/#get)
* **0.1.0** June 8, 2013
  * Extract `octokat` from [`gh`](https://github.com/jingweno/gh)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

octokat is released under the MIT license. See LICENSE.md.
