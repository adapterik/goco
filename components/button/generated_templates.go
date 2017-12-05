// Package button is generated with ftmpl {{{v0.3.1}}}, do not edit!!!! */
package button

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// TMPLERRmarkup evaluates a template markup.tmpl
func TMPLERRmarkup() (string, error) {
	_template := "markup.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<button type="button"></button>`)

	return _ftmpl.String(), nil
}

// T_markup evaluates a template markup.tmpl
func T_markup() string {
	html, err := TMPLERRmarkup()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template markup.tmpl:" + err.Error())
	}
	return html
}
