package mime

mimeTypes := map[string]string {
	"mp3":  "audio/mpeg",
	"flac": "audio/x-flac",
	"aac":  "audio/x-aac",
	"m4a":  "audio/m4a",
	"m4b":  "audio/m4b",
	"ogg":  "audio/ogg",
	"opus": "audio/ogg",
	"wma":  "audio/x-ms-wma",
}

func FromExtension(ext string) (string, bool) {
	v, ok := mimeTypes[ext]
	return v, ok
}