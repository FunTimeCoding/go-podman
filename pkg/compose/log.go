package compose

import "fmt"

func (c *Compose) Log() {
	for _, element := range c.podman.Container() {
		fmt.Printf("Container: %s\n", element.Names[0])
		l := c.podman.Logs(element)
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
