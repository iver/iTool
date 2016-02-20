SHELL		:= bash
ACTUAL := $(shell pwd)
PKGNAME := itool.tar.gz
VERSION := app.version\=\($(shell git rev-parse --short HEAD)\)

export ACTUAL
export PKGNAME
export VERSION

install: base get

get:
	go get -v bitbucket.org/ivan-iver/config;
	go get -v gopkg.in/alecthomas/kingpin.v2;

build: install
	@go build -o bin/itool itool;
	@cp ${ACTUAL}/itool/app.conf bin/;
	@mkdir -p bin/log/ bin/vendor/;
	@cp itool/vendor/* bin/vendor/;

run: build
	@cd bin && ./itool;

base:
	@echo "---- Variables asignadas. ----";
	@echo "| GOPATH: ${GOPATH}";
	@echo "| ACTUAL: ${ACTUAL}";
	@echo "------------------------------";
	# Creating working directory.
	@if [ ! -d $$GOPATH/src/itool ]; then \
		mkdir -p $$GOPATH/src/itool; \
	fi;
	# Checking nubleer link
	@if [[ -L $$GOPATH/src/itool && -d $$GOPATH/src/itool ]]; then \
		echo "Skip Linked"; \
	else \
		ln -sf ${ACTUAL}/itool $$GOPATH/src/itool; \
		echo "Compiling ..."; \
	fi;

package: build
	@sed -i -e '2s/\(.*\)/${VERSION}/g' itool/app.conf
	@cp -r ${ACTUAL}/itool/app.conf bin/;
	@mkdir -p bin/log;
	@tar -zcf ${PKGNAME} bin;
	@rm -rf ${ACTUAL}/bin;
	@echo "Package done! ... you can run deploy.sh script from yout host machine.";

uninstall:
	@unlink ${GOPATH}/src/itool;
	@rm -f ${GOPATH}/bin/itool;

clean:
	@rm -rf ${ACTUAL}/${PKGNAME}
	@rm -rf ${ACTUAL}/bin;

