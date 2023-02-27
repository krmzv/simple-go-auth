build:
	go build -o ./bin/reversejs

run: build
	./bin/reversejs
	
test: 
	go test -verbose ./...