all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./fnormalizer.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.20.6 amd64
	sudo apk update

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/fnormalizer.exe $(DESTDIR)$(PREFIX)/bin/fnormalizer
	strip $(DESTDIR)$(PREFIX)/bin/fnormalizer
