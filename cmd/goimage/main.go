package main

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/client/image"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func main() {
	// TODO: If older version of same image exists, remove it
	// TODO: Checks all images for updates
	c := client.New()
	f := constant.Format
	byName := make(map[string][]*image.Image)

	for _, i := range c.Image2() {
		byName[i.Name] = append(byName[i.Name], i)
	}

	for k, v := range byName {
		if len(v) > 1 {
			fmt.Printf("Multiple images with name %s:\n", k)

			for _, i := range v {
				fmt.Printf("  - %s\n", i.Format(f))
			}
		}
	}
}
