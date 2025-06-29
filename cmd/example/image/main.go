package main

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/system/run"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

type Image struct {
	Id             string            `json:"Id"`
	ParentId       string            `json:"ParentId"`
	RepoTags       []string          `json:"RepoTags"`
	RepoDigests    []string          `json:"RepoDigests"`
	Size           int64             `json:"Size"`
	SharedSize     int               `json:"SharedSize"`
	VirtualSize    int64             `json:"VirtualSize"`
	Labels         map[string]string `json:"Labels"`
	Containers     int               `json:"Containers"`
	Arch           string            `json:"Arch"`
	Digest         string            `json:"Digest"`
	History        []string          `json:"History"`
	IsManifestList bool              `json:"IsManifestList"`
	Names          []string          `json:"Names"`
	Os             string            `json:"Os"`
	Created        int               `json:"Created"`
	CreatedAt      time.Time         `json:"CreatedAt"`
}

func (i *Image) Name() string {
	return i.Names[0]
}

func (i *Image) NameVersion() (string, string) {
	return key_value.Colon(i.Name())
}

func main() {
	// TODO: If older version of same image exists, remove it

	// TODO: Tool chat checks all images for updates

	for _, image := range Images() {
		name, version := image.NameVersion()
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Version: %s\n", version)

		if false {
			fmt.Printf(
				"Create: %s\n",
				image.CreatedAt.Format(timeLibrary.DateMinute),
			)
		}

		if false {
			fmt.Printf("Size: %s\n", units.HumanSize(float64(image.Size)))
		}
	}
}

func Images() []*Image {
	r := run.New()
	r.Start("podman", "image", "ls", "--format", "json")
	var images []*Image
	notation.DecodeStrict(r.OutputString, &images, true)

	return images
}
