package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json: "id"`
	Nome  string `json: "nome"`
	Email string `json: "email"`
}

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	//Fecha a conexão com o banco no final da execução da função
	defer db.Close()

	//PREPARE STATMENT - EVITA ATAQUE DE SQL INJECTION
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	//Fecha a conexão com do statement no final da execução da função
	defer statement.Close()

	//Executa o statement passando os campos para o insert
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}
	//Recupera o id do usuario inserido pelo statement
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao recuperar id do usuário inserido!"))
		return
	}

	//Cria o STATUS CODES e a mensagem de retorno para retornar
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

}

// BuscarUsuarios retorna todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuarios!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON!"))
		return
	}

}

// BuscarUsuario retorna um usuário específico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro para converter o parametro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario!"))
		return
	}

	var usuario usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuario!"))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	if usuario.ID == 0 {
		w.Write([]byte(fmt.Sprintf("Usuário não encontrado com o id %d!", ID)))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON!"))
		return
	}

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	//Ler o parametro ID
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro para converter o parametro para inteiro!"))
		return
	}
	//Lê o corpo da requisição
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	fmt.Println(corpoDaRequisicao)
	if corpoDaRequisicao == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Corpo da requisicao vazio. Tente novamente!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	if usuario.Nome != "" && usuario.Email != "" {
		statement, erro := db.Prepare("update usuarios set nome = ?, email= ? where id = ?")
		if erro != nil {
			w.Write([]byte("Erro ao criar statement!"))
			return
		}
		//Fecha a conexão com do statement no final da execução da função
		defer statement.Close()

		//Executa o statement passando os campos para o insert
		if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
			w.Write([]byte("Erro ao atualizar o usuario!"))
			return
		}
	} else if usuario.Nome != "" && usuario.Email == "" {
		statement, erro := db.Prepare("update usuarios set nome = ? where id = ?")
		if erro != nil {
			w.Write([]byte("Erro ao criar statement!"))
			return
		}
		//Fecha a conexão com do statement no final da execução da função
		defer statement.Close()

		//Executa o statement passando os campos para o insert
		if _, erro := statement.Exec(usuario.Nome, ID); erro != nil {
			w.Write([]byte("Erro ao atualizar o usuario!"))
			return
		}
	} else {
		statement, erro := db.Prepare("update usuarios set email = ? where id = ?")
		if erro != nil {
			w.Write([]byte("Erro ao criar statement!"))
			return
		}
		//Fecha a conexão com do statement no final da execução da função
		defer statement.Close()

		//Executa o statement passando os campos para o insert
		if _, erro := statement.Exec(usuario.Email, ID); erro != nil {
			w.Write([]byte("Erro ao atualizar o usuario!"))
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Usuário de id %d atualizado com sucesso!", ID)))

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro para converter o parametro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	if _, erro = statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuario!"))
		return
	}

	defer statement.Close()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Usuario deletado com sucesso"))

}
