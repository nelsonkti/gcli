package xfile

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

func WriteFile(path, tmpl string, args interface{}) (err error) {

	data, err := ParseTmpl(tmpl, args)
	if err != nil {
		return
	}

	return ioutil.WriteFile(path, data, 0755)
}

func ParseTmpl(tmpl string, args interface{}) ([]byte, error) {

	tmp, err := template.New("").Parse(tmpl)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = tmp.Execute(&buf, args); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
