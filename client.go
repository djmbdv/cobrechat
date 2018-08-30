package main

import (
"net/http"
"log"
 "fmt"
 "io/ioutil" )
func main() {
	 res, err := http.Get("https://diakonia-ec.org/test.php")
	if err != nil {
	    log.Fatal(err)
	}
	contenido, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("%s", contenido)
	 
}