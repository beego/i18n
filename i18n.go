// Copyright 2013 beego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package i18n is for app Internationalization and Localization.
package i18n

import (
	"fmt"
	"reflect"

	"github.com/Unknwon/goconfig"
)

var (
	locales localeStore
)

type locale struct {
	lang    string
	message *goconfig.ConfigFile
}

type localeStore []*locale

// Get locale from localeStore use specify lang string
func (d *localeStore) getLocale(lang string) (*locale, bool) {
	for _, l := range locales {
		if l.lang == lang {
			return l, true
		}
	}
	return nil, false
}

// Get target language string
func (d *localeStore) Get(lang, section, format string) (string, bool) {
	if locale, ok := d.getLocale(lang); ok {
		if section == "" {
			section = "common"
		}
		value, err := locale.message.GetValue(section, format)
		if err == nil {
			return value, true
		}
	}
	return "", false
}

// SetMessage sets the message file for localization.
func SetMessage(lang, filePath string) error {
	message, err := goconfig.LoadConfigFile(filePath)
	if err == nil {
		lc := new(locale)
		lc.lang = lang
		lc.message = message
		locales = append(locales, lc)
	}
	return err
}

// A Locale describles the information of localization.
type Locale struct {
	Lang string
}

// Tr translate content to target language.
func (l Locale) Tr(format string, args ...interface{}) string {
	return Tr(l.Lang, format, args...)
}

// Trs translate content to target language with specify section.
func (l Locale) Trs(section, format string, args ...interface{}) string {
	return Trs(l.Lang, section, format, args...)
}

// Tr translate content to target language.
func Tr(lang, format string, args ...interface{}) string {
	return tr(lang, "", format, args...)
}

// Trs translate content to target language with specify section.
func Trs(lang, section, format string, args ...interface{}) string {
	return tr(lang, section, format, args...)
}

func tr(lang, section, format string, args ...interface{}) string {
	value, ok := locales.Get(lang, section, format)
	if ok {
		format = value
	}

	if len(args) > 0 {
		params := make([]interface{}, 0, len(args))
		for _, arg := range args {
			if arg != nil {
				val := reflect.ValueOf(arg)
				if val.Kind() == reflect.Slice {
					for i := 0; i < val.Len(); i++ {
						params = append(params, val.Index(i).Interface())
					}
				} else {
					params = append(params, arg)
				}
			}
		}
		return fmt.Sprintf(format, params...)
	}
	return fmt.Sprintf(format)
}
