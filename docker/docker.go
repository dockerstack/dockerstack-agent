package main

import (
	_ "dockerstack/dockerstack-agent/config"
	_ "encoding/json"
	_ "errors"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

const (
	DEFAULT_ENDPOINT = "unix:///var/run/docker.sock"
)

type ListImages struct {
	Images []Images `json:"images"`
}

type Images struct {
	Created int64  `json:"created"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
}

func checkSockfile() (value bool) {

	if _, err := os.Stat("/var/run/docker.sock"); err == nil {

		value = true
		return value
	} else {
		value = false
		return value
	}
	return value
}

/*func RunningContainers() {

	if config.Docker("socket") == DEFAULT_ENDPOINT {
		client, _ := docker.NewClient(DEFAULT_ENDPOINT)

	} else {
		fmt.Println("Testing")
	}

}

*/
/*// APIImages represent an image returned in the ListImages call.
type APIImages struct {
	ID          string   `json:"Id" yaml:"Id"`
	RepoTags    []string `json:"RepoTags,omitempty" yaml:"RepoTags,omitempty"`
	Created     int64    `json:"Created,omitempty" yaml:"Created,omitempty"`
	Size        int64    `json:"Size,omitempty" yaml:"Size,omitempty"`
	VirtualSize int64    `json:"VirtualSize,omitempty" yaml:"VirtualSize,omitempty"`
	ParentID    string   `json:"ParentId,omitempty" yaml:"ParentId,omitempty"`
	RepoDigests []string `json:"RepoDigests,omitempty" yaml:"RepoDigests,omitempty"`
}*/
func ImagesList() (data *ListImages) {

	sockexist := checkSockfile()

	if sockexist == true {
		client, _ := docker.NewClient(DEFAULT_ENDPOINT)
		imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})

		for _, img := range imgs {

			data.Images[0].ID = img.ID
			data.Images[0].Name = img.RepoTags
			data.Images[0].Created = img.Created
			data.Images[0].Size = img.VirtualSize

			return data
		}

	} else {
		fmt.Println("Testing")
		return nil
	}

	return data

}

func main() {

	// RunningContainers()

	imgs := ImagesList()

	fmt.Println(imgs)

	fmt.Println(checkSockfile())

}
