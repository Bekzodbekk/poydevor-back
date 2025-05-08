DB_URL = postgres://postgres:1@localhost:5432/poydevor_usersdb?sslmode=disable

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/user.proto
	

migrate-up:
	migrate -path ./db/migration -database ${DB_URL} up

migrate-down:
	migrate -path ./db/migration -database ${DB_URL} down

migrate-force:
	migrate -path=db/migration -database ${DB_URL} verbose -force 1