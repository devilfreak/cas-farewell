package handler

import "time"

type File struct {
	Bule       map[Level]time.Duration `json:"bule"`
	Pb         map[string]Run          `json:"pb,omitempty"`
	Defaultrun string                  `json:"default_run,omitempty"`
}

type Run struct {
	Times      map[Level]time.Duration `json:"times"`
	Levelnames map[Level]string        `json:"level_names,omitempty"`
}

type Level struct {
	Chapter Chapter
	Side    Side
}

type Settings struct {
	Settings map[string]string `json:"settings,omitempty"`
}

type Side int

type Chapter int