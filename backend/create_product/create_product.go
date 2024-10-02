package create_product

import (
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB max request per form

	if err != nil {
		fmt.Println("Error has occured with reading form", err.Error()) // message displayed only on backend for debugging
		http.Error(w, "Problem sa ucitavanjem Vaseg zahtjeva", http.StatusBadRequest)
		return
	}

	sifra := r.FormValue("sifra")
	naziv := r.FormValue("naziv")
	barkod := r.FormValue("barkod")
	gln := r.FormValue("gln")
	dostupnost := r.FormValue("dostupnost")
	minimalna_kolicina := r.FormValue("minimalna_kolicina")

	if sifra == "" {
		fmt.Println("Sifra is not entered")
		http.Error(w, "Niste unijeli sifru", http.StatusBadRequest)
		return
	} else if naziv == "" {
		fmt.Println("Naziv is not entered")
		http.Error(w, "Niste unijeli naziv", http.StatusBadRequest)
		return
	} else if barkod == "" {
		fmt.Println("Barkod is not entered")
		http.Error(w, "Niste unijeli barkod", http.StatusBadRequest)
		return
	} else if gln == "" {
		fmt.Println("Gln is not entered")
		http.Error(w, "Niste unijeli gln", http.StatusBadRequest)
		return
	} else if dostupnost == "" {
		fmt.Println("Dostupnost is not entered")
		http.Error(w, "Niste unijeli dostupnost", http.StatusBadRequest)
		return
	} else if minimalna_kolicina == "" {
		fmt.Println("Minimalna_kolicina is not entered")
		http.Error(w, "Niste unijeli minimalna_kolicina", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Produkt kreiran")
	fmt.Println("Product created")
	fmt.Println(sifra, naziv, barkod, gln, dostupnost, minimalna_kolicina)
}
