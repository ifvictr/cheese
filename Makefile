MAIN=./cmd/cheese.go

all: run

build:
	go build -o ./bin/cheese ${MAIN}

clean:
	rm -rfv ./bin

run:
	go run -race ${MAIN}
