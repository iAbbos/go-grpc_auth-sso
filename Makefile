start:
	 CONFIG_PATH=config/local.yaml go run cmd/sso/main.go

 migration:
	go run ./cmd/migrator/ --storage-path=./storage/sso.db --migrations-path=./migrations

 migration-test:
	go run ./cmd/migrator/migrator.go --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test