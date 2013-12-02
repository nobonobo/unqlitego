# for GNU make
.PHONY: all build install test clean

all: build	

build: src/unqlite.c
	go build

install: src/unqlite.c
	go install

clean:
	rm -rf src/*

test:
	go test ./

src/unqlite.c:
	git submodule init
	git submodule update
