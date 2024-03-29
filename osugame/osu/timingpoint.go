package osu

import (
	"errors"
	"strconv"
	"strings"
)

type TimingPoint struct { // delimiter,
	Time        int     `json:"time"`
	BeatLength  float64 `json:"beatLength"`
	Meter       int     `json:"meter"`
	SampleSet   int     `json:"sampleSet"`   // nofloat
	SampleIndex int     `json:"sampleIndex"` // nofloat
	Volume      int     `json:"volume"`
	Uninherited bool    `json:"uninherited"`
	Effects     int     `json:"effects"` // nofloat
}

func newTimingPoint(line string) (TimingPoint, error) {
	// time,beatLength,meter,sampleSet,sampleIndex,volume,uninherited,effects
	var tp TimingPoint
	vs := strings.Split(line, `,`)
	if len(vs) < 8 {
		return tp, errors.New("invalid timing point: not enough length")
	}
	{
		f, err := strconv.ParseFloat(vs[0], 64)
		if err != nil {
			return tp, err
		}
		tp.Time = int(f)
	}
	{
		f, err := strconv.ParseFloat(vs[1], 64)
		if err != nil {
			return tp, err
		}
		tp.BeatLength = f
	}
	{
		f, err := strconv.ParseFloat(vs[2], 64)
		if err != nil {
			return tp, err
		}
		tp.Meter = int(f)
	}
	{
		i, err := strconv.Atoi(vs[3])
		if err != nil {
			return tp, err
		}
		tp.SampleSet = i
	}
	{
		i, err := strconv.Atoi(vs[4])
		if err != nil {
			return tp, err
		}
		tp.SampleIndex = i
	}
	{
		f, err := strconv.ParseFloat(vs[5], 64)
		if err != nil {
			return tp, err
		}
		tp.Volume = int(f)
	}
	{
		b, err := strconv.ParseBool(vs[6])
		if err != nil {
			return tp, err
		}
		tp.Uninherited = b
	}
	{
		i, err := strconv.Atoi(vs[7])
		if err != nil {
			return tp, err
		}
		tp.Effects = i
	}
	return tp, nil
}

func (tp TimingPoint) IsInherited() bool { return !tp.Uninherited }

func (tp TimingPoint) BPM() (bpm float64, ok bool) {
	if !tp.Uninherited {
		return 0, false
	}
	return 1000 * 60 / tp.BeatLength, true
}

// SpeedFactor returns a speed factor. The standard value is 1.
func (tp TimingPoint) SpeedFactor() (speed float64, ok bool) {
	if tp.Uninherited {
		return 0, false
	}
	return 100 / (-tp.BeatLength), true
}

func (tp TimingPoint) IsKiai() bool { return tp.Effects&1 != 0 }
func (tp TimingPoint) isFirstBarOmitted() bool {
	return tp.Effects&(1<<3) != 0
}
