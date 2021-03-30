build:
	go build -o okki-status

clean:
	rm -f okki-status

install: build
	mv okki-status /usr/local/bin

