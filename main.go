package main

import (
	//"io"
	"os"
	"fmt"

	"github.com/gordonklaus/portaudio"
)

const (
	SampleFormatUnsupported int = 0
	SampleFormatF64 = 1
)

const (
	CodecUnsupported int = 20
	CodecWAV = 21
)

type AudioProps struct {
	Channels int
	SampleRate int
	Format int
	Codec int
}

// looking for the 'RIFF' header, and populate properties
func validateWAV(file *os.File) (props AudioProps, err error) {
	buffer := make([]byte, 4)
	//found, err = io.ReadAt(buffer, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	if string(buffer) == "RIFF" {
		props.Codec = CodecWAV
	}
	return
}

func main() {
	wav, err := os.Open("Periphery_Wildfire_rhythm_guitar_stem.wav")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wav.Close()

	portaudio.Initialize()
	defer portaudio.Terminate()
}
