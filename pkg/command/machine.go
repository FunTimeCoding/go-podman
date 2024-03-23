package command

import "time"

type Machine struct {
	ConfigPath struct {
		Path string `json:"Path"`
	} `json:"ConfigPath"`
	ConnectionInfo struct {
		PodmanSocket any `json:"PodmanSocket"`
		PodmanPipe   struct {
			Path string `json:"Path"`
		} `json:"PodmanPipe"`
	} `json:"ConnectionInfo"`
	Created time.Time `json:"Created"`
	Image   struct {
		IgnitionFilePath struct {
			Path string `json:"Path"`
		} `json:"IgnitionFilePath"`
		ImageStream string `json:"ImageStream"`
		ImagePath   struct {
			Path string `json:"Path"`
		} `json:"ImagePath"`
	} `json:"Image"`
	LastUp    time.Time `json:"LastUp"`
	Name      string    `json:"Name"`
	Resources struct {
		CPUs     int   `json:"CPUs"`
		DiskSize int64 `json:"DiskSize"`
		Memory   int   `json:"Memory"`
	} `json:"Resources"`
	SSHConfig struct {
		IdentityPath   string `json:"IdentityPath"`
		Port           int    `json:"Port"`
		RemoteUsername string `json:"RemoteUsername"`
	} `json:"SSHConfig"`
	State              string `json:"State"`
	UserModeNetworking bool   `json:"UserModeNetworking"`
}
