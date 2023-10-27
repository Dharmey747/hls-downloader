package downloader

import (
	"HLS/client"
	"HLS/common"
	"strconv"
)

func download(url string, index int, enc map[string]any) {
	var path string
	if enc["required"].(bool) {
		path = "./.temp_downloads/" + strconv.Itoa(index) + ".ts"
	} else {
		path = "./.temp_decrypt/" + strconv.Itoa(index) + ".ts"
	}

	for {
		request := client.Request(url, "", map[string]string{})

		if request != nil {
			common.CreateFile(path, request)

			if enc["required"].(bool) {
				if decrypt(path, enc["key"].([]byte), enc["iv"].(string)) {
					break
				}
			}
		}
	}
}
