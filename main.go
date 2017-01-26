package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var err error
	// viper.BindEnv("cert", "SERVING_CERT")
	// viper.BindEnv("key", "SERVING_KEY")
	// viper.SetDefault("ip", "127.0.0.1")
	// viper.SetDefault("port", "8080")

	http.HandleFunc("/", home)

	bind := fmt.Sprintf(":%s", "8024")
	fmt.Printf("listening on %s...", bind)
	err = http.ListenAndServe(bind, nil)
	//err = http.ListenAndServeTLS(bind, viper.GetString("cert"), viper.GetString("key"), nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}
