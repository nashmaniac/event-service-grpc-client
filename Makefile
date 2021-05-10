run:
	go run main.go
gen:
	protoc -I=. proto/*.proto --go_out=plugins=grpc:.
clean:
	rm pb/*.go