package compose

import "fmt"

func (t *Tester) Log() {
	for _, element := range t.podman.Container() {
		fmt.Printf("Container: %s\n", element.Names[0])
		l := t.podman.Logs(element)
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
