package compose

func (t *Tester) Create() {
	for _, element := range t.containers {
		if !t.podman.ImageExists(element) {
			t.podman.Pull(element)
		}

		if !t.podman.Exists(element) {
			t.podman.Create(element)
		}

		if !t.podman.Running(element) {
			t.podman.Start(element)
		}
	}
}
