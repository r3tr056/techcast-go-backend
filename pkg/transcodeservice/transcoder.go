package transcodeservice

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/google/shlex"
)

type Transcoder interface {
	Transcode(ctx context.Context, profile AudioProfile, in string, out io.Writer) error
}

var AudioProfiles = map[string]AudioProfile {
	"mp3": MP3Profile,
	"mp3_rg": MP3RGProfile,
	"opus_car": OpusCarProfile,
	"opus_rg": OpusRGProfile,
	"opus": OpusProfile,
	"opus_128_car": Opus128CarProfile,
	"opus_128": Opus128Profile,
	"opus_128_rg": Opus128RGProfile,
}

var (
	ErrNoProfileParts = fmt.Errorf("not enough details to construct profile")
)

var (
	MP3Profile = NewProfile("audio/mpeg", 128)
	MP3RGProfile = NewProfile("audio/mpeg", 128)
	OpusCarProfile = NewProfile("audio/ogg", 96)
	OpusProfile = NewProfile("audio/ogg", 96)
	OpusRGProfile = NewProfile("audio/ogg", 96)
	Opus128CarProfile = NewProfile("audio/ogg", 128)
	Opus128Profile = NewProfile("audio/ogg", 128)
	Opus128RGProfile = NewProfile("audio/ogg", 128)

	PCM12leProfile = NewProfile("audio/wav", 0)
)

type Bitrate int

type AudioProfile struct {
	bitrate Bitrate
	seek time.Duration
	mime string
	exec string
}

func (ap *AudioProfile) Bitrate() Bitrate { return p.bitrate }
func (ap *AudioProfile) Seek() time.Duration { return p.seek }
func (ap *AudioProfile) MIME() string { return p.mime }

func NewProfile(mime string, bitrate Bitrate, exec string) AudioProfile {
	return AudioProfile{mime: mime, bitrate: bitrate, exec: exec}
}

func WithBitrate(ap AudioProfile, bitrate Bitrate) AudioProfile {
	p.bitrate = bitrate
	return p
}

func WithSeek(ap AudioProfile, seek time.Duration) AudioProfile {
	p.seek = seek
	return p
}

func parseProfile(ap AudioProfile, in string) (string, []string, error) {
	parts, err := shlex.Split(profile.exec)
	if err != nil {
		return "", nil, fmt.Errorf("split command : %w", err)
	}

	if len(parts) == 0 {
		return "", nil, ErrNoProfileParts
	}

	name, err := exec.LookPath(parts[0])
	if err != nil {
		return "", nil, fmt.Errorf("find name : %w", err)
	}

	var args []string
	for _, p := range parts[1:] {
		switch p {
		case "<file>":
			args = append(args, in)
		case "<seek>":
			args = append(args, fmt.Sprintf("%dus", profile.Seek().Microseconds()))
		case "<bitrate>":
			args = append(args, fmt.Sprintf("%dk", profile.BitRate()))
		default:
			args = append(args, p)
		}
	}

	return name, args, nil
}

func startFFMpegProcess(ap AudioProfile, buf io.Reader) <-chan error {

	done := make(chan error)
	go func() {
		err := ffmpeg.Input("pipe:", ffmpeg.KwArgs{"format":"rawaudio"}).Output(ap.mime).OverWirteOutput().WithInput(buf).Run()
		log.Println("ffmpeg process done")
		done <- err
		close(done)
	}
}