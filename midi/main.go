package main

// #cgo windows LDFLAGS: -L. -lMIDIData
// #include <windows.h>
// #include "MIDIData.h"
import "C"

func main() {
	var midiData = C.MIDIData_Create(C.MIDIDATA_FORMAT0, 1, C.MIDIDATA_TPQNBASE, 120)
	var midiTrack = C.MIDIData_GetFirstTrack(midiData)
	C.MIDITrack_InsertTrackNameA(midiTrack, 0, C.CString("doremi"))
	C.MIDITrack_InsertTempo(midiTrack, 0, 60000000/120)
	C.MIDITrack_InsertProgramChange(midiTrack, 0, 0, 1)

	C.MIDITrack_InsertNote(midiTrack, 0, 0, 60, 100, 120)
	C.MIDITrack_InsertNote(midiTrack, 120, 0, 62, 100, 120)
	C.MIDITrack_InsertNote(midiTrack, 240, 0, 64, 100, 120)

	C.MIDITrack_InsertEndofTrack(midiTrack, 360)
	C.MIDIData_SaveAsSMFA(midiData, C.CString("doremi.midi"))
	C.MIDIData_Delete(midiData)
}
