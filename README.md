[![GoDoc](https://godoc.org/github.com/giantswarm/gsclientgen?status.svg)](https://godoc.org/github.com/giantswarm/gsclientgen)
[![Travis-CI Build Badge](https://api.travis-ci.org/giantswarm/gsclientgen.svg?branch=master)](https://travis-ci.org/giantswarm/gsclientgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/gsclientgen)](https://goreportcard.com/report/github.com/giantswarm/gsclientgen)
[![IRC Channel](https://img.shields.io/badge/irc-%23giantswarm-blue.svg)](https://kiwiirc.com/client/irc.freenode.net/#giantswarm)

# Giant Swarm Golang Client (generated)

Experimental Go client for the Giant Swarm API, auto-generated based on an OAI/Swagger specification using Swagger Codegen.

Note: This client currently covers only a part of the API. Expect lots of breaking changes within the code. Use at your own risk.

Documentation can be found in the sub folder `docs`.

## Usage

In your Go package, import like this:

```go
import (
  "github.com/giantswarm/gsclientgen/client"
  "github.com/giantswarm/gsclientgen/client/operations"

  "github.com/go-openapi/strfmt"
)
```

Here is a simplistic way to use the client

```go
// transport is a http.Transport set up with the right endpoint, scheme,
// TLS config, proxy settings, timeouts etc.

c, err := client.New(transport, strfmt.Default)
if err != nil {
  fmt.Println(err)
}

params := operations.NewGetClustersParams()
params.Authorization = "giantswarm " + token

clustersResponse, err := apiClient.Operations.GetClusters(params)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("First cluster ID is '%s'\n", clustersResponse.Payload[0].ID)
```

## Development

The source API specification can be found in `api-spec/oai-spec.yaml`. Changes here will recflect in changes of the generated code.

To generate client library code after changes in above file, run:

```nohighlight
make validate
make generate
```
