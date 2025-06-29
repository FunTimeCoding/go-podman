package image

import "time"

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
