/////////////////////////////////////////////////////
// Simple Go Echo Server
// http://blog.roomanna.com/05-20-2011/simple-go-server
// https://github.com/kurrik/appengine-samples/blob/master/go-echo/echo/echo.go

package echo

import (
	"html/template"
	"net/http"
	"fmt"
)


func Init_echo() {
  http.HandleFunc("/dest", posthandler)
  http.HandleFunc("/echoform", formhandler)
}


func formhandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, `<!DOCTYPE html>
<html>
<body>
<form method="POST" action="/dest">
<input type="text" name="content" placeholder="Put content here" />
<button>Submit</button>
</form>
</body>
</html>`)
}

func posthandler (w http.ResponseWriter, r *http.Request) {
  data := map[string] interface{} {
    "Content": r.FormValue("content"),
    "Host": r.Host,
  }
  
  if err := postTemplate.Execute(w, data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

var postTemplate = template.Must(template.New("post").Parse(postTemplateHTML))
const postTemplateHTML = `
<!DOCTYPE html>
<html>
<body>
Post Content: {{.Content}}<br />
Host: {{.Host}}<br />

<p/><a href="/">Back home</a>
</body>
</html>`


