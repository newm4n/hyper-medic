package modelspec

import "time"

type IAttachment interface {
	GetContentType() string
	GetLanguage() string
	GetData() string
	GetURL() string
	GetSize() uint64
	GetHash() string
	GetTitle() string
	GetCreation() time.Time
	GetHeight() uint
	GetWith() uint
	GetFrames() uint
	GetDuration() uint
	GetPages() uint
}
