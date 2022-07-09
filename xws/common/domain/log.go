package domain

import "time"

type Log struct {
	Time        time.Time `json:"time"`
	Level       string    `json:"level"`
	Service     string    `json:"service"`
	Msg         string    `json:"msg"`
	FullContent string    `json:"fullContent"`
}
