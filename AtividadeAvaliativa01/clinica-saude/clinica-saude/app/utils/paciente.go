package utils

import (
	"database/sql"
	"errors"
	"fmt"
)

type Paciente struct {
	ID             int
	Nome           string
	CPF            string
	DataNascimento string
	Telefone       string
	Email          string
	TipoSanguineo  string
	Convenio       string
	Observacoes    string
	CriadoEm       string
}

func CriarPaciente(p Paciente) error {
	query := `
		INSERT INTO pacientes (nome, cpf, data_nascimento, telefone, email, tipo_sanguineo, convenio, observacoes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := DB.Exec(query,
		p.Nome,
		p.CPF,
		p.DataNascimento,
		p.Telefone,
		p.Email,
		p.TipoSanguineo,
		p.Convenio,
		p.Observacoes,
	)
	if err != nil {
		return fmt.Errorf("erro ao cadastrar paciente: %v", err)
	}
	return nil
}

func BuscarPacientePorCPF(cpf string) (*Paciente, error) {
	query := `SELECT id, nome, cpf, data_nascimento, telefone, email, tipo_sanguineo, convenio, observacoes, criado_em
	          FROM pacientes WHERE cpf = $1`

	row := DB.QueryRow(query, cpf)

	var p Paciente
	var dataNasc, criadoEm []byte

	err := row.Scan(
		&p.ID, &p.Nome, &p.CPF, &dataNasc,
		&p.Telefone, &p.Email, &p.TipoSanguineo,
		&p.Convenio, &p.Observacoes, &criadoEm,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("paciente com CPF %s não encontrado", cpf)
		}
		return nil, fmt.Errorf("erro ao buscar paciente: %v", err)
	}

	p.DataNascimento = string(dataNasc)
	p.CriadoEm = string(criadoEm)
	return &p, nil
}

func AtualizarPaciente(p Paciente) error {
	query := `
		UPDATE pacientes
		SET nome = $1, telefone = $2, email = $3, tipo_sanguineo = $4, convenio = $5, observacoes = $6
		WHERE cpf = $7`

	result, err := DB.Exec(query,
		p.Nome, p.Telefone, p.Email,
		p.TipoSanguineo, p.Convenio, p.Observacoes, p.CPF,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar paciente: %v", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("paciente com CPF %s não encontrado", p.CPF)
	}
	return nil
}

func ExcluirPaciente(cpf string) error {
	query := `DELETE FROM pacientes WHERE cpf = $1`
	result, err := DB.Exec(query, cpf)
	if err != nil {
		return fmt.Errorf("erro ao excluir paciente: %v", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("paciente com CPF %s não encontrado", cpf)
	}
	return nil
}
