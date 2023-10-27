package common

import (
	"os"
)

func Clean(force_delete bool) {
	if force_delete {
		os.RemoveAll("./.temp_decrypt")
		os.RemoveAll("./.temp_downloads")
		return
	}

	_, err := os.Stat("./.temp_decrypt")
	if err != nil {
		os.MkdirAll("./.temp_decrypt", os.ModePerm)
	} else {
		os.RemoveAll("./.temp_decrypt")
	}

	_, err = os.Stat("./.temp_downloads")
	if err != nil {
		os.MkdirAll("./.temp_downloads", os.ModePerm)
	} else {
		os.RemoveAll("./.temp_downloads")
	}
}
