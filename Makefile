.PHONY: run-virtual-environment
run-virtual-environment:
	cd deployments && docker-compose up -d

.PHONY: goose-up
goose-up:
	goose -dir migrations postgres "user=schedule_user dbname=schedule password=user sslmode=disable host=localhost port=5432" up

.PHONY: proto-generate
proto-generate:
	protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out . api/schedule.proto

.PHONY: goose-down
goose-down:
	goose -dir migrations postgres "user=schedule_user dbname=schedule password=user sslmode=disable host=localhost port=5432" down

.PHONY: stop-virtual-environment
stop-virtual-environment:
	cd deployments && docker-compose down
