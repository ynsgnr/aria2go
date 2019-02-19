LDFLAGS=-L.-l aria2 -l aria2go
CXX=g++
CC=gcc
CFLAGS=-I.
CXXFLAGS=$(CFLAGS)

ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

libaria2go.so: aria2go.o aria2Interface.o
	$(CXX) -shared $^ -o $@

aria2go.o aria2Interface.o : CXXFLAGS+=-fPIC

install: libaria2go.so
	install -d $(DESTDIR)$(PREFIX)/lib/
	install -m 644 libaria2go.so $(DESTDIR)$(PREFIX)/lib/