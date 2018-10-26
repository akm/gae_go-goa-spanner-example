##### Convenient command ######

REPO:=github.com/akm/gae_go-goa-spanner-example
GAE_PROJECT ?= dummy-project
BASEDIR = $(CURDIR)
VERSION ?= $(shell cat ./VERSION)

SERVER_BASE=$(REPO)

LOCAL_PORT ?= 8080

dep_init:
	@dep init

dep_ensure:
	@dep ensure

goa_bootstrap:
	@mv vendor vendor.bak
	@goagen bootstrap -d $(SERVER_BASE)/design ; mv vendor.bak vendor

rm_golang_files:
	@rm -f *.go

goa_setup: goa_bootstrap appengine/main.go rm_golang_files goa_controller

init: install bootstrap import
gen: clean generate import

# goagen douments
# https://goa.design/implement/goagen/
# https://goa.design/ja/implement/goagen/

# Rename vendor during executing goagen
#	https://github.com/goadesign/goa/issues/923#issuecomment-290424097
bootstrap: generate main

main: controller appengine/main.go

appengine/main.go:
	@mv vendor vendor.bak
	@goagen main -d $(SERVER_BASE)/design >/dev/null ; mv vendor.bak vendor
	@mkdir -p appengine
	@mv main.go appengine
	@rm *.go
	@echo 'appengine/main.go'
	@echo '1. Change package from "main" to "appengine"'
	@echo '2. Add "net/http" to import section'
	@echo '3. Add "$(SERVER_BASE)/controller" to import section'
	@echo '4. Change "func main()" to "func init()"'
	@echo '5. Add "controller." before each "NewXxxxController"'
	@echo '6. Comment out the lines below the comment "Start service"'
	@echo '7. Add http.HandleFunc("/", service.Mux.ServeHTTP) at the end of init func'

appengine/swagger:
	cd appengine && rm -f swagger && ln -s ../swagger; cd -
appengine/swaggerui:
	cd appengine && rm -f swaggerui && ln -s ../swaggerui; cd -

appengine/app.yaml:
	@echo 'You have to setup appengine/app.yaml to build'

appengine: appengine/main.go appengine/swagger appengine/swaggerui appengine/app.yaml


app:
	@mv vendor vendor.bak
	@goagen app -d $(SERVER_BASE)/design ; mv vendor.bak vendor

controller: goa_controller

goa_controller:
	@mv vendor vendor.bak
	@mkdir -p controller
	@goagen controller  -d $(SERVER_BASE)/design --pkg controller --out controller --app-pkg ../app ; mv vendor.bak vendor

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

goa_gen: gen_app gen_client gen_swagger

gen_app:
	@mv vendor vendor.bak
	@goagen app            -d $(SERVER_BASE)/design; mv vendor.bak vendor

gen_client:
	@mv vendor vendor.bak
	@goagen client         -d $(SERVER_BASE)/design; mv vendor.bak vendor

gen_swagger:
	@mv vendor vendor.bak
	@goagen swagger        -d $(SERVER_BASE)/design; mv vendor.bak vendor


generate: goa_gen

swaggerui:
	rm -rf ./tmp/swaggerui && \
	mkdir -p ./tmp/swaggerui && \
	cd ./tmp/swaggerui && \
		git init && \
		git config core.sparsecheckout true && \
		git remote add origin git@github.com:swagger-api/swagger-ui.git && \
		echo "dist" > .git/info/sparse-checkout && \
		git pull origin master && \
		cd - && \
	mkdir -p swaggerui && \
	cp -R tmp/swaggerui/dist swaggerui/

install: install_dep vendor/github.com/goadesign/goa
install_dep:
	@which dep || go get -u github.com/golang/dep/cmd/dep
	@dep ensure

vendor/github.com/goadesign/goa:
	@go get -u github.com/goadesign/goa/...

setup: install

import:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"

fmt:
	@gofmt -w controller/*.go

test-controller:
	goapp test $(SERVER_BASE)/controller

test-server: build-server test-controller

test: build test-server


build:
	goapp build $(SERVER_BASE)/appengine


.PHONY: version
version:
	@echo $(VERSION)

DEPLOY_PATH=./appengine
deploy:
	appcfg.py -A $(GAE_PROJECT) -V ${VERSION} update $(DEPLOY_PATH)
rollback:
	appcfg.py -A $(GAE_PROJECT) -V ${VERSION} rollback $(DEPLOY_PATH)

local_http_server:
	dev_appserver.py --port=$(LOCAL_PORT) --enable_console --skip_sdk_update_check=yes $(BASEDIR)/appengine/app.yaml

local_https_server:
	dev_appserver.py --port=$(LOCAL_PORT) --enable_console --skip_sdk_update_check=yes $(BASEDIR)/appengine/app.yaml \
		--ssl_certificate_path $(SSL_CERTIFICATE_PATH) \
		--ssl_certificate_key_path $(SSL_CERTIFICATE_KEY_PATH)

local: swaggerui local_http_server
local_https: swaggerui local_https_server

dev: build local
