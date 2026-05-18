package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"clinica-saude/app/utils"
)

func AtualizarPacienteHandler(w http.ResponseWriter, r *http.Request) {
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
	nome := strings.TrimSpace(r.FormValue("nome"))
	telefone := strings.TrimSpace(r.FormValue("telefone"))
	email := strings.TrimSpace(r.FormValue("email"))
	tipoSanguineo := strings.TrimSpace(r.FormValue("tipo_sanguineo"))
	convenio := strings.TrimSpace(r.FormValue("convenio"))
	observacoes := strings.TrimSpace(r.FormValue("observacoes"))

	if cpf == "" {
		utils.SendHTMLResponse(w, "Campo Obrigatório", utils.ResponseError, "CPF é obrigatório para atualizar um paciente.", "")
		return
	}

	if nome == "" {
		utils.SendHTMLResponse(w, "Campo Obrigatório", utils.ResponseError, "Nome é obrigatório para atualizar os dados.", "")
		return
	}

	if !utils.ValidarCPF(cpf) {
		utils.SendHTMLResponse(w, "CPF Inválido", utils.ResponseError, "O CPF fornecido possui um formato inválido.", "")
		return
	}

	if len(nome) < 3 {
		utils.SendHTMLResponse(w, "Nome Inválido", utils.ResponseError, "Nome deve ter pelo menos 3 caracteres.", "")
		return
	}

	if email != "" && !utils.ValidarEmail(email) {
		utils.SendHTMLResponse(w, "Email Inválido", utils.ResponseError, "O email fornecido possui um formato inválido.", "")
		return
	}

	if telefone != "" && !utils.ValidarTelefone(telefone) {
		utils.SendHTMLResponse(w, "Telefone Inválido", utils.ResponseError, "Telefone deve ter entre 10 e 11 dígitos.", "")
		return
	}

	if tipoSanguineo != "" && !utils.ValidarTipoSanguineo(tipoSanguineo) {
		utils.SendHTMLResponse(w, "Tipo Sanguíneo Inválido", utils.ResponseError, "Tipo sanguíneo inválido. Valores aceitos: O-, O+, A-, A+, B-, B+, AB-, AB+", "")
		return
	}

	paciente := utils.Paciente{
		CPF:           cpf,
		Nome:          nome,
		Telefone:      telefone,
		Email:         email,
		TipoSanguineo: tipoSanguineo,
		Convenio:      convenio,
		Observacoes:   observacoes,
	}

	err = utils.AtualizarPaciente(paciente)
	if err != nil {
		utils.SendHTMLResponse(w, "Erro ao Atualizar", utils.ResponseError, fmt.Sprintf("Não foi possível atualizar o paciente. %v", err), "")
		return
	}

	detalhes := map[string]string{
		"CPF":            cpf,
		"Nome":           nome,
		"Telefone":       telefone,
		"Email":          email,
		"Tipo Sanguíneo": tipoSanguineo,
		"Convênio":       convenio,
	}

	utils.SendDetailedResponse(w, "Paciente Atualizado", utils.ResponseSuccess, "Dados do paciente atualizados com sucesso!", detalhes)
}
