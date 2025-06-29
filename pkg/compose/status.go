package compose

import "fmt"

func (c *Compose) Status() {
	for _, o := range c.podman.Pod() {
		fmt.Printf("Pod: %+v\n", o)
	}

	for _, o := range c.podman.Container(false) {
		fmt.Printf("Container %s: %+v\n", o.Name, o)
		l := c.podman.Execute(o.RawList, "ls")

		if s := l.Output.Render(); s != "" {
			fmt.Println(s)
		}

		if s := l.Error.Render(); s != "" {
			fmt.Println(s)
		}
	}
}
