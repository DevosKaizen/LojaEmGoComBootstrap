package routes

import (
	"LOJAEMGO/controllers"
	"net/http"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)

}
