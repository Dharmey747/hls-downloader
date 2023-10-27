package downloader

import (
	"HLS/common"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func encode(file_name string) {
	files, _ := os.ReadDir("./.temp_decrypt")

	fileList := []int{}

	for _, file := range files {
		num, _ := strconv.Atoi(strings.Replace(file.Name(), ".ts", "", 1))

		fileList = append(fileList, num)
	}

	exec.Command("bash", "-c", `cat .temp_decrypt/{0..`+strconv.Itoa(slices.Max(fileList))+`}.ts > all.ts`).Run()
	exec.Command("bash", "-c", `ffmpeg -nostats -loglevel 0 -i all.ts -acodec copy -vcodec copy `+file_name+`.mp4`).Run()
	common.RemoveFile("./all.ts")
}
