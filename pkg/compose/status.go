package compose

import "fmt"

func (t *Tester) Status() {
	for _, element := range t.podman.Pod() {
		fmt.Printf("Pod: %+v\n", element)
	}

	for _, element := range t.podman.Container() {
		fmt.Printf("Container %s: %+v\n", element.Names[0], element)
		l := t.podman.Execute(element, "ls")
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
