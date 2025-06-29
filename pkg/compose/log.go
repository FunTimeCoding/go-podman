package compose

import "fmt"

func (c *Compose) Log() {
	for _, o := range c.podman.Container(false) {
		fmt.Printf("Container: %s\n", o.Name)
		l := c.podman.Logs(o.RawList)

		if s := l.Output.Render(); s != "" {
			fmt.Println(s)
		}

		if s := l.Error.Render(); s != "" {
			fmt.Println(s)
		}
	}
}
