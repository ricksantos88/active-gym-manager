## Rodar migration em ambiente local


`migrate create -ext sql -dir ./infrastructure/migrations/ -seq create_fixed_informations_tables`

`migrate -path ./infrastructure/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=active_gym_db" up`
