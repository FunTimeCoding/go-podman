package compose

import "fmt"

func (c *Compose) Status() {
	for _, element := range c.podman.Pod() {
		fmt.Printf("Pod: %+v\n", element)
	}

	for _, element := range c.podman.Container() {
		fmt.Printf("Container %s: %+v\n", element.Names[0], element)
		l := c.podman.Execute(element, "ls")
		output := l.Output.Format()

		if output != "" {
			fmt.Println(output)
		}

		errorOutput := l.Error.Format()

		if errorOutput != "" {
			fmt.Println(errorOutput)
		}
	}
}
