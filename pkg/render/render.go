package render

import (
	"bytes"
	"text/template"
)

type Render struct {
	data map[string]interface{}
}

func NewRender(opts Options) *Render {
	return &Render{
		data: opts.toMap(),
	}
}

func (r *Render) Render(tmpl string) (string, error) {
	tpl, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, r.data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
