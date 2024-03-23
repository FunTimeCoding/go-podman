package container

type Container struct {
	Name        string
	Alias       string
	Port        int
	Environment map[string]string
}
