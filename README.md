# Octokat [![Build Status](https://travis-ci.org/jingweno/octokat.png?branch=master)](https://travis-ci.org/jingweno/octokat)

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
    user, err := client.User("jingweno", nil)
    // Do something with user
}
```

## List authorizations

```go
package main

import "github.com/jingweno/octokat"

func main() {
    client := octokat.NewClient().WithLogin("LOGIN", "PASSWORD")
    authorizations, err := client.Authorizations(nil)
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
    options := octokat.Options{Params: params}
    pullRequest, err := client.CreatePullRequest(repo, &options)
    // Do something with pullRequest
}
```

## Release Notes

See [Releases](https://github.com/jingweno/octokat/releases).

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

octokat is released under the MIT license. See
[LICENSE.md](https://github.com/jingweno/octokat/blob/master/LICENSE.md).
