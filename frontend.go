package codenames

import (
	"math/rand"
	"net/http"
	"strings"
)

const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <title>Codenames</title>
        <script src="/js/lib/browser.min.js"></script>
        <script src="/js/lib/react.min.js"></script>
        <script src="/js/lib/react-dom.min.js"></script>
        <script src="/js/lib/jquery-3.0.0.min.js"></script>

        <script type="text/babel">
             window.autogeneratedGameID = "{{.AutogeneratedGameID}}";
        </script>

        {{range .JSScripts}}
            <script type="text/babel" src="/js/{{ . }}"></script>
        {{end}}
        {{range .Stylesheets}}
            <link rel="stylesheet" type="text/css" href="/css/{{ . }}" />
        {{end}}

        <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
    </head>
    <body>
        <div id="app">
			<h1>Codenames board game generator</h1>
			<p>Play Codenames across multiple devices: computers, tablets, phones, etc.
			Board state automatically syncs between devices.</p>
		</div>
        <script type="text/babel">
            ReactDOM.render(<window.App />, document.getElementById('app'));
        </script>
    </body>
</html>
`

type templateParameters struct {
	AutogeneratedGameID string
	JSLibs              []string
	JSScripts           []string
	Stylesheets         []string
}

func (s *Server) handleIndex(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(rw, req)
		return
	}

	first := s.words[rand.Intn(len(s.words))]
	second := s.words[rand.Intn(len(s.words))]

	err := s.tpl.Execute(rw, templateParameters{
		AutogeneratedGameID: strings.Replace(first+"-"+second, " ", "", -1),
		JSLibs:              s.jslib.RelativePaths(),
		JSScripts:           s.js.RelativePaths(),
		Stylesheets:         s.css.RelativePaths(),
	})
	if err != nil {
		http.Error(rw, "error rendering", http.StatusInternalServerError)
	}
}
