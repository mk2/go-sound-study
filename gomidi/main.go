package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"
)

type printer struct{}

func (pr printer) noteOn(p *reader.Position, channel, key, vel uint8) {
	fmt.Printf("Track: %v Pos: %v NoteOn (ch %v: key %v vel: %v)\n", p.Track, p.AbsoluteTicks, channel, key, vel)
}

func (pr printer) noteOff(p *reader.Position, channel, key, vel uint8) {
	fmt.Printf("Track: %v Pos: %v NoteOff (ch %v: key %v)\n", p.Track, p.AbsoluteTicks, channel, key)
}

func main() {
	dir, _ := os.Getwd()
	f := filepath.Join(dir, "smf-test.mid")

	if _, err := os.Stat(f); os.IsExist(err) {
		os.Remove(f)
	}

	// defer os.Remove(f)

	var p printer

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

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(p.noteOn),
		reader.NoteOff(p.noteOff),
	)

	err = reader.ReadSMFFile(rd, f)

	if err != nil {
		fmt.Printf("could not read SMF file %v\n", f)
	}

	// Output: Track: 0 Pos: 0 NoteOn (ch 11: key 120 vel: 50)
	// Track: 0 Pos: 120 NoteOff (ch 11: key 120)
	// Track: 0 Pos: 360 NoteOn (ch 11: key 125 vel: 50)
	// Track: 0 Pos: 380 NoteOff (ch 11: key 125)
	// Track: 1 Pos: 0 NoteOn (ch 2: key 120 vel: 50)
	// Track: 1 Pos: 60 NoteOff (ch 2: key 120)
}
