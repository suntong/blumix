/////////////////////////////////////////////////////
// Simple Go Upload Server
// initially from https://github.com/ljgww/web_server_example_in_Go_-golang-

package upload

import (
  "net/http"
  "html/template"
  "io"
  "io/ioutil"
	"log"
)

var uploadTemplate, _ = template.ParseFiles("upload/upload.html")
var errorTemplate, _ = template.ParseFiles("upload/error.html")

func check(err error) { if err != nil { panic(err) } }

func Init_upload() {
    http.HandleFunc("/upload", errorHandler(upload))
    http.HandleFunc("/view", errorHandler(view))
}

func errorHandler(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if e, ok := recover().(error); ok {
                 w.WriteHeader(500)
                 errorTemplate.Execute(w, e)
            }
        }()
        fn(w, r)
    }
}

func upload(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    uploadTemplate.Execute(w, nil)
    return
  }
  
  log.Println(">> Upload Handler")
  defer log.Println("<< Upload Handler")

	f, _, err := r.FormFile("image")
  check(err)
  defer f.Close()
  t, err := ioutil.TempFile("./uploading/", "image-")
  check(err)
  defer t.Close()
  _, err = io.Copy(t, f)
  check(err)
  http.Redirect(w, r, "/view?id="+t.Name()[16:], 302)
}

func view(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image")
    http.ServeFile(w, r, "uploading/image-"+r.FormValue("id"))
}

