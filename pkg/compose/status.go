package compose

import "fmt"

func (c *Compose) Status() {
	for _, o := range c.podman.Pod() {
		fmt.Printf("Pod: %+v\n", o)
	}

	for _, o := range c.podman.Container() {
		fmt.Printf("Container %s: %+v\n", o.Names[0], o)
		l := c.podman.Execute(o, "ls")
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
