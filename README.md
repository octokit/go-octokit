# go-octokit

Go toolkit for the GitHub API.

# Hypermedia-driven client

## Show a user

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
    client := octokit.NewClient()

    user, result := client.User("jingweno") // Internally it's hypermedia-driven
    if result.HasError() {
      // Handle error
    }
    // Do something with user
}
```

or

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
    client := octokit.NewClient()

    // Get root
    root, result := client.Root()
    if result.HasError() {
      // Handle error
    }

    // Get a user
    userURL, _ := root.UserURL.Expand(octokit.M{"user": "jingweno"})
    requester := client.Requester(userURL)
    user := new(User)
    result = requester.Get(user)
    if result.HasError() {
      // Handle error
    }
    // Do something with user
}
```

## Pagination

```go
package main

import "github.com/octokit/go-octokit/octokit"

func main() {
    client := octokit.NewClient()
    repos, result := client.OrgRepos("github")
    if result.HasError() {
      // Handle error
    }
    // Do something with repos

    // next page
    result = client.Requester(result.NextPage).Get(&repos)
    if result.HasError() {
      // Handle error
    }
    // Do something with users
}

```

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
