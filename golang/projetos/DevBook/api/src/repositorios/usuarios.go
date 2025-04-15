package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) BuscarUsuarios() ([]modelos.Usuario, error) {

	var usuarios []modelos.Usuario

	linhas, erro := repositorio.db.Query("select * from usuarios")
	if erro != nil {
		return usuarios, erro
	}

	defer linhas.Close()

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email); erro != nil {
			return usuarios, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}
