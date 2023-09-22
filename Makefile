INSTALL_DIR=${DESTDIR}/usr/bin
CONFIG_DIR=${DESTDIR}/etc/okki-status

build:
	go build -o okki-status -buildvcs=false

clean:
	rm -f okki-status

install: build
	mkdir -p ${INSTALL_DIR}
	mkdir -p ${CONFIG_DIR}
	cp okki-status ${INSTALL_DIR}
	cp etc/okki-refresh ${INSTALL_DIR}
	cp etc/config.yaml ${CONFIG_DIR}
	cp etc/example.yaml ${CONFIG_DIR}

uninstall:
	rm -f ${INSTALL_DIR}/okki-status
	rm -f ${INSTALL_DIR}/okki-refresh
	rm -rf ${CONFIG_DIR}
