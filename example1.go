package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It was a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Kamol Hasan", "iphone X", true},
		{"Masudur Rahman", "Nokia 3600", false},
		{"Fahim Abrar", "", false},
	}
	t := template.Must(template.New("letter").Parse(letter))

	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Fatalf("t.Execute() failed with %s\n", err)
		}
	}
}
