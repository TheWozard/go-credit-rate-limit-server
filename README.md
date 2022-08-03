# go-credit-rate-limit-server
[![Test](https://github.com/TheWozard/go-credit-rate-limit-server/actions/workflows/test.yml/badge.svg)](https://github.com/TheWozard/go-credit-rate-limit-server/actions/workflows/test.yml)

A simple api for coordinating rate limiting through an credit based api. Also an exploration into [Gin](https://github.com/gin-gonic/gin)

# Run Locally

```
$ make run-prod
```

Starts the server on `localhost:8080`

## Endpoints
|Path|Query Params|Description|
|:-|:-|:-|
| `/health` | N/A | Basic health endpoint. |
| `/v1/acquire` | account=string,credits=int | Requests a certain number of credits from the account. Returns the number of credits received as `{"credits":<credits>}` |

# Adding a new account

In `cmd\server\main.go` add a new entry to the `Limiter` map `"example": rate.NewRate(5*time.Second, 15)`
