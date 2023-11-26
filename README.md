# TP3

[![CI: Test](https://github.com/BaptisteBuvron/go-voting/actions/workflows/test.yml/badge.svg)](https://github.com/BaptisteBuvron/go-voting/actions/workflows/test.yml)

Subject:

- [TP3](https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/td3/sujet.md)
- [Serveur de vote](https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/activit%C3%A9s/serveur-vote/api.md)

Run the server:

```bash
go run ia04/cmd/server
```

Example commands for client:

```bash
go run ia04/cmd/client v1 new-ballot majority '2023-11-26T16:27:11+00:00' 'v1,v2,v3' 5 '0,1,2,3,4'
go run ia04/cmd/client v1 vote majority-18c0c24a3245e '0,1,2,3,4'
go run ia04/cmd/client v1 result majority-18c0c24a3245e
```

Run tests:

```bash
go test '-coverprofile=coverage.txt' -v ./...
go tool cover '-html=coverage.txt'
```
