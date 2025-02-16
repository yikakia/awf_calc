package i18n

import (
	"embed"

	"fyne.io/fyne/v2/lang"
)

//go:embed translations
var translations embed.FS

func init() {
	err := lang.AddTranslationsFS(translations, "translations")
	if err != nil {
		panic(err)
	}
}
