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

	"github.com/Unknwon/goconfig"
)

var (
	defautLocale = "en-US"
	message      *goconfig.ConfigFile
)

// SetDefaultLocale sets the default language for localization.
func SetDefaultLocale(locale string) {
	defautLocale = locale
}

// SetMessage sets the message file for localization.
func SetMessage(filePath string) (err error) {
	message, err = goconfig.LoadConfigFile(filePath)
	return err
}

// A Locale describles the information of localization.
type Locale struct {
	CurrentLocale string
}

// Tr translate content to target language.
func Tr(locale, format string, args ...interface{}) string {
	if locale != defautLocale {
		format = message.MustValue(locale, format)
	}
	return fmt.Sprintf(format, args...)
}

// Tr translate content to target language.
func (l Locale) Tr(format string, args ...interface{}) string {
	return Tr(l.CurrentLocale, format, args...)
}
