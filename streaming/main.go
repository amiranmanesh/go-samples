package main

import (
	"log"
	"net/http"
)

//ffmpeg -i [BachGavotteShort.mp3] -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts
//http: //localhost:8080/outputlist.m3u8
func main() {

	const file = "files"

	http.Handle("/", http.FileServer(http.Dir(file)))


	log.Fatal(http.ListenAndServe(":8080", nil))
}
