build:
	cd okki-refresh; \
		go build -o okki-refresh
	cd okki-status; \
		go build -o okki-status

clean:
	rm -f okki-refresh/okki-refresh
	rm -f okki-status/okki-status

install: build
	mv okki-refresh/okki-refresh /usr/local/bin
	mv okki-status/okki-status /usr/local/bin

