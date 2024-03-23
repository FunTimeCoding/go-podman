package compose

func (t *Tester) Destroy() {
	for _, element := range t.containers {
		if !t.podman.Exists(element) {
			return
		}

		if t.podman.Running(element) {
			t.podman.Stop(element)
		}

		t.podman.Remove(element)
	}
}
