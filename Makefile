build:
	cd okki-status; \
		go build -o okki-status

clean:
	rm -f okki-status/okki-status

install: build
	mv okki-status/okki-status /usr/local/bin

