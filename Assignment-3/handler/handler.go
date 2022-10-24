package handler

import (
	"html/template"
	d "microservice/service"
	"net/http"
	"path"
)

func getdescwt(i int) string {
	if i < 5 {
		return "AMAN"
	} else if i >= 6 && i <= 8 {
		return "SIAGA"
	} else {
		return "BAHAYA"
	}

}

func getdescwd(i int) string {
	if i < 6 {
		return "AMAN"
	} else if i >= 7 && i <= 15 {
		return "SIAGA"
	} else {
		return "BAHAYA"
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {

	var html = path.Join("index.html")

	var tmpl, err = template.ParseFiles(html)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	getdata := d.GetService()
	wtrnumber := getdata.Status.Water
	wdnumber := getdata.Status.Wind

	alertiwater := getdescwt(wtrnumber)
	alertwind := getdescwd(wdnumber)

	var Desc = map[string]interface{}{
		"water":     wtrnumber,
		"wind":      wdnumber,
		"waterDesc": alertiwater,
		"windDesc":  alertwind,
	}

	err = tmpl.Execute(w, Desc)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
