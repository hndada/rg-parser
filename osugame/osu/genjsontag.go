// +build ignore

package main

import (
	"fmt"
	"strings"
)

const src = `type Format struct {
	FormatVersion int
	General
	Editor
	Metadata
	Difficulty
	Events
	TimingPoints
	Colours
	HitObjects
}

type General struct { // delimiter:(space)
	AudioFilename            string
	AudioLeadIn              int
	AudioHash                string // deprecated
	PreviewTime              int
	Countdown                int // nofloat
	SampleSet                string
	StackLeniency            float64
	Mode                     int // nofloat
	LetterboxInBreaks        bool
	StoryFireInFront         bool // deprecated
	UseSkinSprites           bool
	AlwaysShowPlayfield      bool // deprecated
	OverlayPosition          string
	SkinPreference           string
	EpilepsyWarning          bool
	CountdownOffset          int
	SpecialStyle             bool
	WidescreenStoryboard     bool
	SamplesMatchPlaybackRate bool
}
type Editor struct { // delimiter:(space)
	Bookmarks       []int // delimiter,
	DistanceSpacing float64
	BeatDivisor     float64
	GridSize        int
	TimelineZoom    float64
}
type Metadata struct { // delimiter:
	Title         string
	TitleUnicode  string
	Artist        string
	ArtistUnicode string
	Creator       string
	Version       string
	Source        string
	Tags          []string // delimiter(space)
	BeatmapID     int      // nofloat
	BeatmapSetID  int      // nofloat
}
type Difficulty struct { // delimiter:
	HPDrainRate       float64
	CircleSize        float64
	OverallDifficulty float64
	ApproachRate      float64
	SliderMultiplier  float64
	SliderTickRate    float64
}
type Events []Event
type TimingPoints []TimingPoint
type Colours struct { // manual
	Combos              [8]color.RGBA
	SliderTrackOverride color.RGBA
	SliderBorder        color.RGBA
}
type HitObjects []HitObject

type HitObject struct { // delimiter,
	X            int
	Y            int
	Time         int
	NoteType     int          // nofloat
	HitSound     int          // nofloat
	EndTime      int          // optional
	SliderParams SliderParams // optional
	HitSample    HitSample    // optional
}
type SliderParams struct { // delimiter,
	CurveType   string   // one letter
	CurvePoints [][2]int // delimiter| // delimiter: // slice of paired integers
	Slides      int
	Length      float64
	EdgeSounds  []int    // delimiter|
	EdgeSets    [][2]int // delimiter| // delimiter:
}
type HitSample struct { // delimiter:
	NormalSet   int // nofloat
	AdditionSet int // nofloat
	Index       int // nofloat
	Volume      int
	Filename    string
}

type TimingPoint struct { // delimiter,
	Time        int
	BeatLength  float64
	Meter       int
	SampleSet   int // nofloat
	SampleIndex int // nofloat
	Volume      int
	Uninherited bool
	Effects     int // nofloat
}
`

func main() {
	for _, line := range strings.Split(src, "\n") {
		vs := strings.Fields(line)
		if len(vs) >= 2 {
			switch vs[1] {
			case "string", "int", "float64", "bool":
				name := vs[0]
				tagName := strings.ToLower(string(name[0])) + name[1:]
				if len(name) >= 2 && name[:2] == "HP" {
					tagName = "hp" + tagName[2:]
				}
				for _, v := range vs[:2] {
					fmt.Printf("%s ", v)
				}
				fmt.Printf(" `json:\"%s\"`", tagName)
				for _, v := range vs[2:] {
					fmt.Printf("%s ", v)
				}
				fmt.Println()
			default:
				fmt.Println(line)
			}
		} else {
			fmt.Println(line)
		}
	}
}
