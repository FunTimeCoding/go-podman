package podman

const (
	LocatorEnvironment            = "PODMAN_LOCATOR"
	ContainerAutoStartEnvironment = "CONTAINER_AUTO_START"
	ImageAutoCleanEnvironment     = "IMAGE_AUTO_CLEAN"
	ImageAutoUpdateEnvironment    = "IMAGE_AUTO_UPDATE"

	Command = "podman"

	Image = "image"

	Execute = "exec"

	Machine = "machine"
	Inspect = "inspect"

	System     = "system"
	Connection = "connection"
	List       = "ls"
	Format     = "--format"
	Notation   = "json"

	FixtureIdentifier = "123456789012"
	FixtureName       = "test:latest"
)
