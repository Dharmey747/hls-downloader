# HLS Downloader

A simple and easy to use golang program to download videos from the internet that are served using the **HTTP Live Streaming** protocol.

## How it works

The HLS protocol sends video to the browser in small bits (**.ts** files) that can range from 2-15 seconds, so the user won't have to load the entire video file before watching it. It gets loaded as you watch.

When the video starts, the browser receives a playlist file that contains all the URLs for those small video bits and some other info like protocol version, encryption data (if any) and so on.

This program takes the playlist URL and first checks if decryption of the files is required, and then starts downloading all of them using multithreading. This takes a lot of bandwidth but it speeds up the download process a lot.

* *Only AES-128-CBC Decryption is supported.*

After the download has completed, the program *merges* all the **.ts** files, and then uses **FFMPEG** to convert it to a compressed **.mp4** file. You can then use this file to watch what you've downloaded on any device without needing an internet connection.

The program also supports a queue system if you'd wish to download multiple videos in one go, you'll simply have to put all the playlist URLs in the `queue` map inside `main.go`.

## Notes

I've coded this for my personal use, that's why the features might seem kinda weird and out of place. Feel free to adapt it and add stuff that fits you best.

Some bugs might of course be present, if you want you can report them by making a pull request and I'll look at it.

This has also only been tested on `macOS`, so I'm not 100% sure if it will work fine also on Windows and Linux.