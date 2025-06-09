package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// InitializeDatabase prepara e retorna a conexão com o banco de dados.
func InitializeDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Erro ao carregar o arquivo .env. Usando variáveis de ambiente do sistema.")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("Erro: A variável de ambiente DB_PATH não está definida.")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Erro ao preparar a conexão com o banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida.")
	return db
}