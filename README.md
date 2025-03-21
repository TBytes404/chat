### Chat Room using Go

TCP Server using [stdlib net](https://pkg.go.dev/net), can broadcast messages.

## Quick Start

Set program alias.

```sh
alias chat="go run ."
```

Create new Chat server on `localhost:8000`.

```sh
chat serve
```

Connect to chat server on `localhost:8000`.

```sh
chat connect
```

Type message and Enter to send. It will broadcast to everyone connected.
