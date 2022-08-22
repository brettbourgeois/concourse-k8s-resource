package main

import (
	"encoding/json"
	"github.com/brettbourgeois/concourse-k8s-resource/pkg/k8s"
	"github.com/brettbourgeois/concourse-k8s-resource/pkg/k8s/kubectl"
	"github.com/brettbourgeois/concourse-k8s-resource/pkg/models"
	"github.com/brettbourgeois/concourse-k8s-resource/pkg/utils"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"log"
	"os"
	"reflect"
	"time"
)

var streams = genericclioptions.IOStreams{
	In:     os.Stdin,
	Out:    os.Stderr, // concourse console
	ErrOut: os.Stderr,
}

func main() {

	var request models.OutRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatalln("Illegal input format", err)
	}

	utils.Debug(&request.Source, "request: ", request)
	utils.ChangeWorkingDir()

	clientset, clientConfig := k8s.NewClientSet(&request.Source)
	if request.Source.Namespace == "" {
		request.Source.Namespace = "default"
	}

	factory := kubectl.NewCommandFactory(&request.Params)
	commandConfig := &kubectl.CommandConfig{
		Clientset:    clientset,
		Discovery:    toDiscoveryInterface(clientset),
		ClientConfig: clientConfig,
		Streams:      streams,
		Namespace:    request.Source.Namespace,
		Resources:    request.Source.WatchResources,
		Params:       &request.Params,
	}
	if err := kubectl.RunCommand(factory, commandConfig); err != nil {
		log.Fatalln("cannot run kubectl command", err)
	}
	if requireStatusCheck(request.Params) {
		time.Sleep(5 * time.Second)
		log.Println("check status for", request.Source.WatchResources)
		if ok := k8s.CheckResourceStatus(clientset, request.Source.Namespace, request.Source.WatchResources, request.Params.StatusCheckTimeout); !ok {
			log.Fatalln("resource is not running...")
		}
	}

	response := createResponse(request, clientset)

	utils.Debug(&request.Source, "response: ", *response)
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		log.Fatalln("Output Failure", err)
	}
}

func requireStatusCheck(params models.OutParams) bool {
	return !params.Delete && !params.ServerDryRun && !params.Diff
}

func createResponse(request models.OutRequest, clientset kubernetes.Interface) *models.OutResponse {

	if request.Params.Delete {
		// resources is deleted, so just return empty response
		return emptyResponse()
	}

	// apply or undo
	version, err := k8s.GetCurrentVersion(&request.Source, clientset)
	if err != nil {
		return emptyResponse()
	}
	metadatas, err := k8s.GenerateMetadatas(&request.Source, clientset)
	if err != nil {
		return emptyResponse()
	}

	response := models.OutResponse{
		Version:  *version,
		Metadata: metadatas,
	}
	return &response
}

func emptyResponse() *models.OutResponse {
	return &models.OutResponse{
		Version:  models.Version{},
		Metadata: nil,
	}
}

func toDiscoveryInterface(obj interface{}) discovery.DiscoveryInterface {
	if discoveryIf, ok := obj.(discovery.DiscoveryInterface); ok {
		return discoveryIf
	}
	log.Fatalf("cannot cast to discovery interface from %s", reflect.TypeOf(obj))
	return nil
}
