API_VERSION ?= 1
HANDLER ?= handler
PACKAGE ?= package

ifdef DOTENV
	DOTENV_TARGET = dotenv
else
	DOTENV_TARGET = .env
endif

.env:
	@echo "Create .env with .env.template"
	cp .env.template .env

dotenv:
	@echo "Overwrite .env with $(DOTENV)"
	cp $(DOTENV) .env

clean:
	rm -fr Gopkg.lock $(HANDLER).so $(PACKAGE).zip .serverless

deps:
	docker-compose down --remove-orphans
	docker-compose run --rm google-go-shim make _depsGo

test: $(DOTENV_TARGET)
	docker-compose run --rm google-go-shim make _test

buildLambda: $(DOTENV_TARGET) deps
	docker-compose run --rm aws-go-shim make _buildLambda

buildCloudFunction: $(DOTENV_TARGET) deps
	docker-compose run --rm google-go-shim make _buildCloudFunction

deploy: $(DOTENV_TARGET)
	docker-compose run --rm serverless make _deploy

_depsGo:
	dep ensure
	dep ensure -update

_test:
	cat test.json | cloud-functions-go-shim -entry-point $(ENTRY_POINT) -event-type http -plugin-path $(HANDLER).so

_build:
	build $(HANDLER)

_buildLambda: _build
	pkg $(HANDLER) $(PACKAGE)

_buildCloudFunction: _build
	cloud-functions-go -entry-point $(ENTRY_POINT) -event-type http -plugin-path $(HANDLER).so -o $(PACKAGE).zip

node_modules:
	npm install --save serverless-google-cloudfunctions

_deploy: node_modules
	rm -fr .serverless
	sls deploy -v -s $(ENV)

shellGoAWS: $(DOTENV_TARGET)
	docker-compose run --rm aws-go-shim bash

shellGoGoogle: $(DOTENV_TARGET)
	docker-compose run --rm google-go-shim bash

shellServerless: $(DOTENV_TARGET)
	docker-compose run --rm serverless bash
