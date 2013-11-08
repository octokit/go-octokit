# go-octokit

Go toolkit for the GitHub API.

## Status

Very experimental.

## Motivation

`go-octokit` is designed to be a hypermedia API client that [wraps](http://wynnnetherland.com/journal/what-makes-a-good-api-wrapper) the [GitHub API](http://developer.github.com/).

## Hypermedia agent

`go-octokit` is hypermedia-driven by default.
Under the hood, it uses [`go-sawyer`](https://github.com/lostisland/go-sawyer), the Go version of [Ruby Sawyer](https://github.com/lostisland/sawyer).

### Hypermedia in go-octokit

Resources in `go-octokit` contain not only data but hypermedia links:

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
    client := octokit.NewClient(nil)
    usersService, err := client.Users(&octokit.UserURL, octokit.M{"user": "jingweno"})
    if err != nil  {
      // Handle error
    }

    user, result := usersService.Get()
    if result.HasError() {
      // Handle error
    }

    fmt.Println(user.ReposURL) // https://api.github.com/users/jingweno/repos
}
```

### URI templates

Many hypermedia links have variable placeholders. `go-octokit` supports [URI Templates](http://tools.ietf.org/html/rfc6570) for parameterized URI expansion:

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
    url, _ := octokit.UserURL.Expand(octokit.M{"user": "jingweno"})
    fmt.Println(url) // https://api.github.com/users/jingweno
}
```

### The Full Hypermedia Experience™

If you want to use `go-octokit` as a pure hypermedia API client, you can
start at the API root and follow hypermedia links which drive the application state transitions:

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
  rootService, _ := client.Root(nil)
  root, _ := rootService.Get()

  usersService, _ := client.Users(root.UserURL, octokit.M{"users": "jingweno"})
  user, _ := usersService.Get()
}
```

### Pagination

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
    client := octokit.NewClient(nil)
    usersService, err := client.Users(&octokit.UserURL, nil)
    if err != nil  {
      // Handle error
    }

    users, result := usersService.GetAll()
    if result.HasError() {
      // Handle error
    }

    // Do something with users

    // Next page
    usersService, err := client.Users(result.NextPage, nil)
    if result.HasError() {
      // Handle error
    }

    // Do something with users
}

```

### Caching

Client-side caching is the #1 thing to do to make a hypermedia client more performant.
We plan to support this in the near future.

More [examples](https://github.com/octokit/go-octokit/blob/master/examples/example.go) are available.

## Release Notes

See [Releases](https://github.com/octokit/go-octokit/releases).

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

go-octokit is released under the MIT license. See
[LICENSE.md](https://github.com/octokit/go-octokit/blob/master/LICENSE.md).
