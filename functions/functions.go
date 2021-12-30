package functions

import (
	"fmt"
	"net/http"
)

func mensagem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

//Não estou conseguindo chamar essa função no main
