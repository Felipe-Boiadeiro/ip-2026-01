package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com banco de dados: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}

func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS pacientes (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(150) NOT NULL,
		cpf VARCHAR(14) NOT NULL UNIQUE,
		data_nascimento DATE NOT NULL,
		telefone VARCHAR(20),
		email VARCHAR(150),
		tipo_sanguineo VARCHAR(5),
		convenio VARCHAR(100),
		observacoes TEXT,
		criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela de pacientes: %v", err)
	}

	fmt.Println("Tabela 'pacientes' verificada/criada com sucesso!")
}
