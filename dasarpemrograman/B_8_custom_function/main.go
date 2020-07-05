package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var funcMap = template.FuncMap{
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"avg": func(n ...int) int {
		var total = 0
		for _, each := range n {
			total += each
		}
		return total / len(n)
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl1 = template.Must(template.New("view.html").
			Funcs(funcMap).
			ParseFiles("view.html"))
		if err := tmpl1.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

// Berikut merupakan penjelasan step-by-step mengenai kode panjang untuk parsing dan rendering template di atas.
// 1. Sebuah template disipakan dengan nama view.html. Pembuatan instance template dilakukan melalui fungsi template.New().
// 2. Fungsi custom yang telah kita buat, diregistrasikan agar dikenali oleh template tersebut. Bisa dilihat pada pemanggilan method Funcs().
// 3. Setelah itu, lewat method ParseFiles(), view view.html di-parsing. Akan dicari dalam file tersebut apakah ada template yang didefinisikan dengan nama view.html. Karena di dalam template view tidak ada deklarasi template sama sekali ({{template "namatemplate"}}), maka akan dicari view yang namanya adalah view.html. Keseluruhan isi view.html akan dianggap sebagai sebuah template dengan nama template adalah nama file itu
// sendiri.
