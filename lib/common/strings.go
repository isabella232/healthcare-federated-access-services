// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"github.com/pborman/uuid"
)

const (
	RealmVariable = "{realm}"
)

func ListContains(list []string, find string) bool {
	for _, entry := range list {
		if entry == find {
			return true
		}
	}
	return false
}

func RequestAbstractPath(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")
	// Path starts with a "/", so first part is always empty.
	if len(parts) > 3 {
		parts[3] = RealmVariable
	}
	return strings.Join(parts, "/")
}

func GenerateGUID() string {
	return uuid.New()
}

func ParseGUID(in string) (uuid.UUID, error) {
	return uuid.Parse(in), nil
}

// JoinNonEmpty filters empty strings and joins remainder together.
func JoinNonEmpty(in []string, separator string) string {
	out := []string{}
	for _, v := range in {
		if len(v) > 0 {
			out = append(out, v)
		}
	}
	return strings.Join(out, separator)
}

// FilterStringsByPrefix filters returns only strings that do NOT have a given prefix.
func FilterStringsByPrefix(in []string, prefix string) []string {
	var out []string
	for _, v := range in {
		if !strings.HasPrefix(v, prefix) {
			out = append(out, v)
		}
	}
	return out
}

// ToTitle does some auto-formatting on camel-cased or snake-cased strings to make them look like titles.
func ToTitle(str string) string {
	out := ""
	l := 0
	for i, ch := range str {
		if unicode.IsUpper(ch) && i > 0 && str[i-1] != ' ' {
			out += " "
			l++
		} else if ch == '_' {
			out += " "
			l++
			continue
		}
		if l > 0 && out[l-1] == ' ' {
			ch = rune(unicode.ToUpper(rune(ch)))
		}
		out += string(ch)
		l++
	}
	return strings.Title(out)
}

// IsURL returns true if the format of the string appears to be a fully qualified URL
func IsURL(v string) bool {
	if len(v) < 7 || !(strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")) || !(strings.Contains(v, ".") || strings.Contains(v, "localhost")) || len(strings.Split(v, ":")) > 3 || strings.Contains(v, "//http") {
		return false
	}
	if _, err := url.Parse(v); err != nil {
		return false
	}
	return true
}

// ReplaceVariables replaces all substrings of the form "${var-name}"
// based on args like {"var-name":"var-value"}.
func ReplaceVariables(v string, args map[string]string) (string, error) {
	if idx := strings.Index(v, "${"); idx < 0 {
		return v, nil
	}
	parts := strings.Split(v, "${")
	out := parts[0]
	for i := 1; i < len(parts); i++ {
		p := strings.SplitN(parts[i], "}", 2)
		if len(p) < 2 {
			out += parts[i]
			continue
		}
		arg := p[0]
		val, ok := args[arg]
		if !ok {
			return "", fmt.Errorf("variable %q not defined", arg)
		}
		out += val + p[1]
	}
	return out, nil
}

// ExtractVariables returns a map of variable names found within an input string.
func ExtractVariables(v string) (map[string]bool, error) {
	args := make(map[string]bool)
	parts := strings.Split(v, "${")
	for i := 1; i < len(parts); i++ {
		p := strings.SplitN(parts[i], "}", 2)
		if len(p) < 2 {
			return nil, fmt.Errorf("variable mismatched brackets")
		}
		args[p[0]] = true
	}
	return args, nil
}
