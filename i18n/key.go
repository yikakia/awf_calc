package i18n

import (
	"fyne.io/fyne/v2/lang"
)

type Key string

// T translation string. fallback to itself
func (s Key) T() string {
	return lang.X(s.Raw(), s.Raw())
}

func (s Key) Raw() string {
	return string(s)
}
