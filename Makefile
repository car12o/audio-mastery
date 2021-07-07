### SERVER ###

main:
	@docker-compose up

serve: install.devdeps
	@air -c .air.toml

serve.docs:
	@docker run --rm -v "${PWD}:/audio-mastery" -p 9000:9000 quay.io/goswagger/swagger serve \
		--no-open \
		--flavor swagger \
		--port 9000 \
		/audio-mastery/api/generated/swagger.yaml

install.devdeps:
	@GO111MODULE=off go get -v github.com/cosmtrek/air

### WEB ###

web.serve:
	@cd web && npm run start

web.install:
	@cd web && npm i

### GENERATE ###

gen: gen.clean gen.bundle gen.api gen.web
	@

gen.api:
	@docker run --rm -it -v "${PWD}:/audio-mastery" quay.io/goswagger/swagger generate server \
		--exclude-main \
  	-A audio-mastery \
		-P models.Principal \
  	-f /audio-mastery/api/generated/swagger.yaml \
  	-t /audio-mastery/api/generated && \
		sudo chown -R ${USER}:${$USER} ./api/generated

gen.web:
	@docker run --rm -v "${PWD}:/audio-mastery" swaggerapi/swagger-codegen-cli generate \
		-i /audio-mastery/api/generated/swagger.yaml \
		-l typescript-fetch \
		-o /audio-mastery/web/src/api/generated && \
		sudo chown -R ${USER}:${$USER} ./web/src/api/generated && \
		cd web/src/api/generated && ls -a -I '.' -I '..' | grep -v .ts | xargs rm -rf

gen.bundle:
	@docker run --rm -v "${PWD}:/audio-mastery" jeanberu/swagger-cli swagger-cli bundle \
		-r \
		-t yaml \
		-o /audio-mastery/api/generated/swagger.yaml \
		/audio-mastery/api/spec/swagger.yaml && \
		sudo chown -R ${USER}:${$USER} ./api/generated/swagger.yaml

gen.clean:
	@rm -rf ./api/generated/* && rm -rf ./web/src/api/generated/*
