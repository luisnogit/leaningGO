run: build
	./bin/main

build: cmd/main.go
	go build -o bin/main cmd/*.go

clean:
	rm bin/main