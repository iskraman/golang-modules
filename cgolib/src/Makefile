all: static

dynamic:
	gcc -fPIC -c a.c
	gcc -shared -o libmylib.so a.o
	rm -f a.o

static:
	arch -x86_64 gcc -c a.c
	ar rcs libmylib.a a.o
	ranlib libmylib.a
	rm -f a.o
