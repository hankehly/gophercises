package cyoa

import (
	"encoding/json"
	"io"
)

var Template string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure - {{ .Title }}</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
    {{ range .Story }}
    <p>{{ . }}</p>
    {{ end }}
    <ul>
        {{ range .Options }}
        <li>
            <a href="/chapters/{{ .Arc }}">{{ .Text }}</a>
        </li>
        {{ end }}
    </ul>
</body>
</html>
`

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
