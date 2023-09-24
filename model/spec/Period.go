package modelspec

import "time"

type IPeriod interface {
	GetBegin() time.Time
	GetEnd() time.Time
}
