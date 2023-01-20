package streaming

import (
	"github.com/pyroscope-io/pyroscope/pkg/storage/segment"
)

const (
	profSampleType        = 1
	profSample            = 2
	profMapping           = 3
	profLocation          = 4
	profFunction          = 5
	profStringTable       = 6
	profDropFrames        = 7
	profKeepFrames        = 8
	profTimeNanos         = 9
	profDurationNanos     = 10
	profPeriodType        = 11
	profPeriod            = 12
	profComment           = 13
	profDefaultSampleType = 14

	stType = 1
	stUnit = 2

	locID   = 1
	locLine = 4

	lineFunctionID = 1

	funcID   = 1
	funcName = 2

	sampleLocationID = 1
	sampleValue      = 2
	sampleLabel      = 3

	labelKey = 1
	labelStr = 2
)

var (
	profileIDLabel = []byte(segment.ProfileIDLabelName)
)

type valueType struct {
	Type int64
	unit int64

	resolvedType string
	resolvedUnit string
}
type function struct {
	id       uint64
	name     int32
	filename int32
}

//const noFunction = 0xffffffffffffffff

type location struct {
	id uint64
	// packed from << 32 | to into values
	linesRef uint64
}

type line struct {
	functionID uint64
	line       int64
}

// from,to into profile buffer
type istr uint64

type sample struct {
	tmpValues []int64
	// k<<32|v
	//type labelPacked uint64
	tmpLabels   []uint64
	tmpStack    [][]byte
	tmpStackLoc []uint64
}

func (s *sample) reset() {
	// 64 is max pc for golang + speculative number of inlines
	if s.tmpStack == nil {
		s.tmpStack = make([][]byte, 0, 64+8)
		s.tmpStackLoc = make([]uint64, 0, 64+8)
		s.tmpValues = make([]int64, 0, 4)
		s.tmpLabels = make([]uint64, 0, 4)
	} else {
		s.tmpStack = s.tmpStack[:0]
		s.tmpStackLoc = s.tmpStackLoc[:0]
		s.tmpValues = s.tmpValues[:0]
		s.tmpLabels = s.tmpLabels[:0]
	}
}