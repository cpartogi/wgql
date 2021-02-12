
local:
	go run github.com/99designs/gqlgen generate 
	go build main.go && ./main
