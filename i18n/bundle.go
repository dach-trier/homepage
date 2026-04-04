package i18n

import "fmt"
import "io/fs"
import "encoding/json"
import "golang.org/x/text/language"

type Bundle struct {
	translations map[language.Tag]map[string]string
}

func NewBundle() *Bundle {
	b := &Bundle{}
	b.translations = make(map[language.Tag]map[string]string)
	return b
}

func (b *Bundle) AddTranslation(lang language.Tag, id string, translation string) error {
	if _, exists := b.translations[lang]; !exists {
		b.translations[lang] = make(map[string]string)
	}

	if _, exists := b.translations[lang][id]; exists {
		return ErrTranslationDuplicate{Lang: lang, Id: id}
	}

	b.translations[lang][id] = translation

	return nil
}

func (b *Bundle) MustAddTranslation(lang language.Tag, id string, translation string) {
	err := b.AddTranslation(lang, id, translation)
	if err != nil {
		panic(err)
	}
}

func (b *Bundle) loadRawTranslations(lang language.Tag, parentId string, raw any) error {
	switch value := raw.(type) {
	case map[string]any:
		for id, child := range value {
			nextId := id
			if len(parentId) > 0 {
				nextId = parentId + "." + id
			}
			err := b.loadRawTranslations(lang, nextId, child)

			if err != nil {
				return err
			}
		}

	case string:
		if len(parentId) == 0 {
			return fmt.Errorf("unexpected string value")
		}

		id, translation := parentId, value
		err := b.AddTranslation(lang, id, translation)
		if err != nil {
			return err
		}

	case nil:
		return fmt.Errorf("unexpected nil value")
	default:
		return fmt.Errorf("unexpected value")
	}

	return nil
}

func (b *Bundle) LoadRawTranslations(lang language.Tag, raw any) error {
	err := b.loadRawTranslations(lang, "", raw)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bundle) MustLoadRawTranslations(lang language.Tag, raw any) {
	err := b.LoadRawTranslations(lang, raw)
	if err != nil {
		panic(err)
	}
}

func (b *Bundle) LoadJsonTranslationsFS(lang language.Tag, filesystem fs.FS, path string) error {
	file, err := fs.ReadFile(filesystem, path)
	if err != nil {
		return err
	}

	var raw any
	err = json.Unmarshal(file, &raw)
	if err != nil {
		return err
	}

	err = b.LoadRawTranslations(lang, raw)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bundle) MustLoadJsonTranslationsFS(lang language.Tag, filesystem fs.FS, path string) {
	err := b.LoadJsonTranslationsFS(lang, filesystem, path)
	if err != nil {
		panic(err)
	}
}
