//DEUS é MARAVILHOSO!

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // _ significa "importar apenas para side-effects"
	// oo que siginifica importar com _ (underscore)? R: importa o pacote apenas para executar sua função init(), sem usar diretamente nenhum dos seus elementos. Isso é útil para registrar drivers ou plugins, como no caso do driver PostgreSQL aqui.
	// em que momentos eu devo usar o _ (underscore) ao importar um pacote? R: Use _ ao importar um pacote quando você precisa apenas dos efeitos colaterais do pacote, como a execução da função init(), sem precisar acessar diretamente suas funções, tipos ou variáveis. Isso é comum em casos como registro de drivers de banco de dados, plugins ou inicialização de configurações globais.
	// qual seria o equivalemnnte disso no javascript?
	// entendi.. entaio quendo quiser importar MAS ainda nao tiver usado ele no codigo E evitar que o GO apague a importacao, eu uso o _ (underscore) antes do nome do pacote
)

// Variáveis de configuração do banco (pacote os)
var dbHost string = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbUser = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbDatabaseName = os.Getenv("DB_DATABASE_NAME")

// SetupDB - função pública (letra maiúscula)
func SetupDB() *sql.DB {
	// Carrega .env (equivalente ao require('dotenv').config())
	err := godotenv.Load()
	if err != nil { // nil = null em JavaScript
		log.Fatal("Error loading .env file: ", err)
	}

	// var (
	// 	dbHost         = os.Getenv("DB_HOST")
	// 	dbPort         = os.Getenv("DB_PORT")
	// 	dbUser         = os.Getenv("DB_USERNAME")
	// 	dbPassword     = os.Getenv("DB_PASSWORD")
	// 	dbDatabaseName = os.Getenv("DB_DATABASE_NAME")
	// )

	// fmt.Sprintf = template string do JavaScript
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbDatabaseName)

	// Abre conexão com banco
	dbConnection, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Testa a conexão (como um ping)
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Database connected!")
	return dbConnection // retorna o ponteiro da conexão
}
