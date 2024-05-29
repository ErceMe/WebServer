package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Email, Address, Mood string
}

var Persons = []Person{
	{Email: "akubisa@gmail.com", Address: "Jl. Menuju Masa Depan", Mood: "Senang"},
	{Email: "galau@yahoo.com", Address: "Jl. Menuju Masa Lalu", Mood: "Sedih"},
	{Email: "genz2000@hotmail.com", Address: "Jl. Menuju Indonesia Cemas", Mood: "Cemas"},
	{Email: "mencoba@gmail.com", Address: "Jl. Sedang Berproses", Mood: "Bingung"},
	{Email: "bisagila@hotmail.com", Address: "Jl. Menuju Kebinasaan", Mood: "Gewlo"},
	{Email: "seblak@cewe.com", Address: "Jl. Menyala Abangkuhh", Mood: "hepihepihepi"},
}

func main() {
	http.HandleFunc("/index", getPerson)
	http.HandleFunc("/process", authorizeEmail)

	http.ListenAndServe(":1010", nil)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tpl.Execute(w, Persons)
		return
	}
	http.Error(w, "Errorr Invalid Method", http.StatusBadRequest)
}

func authorizeEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, err := template.ParseFiles("template.html")
		var email = r.FormValue("email")
		var message string
		for _, orang := range Persons {
			if email == orang.Email && err == nil {
				data := map[string]interface{}{
					"Email":   orang.Email,
					"Address": orang.Address,
					"Mood":    orang.Mood,
				}
				message = "berhasil"
				tmp, _ := template.ParseFiles("submit.html")
				tmp.Execute(w, data)
				break
			}
		}

		if message != "berhasil" {
			data2 := map[string]interface{}{
				"error":  "Email tidak ada",
				"error2": "-",
				"error3": "-",
			}
			tmp, _ := template.ParseFiles("submit.html")
			tmp.Execute(w, data2)
		}
	}
}
