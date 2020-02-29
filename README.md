# go-strawpoll

go-strawpoll provides an easy way to fetch and create polls from/tostrawpoll.me.

## Getting Started

**Add go-strawpoll to your dependencies:**
```
go get -u github.com/jozsefsallai/go-strawpoll
```

**To fetch a poll:**

```go
func main() {
  poll, err := strawpoll.Get(1)
  if err != nil {
    panic(err)
  }
  fmt.Println(poll.Title)
}
```

(where `1` is the ID of the poll on strawpoll.me)

**To create a poll:**

```go
func main() {
  poll, err := strawpoll.Create(
    "title of the poll",
    []string{"option 1", "option 2", "option 3"}, // at least 2
    false, // multi-choice
    strawpoll.DupcheckNormal, // duplication checking level
    false, // require CAPTCHA
  )

  if err != nil {
    panic(err)
  }

  fmt.Println(poll.ID)
}
```

## Documentation

https://godoc.org/github.com/jozsefsallai/go-strawpoll

## License

MIT.
