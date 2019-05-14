package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/containerd/containerd/remotes/docker"
	"github.com/deislabs/oras/pkg/content"
	"github.com/deislabs/oras/pkg/oras"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	// Remote reference
	// Command used to run a local test registry on port 5000:
	//   docker run --rm -it -p 5000:5000 registry
	RemoteRef = "localhost:5000/opa/mybundle:1.0.0"

	// OPA-specific file names
	// Command used to create dummy files created for test purposes:
	//   for f in ".manifest" "helpers.rego" "deny.rego" "samples.json"; do echo "hello world" > $f; done
	OpenPolicyAgentManifestFileName = ".manifest"
	OpenPolicyAgentRegoFileName     = "helpers.rego"
	OpenPolicyAgentPolicyFileName   = "deny.rego"
	OpenPolicyAgentDataFileName     = "samples.json"

	// OPA-specific media types
	OpenPolicyAgentConfigMediaType        = "application/vnd.cncf.openpolicyagent.config.v1+json"
	OpenPolicyAgentManifestLayerMediaType = "application/vnd.cncf.openpolicyagent.manifest.layer.v1+json"
	OpenPolicyAgentRegoLayerMediaType     = "application/vnd.cncf.openpolicyagent.rego.layer.v1+rego"
	OpenPolicyAgentPolicyLayerMediaType   = "application/vnd.cncf.openpolicyagent.policy.layer.v1+rego"
	OpenPolicyAgentDataLayerMediaType     = "application/vnd.cncf.openpolicyagent.data.layer.v1+json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileContents(fileName string) []byte {
	b, err := ioutil.ReadFile(fileName)
	check(err)
	return b
}

func main2() {
	// Setup
	var contents []byte
	var layer ocispec.Descriptor
	var layers []ocispec.Descriptor
	resolver := docker.NewResolver(docker.ResolverOptions{})
	memoryStore := content.NewMemoryStore()

	// Add layers
	contents = getFileContents(OpenPolicyAgentManifestFileName)
	layer = memoryStore.Add(OpenPolicyAgentManifestFileName, OpenPolicyAgentManifestLayerMediaType, contents)
	layers = append(layers, layer)

	contents = getFileContents(OpenPolicyAgentRegoFileName)
	layer = memoryStore.Add(OpenPolicyAgentRegoFileName, OpenPolicyAgentRegoLayerMediaType, contents)
	layers = append(layers, layer)

	contents = getFileContents(OpenPolicyAgentPolicyFileName)
	layer = memoryStore.Add(OpenPolicyAgentPolicyFileName, OpenPolicyAgentPolicyLayerMediaType, contents)
	layers = append(layers, layer)

	contents = getFileContents(OpenPolicyAgentDataFileName)
	layer = memoryStore.Add(OpenPolicyAgentDataFileName, OpenPolicyAgentDataLayerMediaType, contents)
	layers = append(layers, layer)

	// Push
	fmt.Printf("Pushing OPA bundle to %s...\n", RemoteRef)
	extraOpts := []oras.PushOpt{oras.WithConfigMediaType(OpenPolicyAgentConfigMediaType)}
	manifest, err := oras.Push(context.Background(), resolver, RemoteRef, memoryStore, layers, extraOpts...)
	check(err)
	fmt.Printf("Pushed OPA bundle to %s with digest %s\n", RemoteRef, manifest.Digest)
}
