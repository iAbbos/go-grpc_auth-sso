start:
	 CONFIG_PATH=config/local.yaml go run cmd/ssh/main.go

 migration:
	go run ./cmd/migrator/ --storage-path=./storage/sso.db --migrations-path=./migrations