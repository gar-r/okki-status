build:
	go build -o okki-status -buildvcs=false

clean:
	rm -f okki-status

install: build
	mv okki-status /usr/local/bin

