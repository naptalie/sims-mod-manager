package models

import (
	"time"
)

// Mod struct to hold mod information
type Mod struct {
	Name    string
	Path    string
	Version string
	Size    int64
	Updated time.Time
}