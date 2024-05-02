package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	v8 "rogchap.com/v8go"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get the path of the request
		path := r.URL.Path

		// check if the path exists in the "pub" folder
		_, err := os.Stat("pub" + path)
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}

		if path == "/" {
			path = "/index.html"
		}

		// check if it ends with .html or .htm
		if path[len(path)-5:] == ".html" || path[len(path)-4:] == ".htm" {
			// open the file and read its contents
			bytes, err := os.ReadFile("pub" + path)
			if err != nil {
				http.Error(w, "Error reading file", 500)
				return
			}

			// convert the bytes to a string
			htmlContent := string(bytes)

			// The regular expression
			re := regexp.MustCompile(`(?i)<server>(.*?)</server>`)

			// Find all matches
			matches := re.FindAllStringSubmatch(htmlContent, -1)

			if len(matches) > 0 {
				// Loop over the matches and print them
				for _, match := range matches {
					fmt.Println(match[1])

					// create a new v8 context
					v8Context := v8.NewContext()

					// run the javascript code
					value, _ := v8Context.RunScript(match[1], "server.js")

					// replace the server tag with the result
					htmlContent = strings.Replace(htmlContent, match[0], value.String(), -1)
				}
			}

			// write the contents to the response writer
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(htmlContent))

			return
		} else {
			// return the contents with the correct mime type
			http.ServeFile(w, r, "pub"+path)
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
