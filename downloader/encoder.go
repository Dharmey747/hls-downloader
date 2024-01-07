package downloader

import (
	"HLS/common"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func encode(file_name string, format string) {
	files, _ := os.ReadDir("./.temp_decrypt")

	fileList := []int{}

	for _, file := range files {
		num, _ := strconv.Atoi(strings.Replace(strings.Replace(file.Name(), ".ts", "", 1), fmt.Sprintf("%s_", format), "", 1))

		fileList = append(fileList, num)
	}

	exec.Command("bash", "-c", `cat .temp_decrypt/`+format+`_{0..`+strconv.Itoa(slices.Max(fileList))+`}.ts > `+format+`_all.ts`).Run()
	exec.Command("bash", "-c", `ffmpeg -nostats -loglevel 0 -i `+format+`_all.ts -acodec copy -vcodec copy `+format+`_`+file_name+`.mp4`).Run()
	common.RemoveFile("./" + format + "_all.ts")
}

func render(queue_index string) {
	exec.Command("bash", "-c", `ffmpeg -i VIDEO_`+queue_index+`.mp4 -i AUDIO_`+queue_index+`.mp4 -shortest `+queue_index+`.mp4`).Run()
	common.RemoveFile("./VIDEO_" + queue_index + ".mp4")
	common.RemoveFile("./AUDIO_" + queue_index + ".mp4")
}
