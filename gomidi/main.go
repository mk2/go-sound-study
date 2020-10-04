package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/gomidi/midi/writer"
)

func main() {
	dir, _ := os.Getwd()
	f := filepath.Join(dir, "smf-test.mid")

	if _, err := os.Stat(f); os.IsExist(err) {
		os.Remove(f)
	}

	err := writer.WriteSMF(f, 1, func(wr *writer.SMF) error {
		writer.TrackSequenceName(wr, "doremi")
		writer.TempoBPM(wr, 120)
		writer.ProgramChange(wr, 1)

		writer.NoteOn(wr, 60, 100)
		wr.SetDelta(960)
		writer.NoteOff(wr, 60)

		writer.NoteOn(wr, 62, 100)
		wr.SetDelta(960)
		writer.NoteOff(wr, 62)

		writer.NoteOn(wr, 64, 100)
		wr.SetDelta(960)
		writer.NoteOff(wr, 64)
		writer.EndOfTrack(wr)
		return nil
	})

	if err != nil {
		fmt.Printf("could not write SMF file %v\n", f)
		return
	}

}
