package common

import "fmt"

func Log(payload string) {
	fmt.Print(fmt.Sprintf("\r%s", payload))
}
