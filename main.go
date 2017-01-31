package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	var err error
	// viper.BindEnv("cert", "SERVING_CERT")
	// viper.BindEnv("key", "SERVING_KEY")
	// viper.SetDefault("ip", "127.0.0.1")
	// viper.SetDefault("port", "8080")

	d1 := []byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go-Web-Go</title>
</head>
<body>
    <h1>Golang OpenShift Container Platform Storage Test</h1>
</body>
</html>`)
	err = ioutil.WriteFile("/data/home.html", d1, 0644)
	if err != nil {
		panic(err)
	}

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
	t, _ := template.ParseFiles("/data/home.html")
	t.Execute(w, nil)
}
