PWD := $(shell pwd)

generate: clean
	# pull spec
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/master/spec.yaml > api-spec/spec.yaml
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/master/responses.yaml > api-spec/responses.yaml
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/master/parameters.yaml > api-spec/parameters.yaml
	curl -s https://raw.githubusercontent.com/giantswarm/api-spec/master/definitions.yaml > api-spec/definitions.yaml
	docker run --rm -it \
	  -e GOPATH=${HOME}/go:/go \
		-v ${HOME}:${HOME} \
		-w ${PWD}/api-spec \
		quay.io/goswagger/swagger:0.14.0 \
			generate client \
			--spec spec.yaml \
			--name gsclientgen \
			--default-scheme https \
			--target ..
	gofmt -s -l -w client
	gofmt -s -l -w models

# removal of all generated files
clean:
	rm -rf client
	rm -rf models
	rm -rf api-spec/*

# validate the spec
validate:
	docker run --rm -it \
	    -v ${PWD}/api-spec:/workdir \
	    boiyaa/yamllint:1.8.1 ./oai-spec.yaml
	docker run --rm -it \
		-v ${PWD}/api-spec:/swagger-api/yaml \
		jimschubert/swagger-codegen-cli:2.2.3 generate \
		--input-spec /swagger-api/yaml/oai-spec.yaml \
		--lang swagger \
		--output /tmp/

# validate the code through building
build:
	dep ensure
	go build github.com/giantswarm/gsclientgen/models
	go build github.com/giantswarm/gsclientgen/client
