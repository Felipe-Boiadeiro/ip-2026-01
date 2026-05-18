package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"clinica-saude/app/utils"
)

func BuscarPacienteHandler(w http.ResponseWriter, r *http.Request) {
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
		utils.SendHTMLResponse(w, "Campo Obrigatório", utils.ResponseError, "CPF é obrigatório para realizar a busca.", "")
		return
	}

	if !utils.ValidarCPF(cpf) {
		utils.SendHTMLResponse(w, "CPF Inválido", utils.ResponseError, "O CPF fornecido possui um formato inválido.", "")
		return
	}

	paciente, err := utils.BuscarPacientePorCPF(cpf)
	if err != nil {
		utils.SendHTMLResponse(w, "Paciente Não Encontrado", utils.ResponseError, fmt.Sprintf("Nenhum paciente encontrado com o CPF: %s", cpf), "")
		return
	}

	detalhes := map[string]string{
		"ID":                 fmt.Sprintf("%d", paciente.ID),
		"Nome":               paciente.Nome,
		"CPF":                paciente.CPF,
		"Data de Nascimento": paciente.DataNascimento,
		"Telefone":           paciente.Telefone,
		"Email":              paciente.Email,
		"Tipo Sanguíneo":     paciente.TipoSanguineo,
		"Convênio":           paciente.Convenio,
		"Cadastrado em":      paciente.CriadoEm,
	}

	if paciente.Observacoes != "" {
		detalhes["Observações"] = paciente.Observacoes
	}

	utils.SendDetailedResponse(w, "Dados do Paciente", utils.ResponseInfo, "Informações do paciente encontradas no sistema.", detalhes)
}
