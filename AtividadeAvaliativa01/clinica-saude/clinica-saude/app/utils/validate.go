package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func ValidarCPF(cpf string) bool {
	cpf = strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)

	if len(cpf) != 11 {
		return false
	}

	firstDigit := cpf[0]
	allSame := true
	for i := 1; i < 11; i++ {
		if cpf[i] != firstDigit {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	return true
}

func ValidarEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func ValidarTelefone(telefone string) bool {
	if telefone == "" {
		return true
	}

	cleaned := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, telefone)

	return len(cleaned) == 10 || len(cleaned) == 11
}

func ValidarDataNascimento(dataNascimento string) error {
	parsedDate, err := time.Parse("2006-01-02", dataNascimento)
	if err != nil {
		return fmt.Errorf("formato de data inválido. Use: YYYY-MM-DD")
	}

	if parsedDate.After(time.Now()) {
		return fmt.Errorf("data de nascimento não pode ser no futuro")
	}

	return nil
}

func ValidarTipoSanguineo(tipo string) bool {
	if tipo == "" {
		return true
	}

	tiposValidos := map[string]bool{
		"O-":  true,
		"O+":  true,
		"A-":  true,
		"A+":  true,
		"B-":  true,
		"B+":  true,
		"AB-": true,
		"AB+": true,
	}

	return tiposValidos[tipo]
}

func ValidarPacienteCompleto(nome, cpf, dataNascimento, email, telefone, tipoSanguineo string) error {
	nome = strings.TrimSpace(nome)
	if nome == "" {
		return fmt.Errorf("nome é obrigatório")
	}
	if len(nome) < 3 {
		return fmt.Errorf("nome deve ter pelo menos 3 caracteres")
	}
	if len(nome) > 150 {
		return fmt.Errorf("nome não pode exceder 150 caracteres")
	}

	cpf = strings.TrimSpace(cpf)
	if cpf == "" {
		return fmt.Errorf("CPF é obrigatório")
	}
	if !ValidarCPF(cpf) {
		return fmt.Errorf("CPF inválido")
	}

	if dataNascimento == "" {
		return fmt.Errorf("data de nascimento é obrigatória")
	}
	if err := ValidarDataNascimento(dataNascimento); err != nil {
		return err
	}

	email = strings.TrimSpace(email)
	if email != "" && !ValidarEmail(email) {
		return fmt.Errorf("email inválido")
	}

	telefone = strings.TrimSpace(telefone)
	if telefone != "" && !ValidarTelefone(telefone) {
		return fmt.Errorf("telefone deve ter entre 10 e 11 dígitos")
	}

	if tipoSanguineo != "" && !ValidarTipoSanguineo(tipoSanguineo) {
		return fmt.Errorf("tipo sanguíneo inválido. Valores aceitos: O-, O+, A-, A+, B-, B+, AB-, AB+")
	}

	return nil
}
