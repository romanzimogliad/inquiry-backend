migrate -path ./db/migrations -database "postgres://user:password@localhost:5437/inquiry?sslmode=disable" down

migrate -path ./db/migrations -database "postgres://user:password@localhost:5437/inquiry?sslmode=disable" up
