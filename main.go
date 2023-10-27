package main

import (
	"HLS/common"
	"HLS/downloader"
	"fmt"
	"strconv"
)

func main() {
	common.Clean(true)

	queue := []string{}

	fmt.Println("Preparing to download " + strconv.Itoa(len(queue)) + " elements.")

	for index, url := range queue {
		common.Clean(false)

		downloader.DownloadFile(url, index+1)
		common.Clean(false)
	}

	common.Clean(true)
	fmt.Println("All downloads have been completed. Enjoy!")
}
