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

func DownloadFile(url string, queue_index int) {
	EncryptionData := map[string]any{
		"required": false,
	}

	common.Log("Getting Encryption Data...")

	playlist := string(client.Request(url, "", map[string]string{}))
	if strings.Contains(playlist, "EXT-X-KEY:METHOD=AES-128") {
		common.Log("[AES] Decryption Required. Getting Keys...")

		EncryptionData["required"] = true

		EncryptionData["iv"] = strings.Split(strings.Split(playlist, ",IV=0x")[1], "\n")[0]

		keyURL := "https://" + strings.Split(strings.Split(url, "https://")[1], "/")[0] + strings.Split(strings.Split(playlist, `URI="`)[1], `"`)[0]
		key := client.Request(keyURL, "", map[string]string{})

		EncryptionData["key"] = key

		common.Log("[AES] Successfully Retrieved Decryption Keys.")
	} else {
		common.Log("[AES] No Encryption Detected.")
	}

	fmt.Println("")

	tsPlaylist := []string{}

	for _, line := range strings.Split(playlist, "\n") {
		if !strings.Contains(line, "#") {
			tsPlaylist = append(tsPlaylist, line)
		}
	}

	for index, bit := range tsPlaylist {
		go download(bit, index, EncryptionData)
	}

	for {
		count, _ := os.ReadDir("./.temp_decrypt")

		if len(count) == len(tsPlaylist) {
			time.Sleep(5 * time.Second)
			fmt.Println("\nDownload Complete.")

			fmt.Println("Encoding mp4 file...")
			encode(strconv.Itoa(queue_index))
			fmt.Println("Encoding Completed.")
			return
		} else {
			currentAmount, _ := os.ReadDir("./.temp_decrypt")
			progress := strconv.Itoa(int(len(currentAmount)*100/len(tsPlaylist))) + " %"
			amount := " [" + strconv.Itoa(len(currentAmount)) + " / " + strconv.Itoa(len(tsPlaylist)) + "] "

			common.Log("(" + progress + ") " + amount + "Downloading...")
		}
	}
}
