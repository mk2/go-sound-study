package main

// #cgo LDFLAGS: -lwinmm
// #include <stdlib.h>
// #include <windows.h>
// #include <mmsystem.h>
import "C"
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"unsafe"
)

const (
	SRATE    = 44410
	PI       = 3.14159286
	B_TIME   = 1.0
	F0       = 440.0
	AMP      = 40.0
	DATA_LEN = int(SRATE * B_TIME)
)

func main() {
	var (
		hWave C.HWAVEOUT
		whdr  C.WAVEHDR
		wfe   C.WAVEFORMATEX
	)

	bWave := (*[DATA_LEN]byte)(C.malloc(C.ulonglong(DATA_LEN)))

	for cnt := 0; cnt < DATA_LEN; cnt++ {
		bWave[cnt] = byte(AMP * math.Sin(float64(2.0*PI*F0*float32(cnt)/SRATE)))
	}

	wfe.wFormatTag = C.WAVE_FORMAT_PCM
	wfe.nChannels = 1
	wfe.nSamplesPerSec = SRATE
	wfe.nAvgBytesPerSec = SRATE
	wfe.wBitsPerSample = 8
	wfe.nBlockAlign = wfe.nChannels * wfe.wBitsPerSample / 8

	C.waveOutOpen(&hWave, C.WAVE_MAPPER, &wfe, 0, 0, C.CALLBACK_NULL)

	whdr.lpData = C.LPSTR(unsafe.Pointer(bWave))
	whdr.dwBufferLength = C.ulong(DATA_LEN)
	whdr.dwFlags = C.WHDR_BEGINLOOP | C.WHDR_ENDLOOP
	whdr.dwLoops = 1

	C.waveOutPrepareHeader(hWave, &whdr, C.uint(unsafe.Sizeof(C.WAVEHDR{})))
	C.waveOutWrite(hWave, &whdr, C.uint(unsafe.Sizeof(C.WAVEHDR{})))

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enterを押したら終了します...")
	reader.ReadString('\n')
}
