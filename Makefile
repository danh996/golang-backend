sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/danh996/go-school/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
