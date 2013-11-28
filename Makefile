# for GNU make

all: libunqlite.a
	go install

src/unqlite.c:
	git submodule init
	git submodule update

libunqlite.a: src/unqlite.c src/unqlite.h
	gcc -c src/unqlite.c -I./src -DUNQLITE_ENABLE_THREADS=1
	ar rv libunqlite.a unqlite.o
