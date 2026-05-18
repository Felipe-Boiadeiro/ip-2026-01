package main

import (
	"fmt"
	"log"
	"net/http"

	"clinica-saude/app/handlers"
	"clinica-saude/app/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	utils.ConnectToDB()
	defer utils.DB.Close()

	utils.CreateTable()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/forms/cadastrarPaciente.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/forms/cadastrarPaciente.html")
	})
	http.HandleFunc("/forms/buscarPaciente.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/forms/buscarPaciente.html")
	})
	http.HandleFunc("/forms/atualizarPaciente.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/forms/atualizarPaciente.html")
	})
	http.HandleFunc("/forms/excluirPaciente.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/forms/excluirPaciente.html")
	})

	http.HandleFunc("/api/cadastrar", handlers.CadastrarPacienteHandler)
	http.HandleFunc("/api/buscar", handlers.BuscarPacienteHandler)
	http.HandleFunc("/api/atualizar", handlers.AtualizarPacienteHandler)
	http.HandleFunc("/api/excluir", handlers.ExcluirPacienteHandler)

	addr := "127.0.0.1:3000"
	fmt.Printf("Servidor rodando em http://%s/\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
