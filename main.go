package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ip", "127.0.0.1")
	viper.SetDefault("port", "8080")

	http.HandleFunc("/", home)

	bind := fmt.Sprintf("%s:%s", viper.Get("ip"), viper.Get("port"))
	fmt.Printf("listening on %s...", bind)
	http.ListenAndServe(bind, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}
