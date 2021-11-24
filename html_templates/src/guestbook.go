package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines

}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func InvoiceHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("bill.html")
	check(err)
	bill := Invoice{
		Name:    "Mary Gibbs",
		Paid:    true,
		Charges: []float64{23.19, 1.13, 42.79},
		Total:   67.11,
	}
	err = html.Execute(writer, bill)
	check(err)

}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	//write the text to file
	_, err = fmt.Fprintln(file, signature)
	check(err)
	//writes to server/terminal
	// _, err = writer.Write([]byte(signature))
	// check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)

}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	http.HandleFunc("/invoice", InvoiceHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}
