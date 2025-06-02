package compose

import "fmt"

func (c *Compose) Log() {
	for _, o := range c.podman.Container() {
		fmt.Printf("Container: %s\n", o.Names[0])
		l := c.podman.Logs(o)
		output := l.Output.Render()

		if output != "" {
			fmt.Println(output)
		}

		errorOutput := l.Error.Render()

		if errorOutput != "" {
			fmt.Println(errorOutput)
		}
	}
}
