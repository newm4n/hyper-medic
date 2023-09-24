package types

import "time"

type DeceaseStatus struct {
	Deceased    bool
	DeceaseTime time.Time
}
