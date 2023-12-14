# Anchors Server

1) Have a postgres database named `anchors` with user `postgres` on host
`localhost` bound to port `5432`.
2) Run `run go main.go` and you'll have a simple
API that exposes the endpoints at `localhost:8080`:
  - GET /tours
  - GET /tour/:id
  - GET /chapters
