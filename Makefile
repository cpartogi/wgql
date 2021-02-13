
PROJECT_NAME=wgql

install:
	cd .. && go get -u github.com/cosmtrek/air && cd ${PROJECT_NAME} 

local:
	air -c config/.air.toml

generate:
	go run github.com/99designs/gqlgen generate 


compose-up:
	docker-compose up -d --build

compose-down:
	docker-compose down