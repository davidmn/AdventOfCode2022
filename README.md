# Advent of Code 2022

I've never finished an Advent of Code before. Maybe I will this year. Doing it in Go this time. I am not a Go developer. Nothing here is best practice, if it is it probably wasn't intentional.

## Cool things I've found along the way

You can throw away the errors with `_` if you want to fly by the seat of your pants

```go
file, _ := os.Open(path)
```

Running all tests in your module

```zsh
go test ./...
```

Which is nice to alias with in your .zshrc

```zsh
alias gt="go test ./..."
```
