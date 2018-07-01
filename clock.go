package main

import (
    "time"
)

type Clock interface {
    Now() time.Time
}

type frozenClock struct{}

func (frozenClock) Now() time.Time {
    t, _ := time.Parse("2006-01-02T15:04:05", "2015-01-01T00:00:00")

    return t
}

type realClock struct{}

func (realClock) Now() time.Time {
    return time.Now()
}
