# go-octokit

Go toolkit for the GitHub API.

## Hypermedia-driven client

### Show a user

```go
package main

import "github.com/octokit/go-octokit"

func main() {
    client := octokit.NewClient()
    usersService, err := client.Users(&octokit.UsersHyperlink, octokit.M{"user": "jingweno"})
    if err != nil  {
      // Handle error
    }

    user, result := usersService.Get()
    if result.HasError() {
      // Handle error
    }

    // Do something with user
}
```

### Pagination

```go
package main

import "github.com/octokit/go-octokit"

func main() {
    client := octokit.NewClient()
    usersService, err := client.Users(&octokit.AllUsersHyperlink, nil)
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

### Exploring hypermedia APIs

```go
rootService, _ := client.Root(nil)
root, _ := rootService.Get()

usersService, _ := client.Users(root.UserURL, octokit.M{"users": "jingweno"})
user, _ := usersService.Get()
```

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
