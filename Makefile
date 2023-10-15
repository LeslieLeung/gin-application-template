run:
	go run main.go serve -c config/example.config.yaml

dry-build:
	go build -o app main.go
	rm app