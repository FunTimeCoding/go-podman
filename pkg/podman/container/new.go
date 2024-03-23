package container

func New(
	name string,
	alias string,
	port int,
	environment map[string]string,
) *Container {
	return &Container{
		Name:        name,
		Alias:       alias,
		Port:        port,
		Environment: environment,
	}
}
