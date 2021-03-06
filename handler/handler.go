package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// root access
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// request halaman root
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//jooin file
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	// error testing index html
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data static untuk read di html untuk interface bisa mereturnk bebas tipe data
	// data := map[string]interface{}{
	// 	"title":   "i'm learn golang web",
	// 	"content": "I'm learning golang web with zaky",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 20200202020, Stock: 12},
		{ID: 2, Name: "Damkar", Price: 40000000020, Stock: 7},
		{ID: 3, Name: "Toyota", Price: 894289321020, Stock: 1},
		{ID: 4, Name: "xpander", Price: 30030302020, Stock: 8},
	}

	// error testing res req
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

// /hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world, saya sedang belajar"))
}

// /mario
func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello mario, sedang apa?"))
}

// /product?id=1
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//w.Write([]byte("Product Page"))

	// fmt.Fprintf(w, "Product page: %d", idNumb)

	data := map[string]interface{}{
		"content": idNumb,
	}

	// template can use for call page
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	// error testing index html
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// error testing res req
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

// func for get or post method ~~harus huruf kapital untuk eksport func
func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method //get or post

	switch method {
	case "GET":
		w.Write([]byte("ini adalh Get"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "error is heppening, keep calm", http.StatusBadRequest)
	}

}
