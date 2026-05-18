package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"clinica-saude/app/utils"
)

func ExcluirPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método HTTP não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar o formulário.", http.StatusBadRequest)
		return
	}

	cpf := strings.TrimSpace(r.FormValue("cpf"))
	if cpf == "" {
		utils.SendHTMLResponse(w, "Campo Obrigatório", utils.ResponseError, "CPF é obrigatório para excluir um paciente.", "")
		return
	}

	if !utils.ValidarCPF(cpf) {
		utils.SendHTMLResponse(w, "CPF Inválido", utils.ResponseError, "O CPF fornecido possui um formato inválido.", "")
		return
	}

	err = utils.ExcluirPaciente(cpf)
	if err != nil {
		utils.SendHTMLResponse(w, "Erro ao Excluir", utils.ResponseError, fmt.Sprintf("Não foi possível excluir o paciente. %v", err), "")
		return
	}

	detalhes := map[string]string{
		"CPF Excluído": cpf,
		"Status":       "Paciente removido permanentemente do sistema",
	}

	utils.SendDetailedResponse(w, "Paciente Excluído", utils.ResponseWarning, "Paciente foi removido com sucesso do sistema.", detalhes)
}
