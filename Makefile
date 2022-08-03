.PHONY: test clean run-prod

test:
	 go test ./...

clean:
	git clean -fXd

run-prod:
	GIN_MODE=release PORT=8080 go run cmd/server/main.go

