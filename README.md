# Octokat [![Build Status](https://drone.io/github.com/jingweno/octokat/status.png)](https://drone.io/github.com/jingweno/octokat/latest)

Simple Go wrapper for the GitHub API. It's written by following the implementation of the official [Ruby wrapper](https://github.com/octokat/octokat.rb).

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
    repository := octokat.Repository{Name: "octokat", UserName: "jingweno"}
    params := octokat.PullRequestParams{Base: "master", Head: "feature", Title: "A pull request", Body: "A body"}
    pullRequest, err := client.CreatePullRequest(repository, params)
    // Do something with pullRequest
}
```

## Release Notes

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
