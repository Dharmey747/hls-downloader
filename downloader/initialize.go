package downloader

import (
	"HLS/client"
	"HLS/common"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DownloadFile(downloadData map[string]string, queue_index int) {
	count := 0
	common.Log("Getting Encryption Data...")

	queue_index_string := strconv.Itoa(queue_index)

	for download_type, download_link := range downloadData {
		if download_link == "" {
			fmt.Println(fmt.Sprintf("\nSkipping %s download because the URL is empty...\n", download_type))
			continue
		}

		count++

		PlaylistData := map[string]any{
			"required": false,
		}

		playlist := string(client.Request(download_link, "", map[string]string{}))
		if strings.Contains(playlist, "EXT-X-KEY:METHOD=AES-128") {
			common.Log("[AES] Decryption Required. Getting Keys...")

			PlaylistData["required"] = true

			PlaylistData["iv"] = strings.Split(strings.Split(playlist, ",IV=0x")[1], "\n")[0]

			keyURL := "https://" + strings.Split(strings.Split(download_link, "https://")[1], "/")[0] + strings.Split(strings.Split(playlist, `URI="`)[1], `"`)[0]
			key := client.Request(keyURL, "", map[string]string{})

			PlaylistData["key"] = key

			common.Log("[AES] Successfully Retrieved Decryption Keys.")
		} else {
			common.Log("[AES] No Encryption Detected.")
		}

		fmt.Println("")

		tsPlaylist := []string{}

		for _, line := range strings.Split(playlist, "\n") {
			if !strings.Contains(line, "#") && line != "" {
				tsPlaylist = append(tsPlaylist, line)
			}
		}

		for index, bit := range tsPlaylist {
			go download(bit, download_type, index, PlaylistData)
		}

		for {
			count, _ := os.ReadDir("./.temp_decrypt")

			if len(count) == len(tsPlaylist) {
				time.Sleep(5 * time.Second)
				fmt.Println("\nDownload Complete.")

				fmt.Println("Encoding mp4 file...")
				encode(queue_index_string, download_type)
				fmt.Println("Encoding Completed.")

				common.Clean(false)
				common.Clean(false)
				break
			} else {
				currentAmount, _ := os.ReadDir("./.temp_decrypt")
				progress := strconv.Itoa(int(len(currentAmount)*100/len(tsPlaylist))) + " %"
				amount := " [" + strconv.Itoa(len(currentAmount)) + " / " + strconv.Itoa(len(tsPlaylist)) + "] "

				common.Log("(" + progress + ") " + amount + "Downloading...")
			}
		}
	}

	if count == 2 { // rendering is needed to merge video and audio into a single .mp4 file
		fmt.Println("\nRendering files...")
		render(queue_index_string)
	} else {
		os.Rename(fmt.Sprintf("VIDEO_%s.mp4", queue_index_string), fmt.Sprintf("%s.mp4", queue_index_string)) // rename the file from VIDEO_X.mp4 to X.mp4, beause rendering isn't needed.
	}
}
