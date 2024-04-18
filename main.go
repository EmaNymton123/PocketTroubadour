package main

import (
	"net/http"
	"log"
	"github.com/EmaNymton123/PocketTroubadour/controllers"
	"flag"
)

func main(){
	var listenFlag *string = flag.String("listen", ":8080", "description adresse d'écoute")
	flag.Parse()
	s := &http.Server{
		Addr:           *listenFlag,
		Handler:        controllers.Routes(),
	}
	log.Printf("starting PocketTroubadour on %v", *listenFlag)
	unpeucommeonveut("vous qui lisez ceci.")

	for i := 0 ; i < 10 ; i++ {
		resfib := fibonacci(i)
		log.Print(i, resfib)
	}

	log.Fatal(s.ListenAndServe())
}
func unpeucommeonveut(nameparexemple string){
	log.Print("Bonjour ", nameparexemple)
}
func fibonacci(rang int) int {
	//f(n)=f(n-1)+f(n-2) pour >=2
	var res int
	//log.Printf("rang = %v", rang)
	if rang >= 2{
		res = fibonacci(rang-1)+fibonacci(rang-2)
	} else {
		res = 1
	}
	//log.Printf("rang = %v, res = %v", rang, res)
	return res
}
