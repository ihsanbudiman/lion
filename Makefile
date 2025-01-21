migrate-source-up:
	soda migrate -c db_source.yaml -p db_source_migrations

migrate-source-down:
	soda migrate -c db_source.yaml -p db_source_migrations down

migrate-destination-up:
	soda migrate -c db_destination.yaml -p db_destination_migrations

migrate-destination-down:
	soda migrate -c db_destination.yaml -p db_destination_migrations down

seed:
	go run main.go seed