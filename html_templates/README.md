## html templates

* insert data into html with package
* powerful way to include data in your app's HTML responses
* `html/template` - template package used to parse files
  * `text/template` -same but without security features of html (NOTE:check go page 462)
    * **actions** - actions to template text like inserting data which is identified in text as `{{.}}`
      * if actions can also be run using `{{if .}} {{end}}`
      * range actions can also be run using `{{range .}} {{end}}`
      * can also pass struct fields to actions and any data passed to action can be identifed by the inserted struct as such `.Name`
* `http.ResponseWriter` and `os.StdOut` statisfy io.Writer interface which requires a `Write` function
  * any data can be passed including nil
* **http redirects** - use `http.Redirect` to redirect back to a page
