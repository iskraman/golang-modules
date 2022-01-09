all: static

dynamic:
	gcc -fPIC -c a.c
	gcc -shared -o libmylib.so a.o
	rm -f a.o

static:
	gcc -c a.c
	ar rc libmylib.a a.o
	ranlib libmylib.a
	rm -f a.o
