PWD := $(shell pwd)

generate: clean
	docker run --rm -it \
	  -e GOPATH=${HOME}/go:/go \
		-v ${HOME}:${HOME} \
		-w ${PWD} \
		quay.io/goswagger/swagger:0.14.0 \
			generate client \
			--spec api-spec/oai-spec.yaml \
			--name gsclientgen \
			--default-scheme https
	gofmt -s -l -w client
	gofmt -s -l -w models

# removal of all generated files
clean:
	rm -rf client
	rm -rf models

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
