package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	//get function which satisfies io.writer which writes the value
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {

	executeTemplate("Dot is: {{.}}\n", true)
	executeTemplate("Dot is: {{.}}\n", false)

	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", false)

	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})

	templateText = "Prices:\n{{range .}}{{.}}\n{{end}}\n"
	executeTemplate(templateText, []float64{1.25, .99, 27})
	executeTemplate(templateText, nil)

	templateText = "Name: {{.Name}}\nCount:{{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})

	templateText = "Name: {{.Name}}{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "Abhilash Nair", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "Jason Tiel", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)

}
