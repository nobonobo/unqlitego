# for GNU make
.PHONY: all build test clean

all: build	

build: libunqlite.a
	go build

install:
	go install

clean:
	rm -f libunqlite.a
	rm -rf src/*

test:
	go test ./

src/unqlite.c:
	git submodule init
	git submodule update

libunqlite.a: src/unqlite.c src/unqlite.h
	gcc -c src/unqlite.c -I./src -DUNQLITE_ENABLE_THREADS=1
	ar rv libunqlite.a unqlite.o
