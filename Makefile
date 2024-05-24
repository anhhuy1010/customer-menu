build-app:
	docker-compose build app
start:
	docker-compose up
restart:
	docker-compose restart
logs:
	docker logs -f customer-menu
ssh-app:
	docker exec -it customer-menu bash
swagger:
	swag init ./controllers/*
proto-user:
	protoc -I grpc/proto/user/ \
		-I /usr/include \
		--go_out=paths=source_relative,plugins=grpc:grpc/proto/user/ \
		grpc/proto/user/user.proto
