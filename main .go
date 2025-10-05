package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vieiraes/opportunities/config"
)

func main() {
	dbConnection := config.SetupDB()
	defer dbConnection.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	log.Println("Server starting on port 3344...")
	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
	// return dbConnection
	//pergunta: o mian() pode ser uam funcao que nao retorna nada?
	//respposta: Sim, a função main em Go pode ser definida como uma função que não retorna nada. Na verdade, a assinatura correta da função main é func main() sem nenhum tipo de retorno. A função main é o ponto de entrada do programa e não deve retornar valores. Portanto, a forma correta de definir a função main é exatamente como você fez na segunda versão:

} // a funcao main() nao deve retornar valores no GO
