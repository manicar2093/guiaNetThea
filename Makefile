migrateup:
	migrate -path db/migrations -database [url] -verbose up

migratedown:
	migrate -path db/migrations -database [url] -verbose down

.PHONY: migrateup migratedown