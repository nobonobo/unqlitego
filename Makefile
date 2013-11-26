# for GNU make

libunqlite.a: src/unqlite.c src/unqlite.h
	gcc -c src/unqlite.c -I./src -DUNQLITE_ENABLE_THREADS=1
	ar rv libunqlite.a unqlite.o
