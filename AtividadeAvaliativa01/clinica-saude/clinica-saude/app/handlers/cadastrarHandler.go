package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"clinica-saude/app/utils"
)

func CadastrarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método HTTP não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao processar o formulário. Verifique os dados enviados.", http.StatusBadRequest)
		return
	}

	nome := strings.TrimSpace(r.FormValue("nome"))
	cpf := strings.TrimSpace(r.FormValue("cpf"))
	dataNascimento := strings.TrimSpace(r.FormValue("data_nascimento"))
	telefone := strings.TrimSpace(r.FormValue("telefone"))
	email := strings.TrimSpace(r.FormValue("email"))
	tipoSanguineo := strings.TrimSpace(r.FormValue("tipo_sanguineo"))
	convenio := strings.TrimSpace(r.FormValue("convenio"))
	observacoes := strings.TrimSpace(r.FormValue("observacoes"))

	if err := utils.ValidarPacienteCompleto(nome, cpf, dataNascimento, email, telefone, tipoSanguineo); err != nil {
		utils.SendHTMLResponse(w, "Erro na Validação", utils.ResponseError, err.Error(), "")
		return
	}

	paciente := utils.Paciente{
		Nome:           nome,
		CPF:            cpf,
		DataNascimento: dataNascimento,
		Telefone:       telefone,
		Email:          email,
		TipoSanguineo:  tipoSanguineo,
		Convenio:       convenio,
		Observacoes:    observacoes,
	}

	err = utils.CriarPaciente(paciente)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			utils.SendHTMLResponse(w, "CPF Duplicado", utils.ResponseError, "Este CPF já está cadastrado no sistema.", "")
			return
		}
		utils.SendHTMLResponse(w, "Erro ao Cadastrar", utils.ResponseError, fmt.Sprintf("Não foi possível cadastrar o paciente: %v", err), "")
		return
	}

	detalhes := map[string]string{
		"Nome":               nome,
		"CPF":                cpf,
		"Data de Nascimento": dataNascimento,
		"Telefone":           telefone,
		"Email":              email,
		"Tipo Sanguíneo":     tipoSanguineo,
		"Convênio":           convenio,
	}

	utils.SendDetailedResponse(w, "Paciente Cadastrado", utils.ResponseSuccess, "Paciente cadastrado com sucesso no sistema!", detalhes)
}
