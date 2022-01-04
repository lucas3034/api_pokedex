package functions

import (
	"fmt"
	"net/http"
)

func Mensagem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindos a Pokedéx")
}


//Não estou conseguindo chamar essa função no main
