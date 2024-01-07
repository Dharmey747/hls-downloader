package main

import (
	"HLS/common"
	"HLS/downloader"
	"fmt"
	"strconv"
)

func main() {
	common.Clean(true)

	queue := []map[string]string{
		{
			"VIDEO": "",
			"AUDIO": "",
		},
	}

	fmt.Println("Preparing to download " + strconv.Itoa(len(queue)) + " elements.")

	for index, download := range queue {
		common.Clean(false)

		downloader.DownloadFile(download, index+1)
		common.Clean(false)
	}

	common.Clean(true)
	fmt.Println("All downloads have been completed. Enjoy!")
}
