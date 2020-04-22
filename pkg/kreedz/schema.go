package kreedz

import (
	"net/http"
)

type Organization uint

const (
	_ Organization = iota
	XtremeJumps
	CosyClimbing
	WorldSurf
	DebugWorldRecord = 0
)

type WorldRecord struct {
	Organization     Organization
	recordFileHeader http.Header
	mapNameCache     map[string]*RecordInfo
	MapSaveDir       string
}

type RecordInfo struct {
	MapName string
	Holder  string
	Region  string
	Time    float64
	Route   string
}
