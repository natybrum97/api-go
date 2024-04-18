package main

//Declaração do pacote principal.

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Importações de pacotes necessários:

//"encoding/json": Para codificar e decodificar JSON.
//"fmt": Para formatar a saída.
//"log": Para fazer log de mensagens.
//"net/http": Para lidar com solicitações HTTP.
//"github.com/gorilla/mux": Um roteador HTTP para Go.

// Definindo a estrutura do usuário
// Definição da estrutura User para representar um usuário com campos ID, Username e Email.
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Declaração da variável users, que armazenará todos os usuários em memória.
var users []User

// Função main(), que é a função de inicialização do programa.
func main() {
	// Inicializando o roteador do Gorilla mux
	router := mux.NewRouter()

	//Definição de várias rotas para manipular solicitações HTTP para diferentes endpoints da API.
	// Rota para obter todos os usuários
	router.HandleFunc("/users", GetUsers).Methods("GET")

	// Rota para obter um usuário específico por ID
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	// Rota para criar um novo usuário
	router.HandleFunc("/users", CreateUser).Methods("POST")

	// Rota para atualizar um usuário existente
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")

	// Rota para excluir um usuário
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Inicia o servidor na porta 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Função para obter todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//Recebe dois parâmetros: w (http.ResponseWriter) para escrever a resposta HTTP e
	//r (http.Request) para acessar os detalhes da solicitação HTTP.

	w.Header().Set("Content-Type", "application/json")
	//Define o cabeçalho da resposta HTTP para indicar que o conteúdo é JSON.
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(users)
	//Codifica o slice de usuários (users) em formato JSON e escreve na resposta HTTP
	// usando o http.ResponseWriter fornecido.
}

// Função para obter um usuário por ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Extrai as variáveis(parametros) de rota da solicitação HTTP usando o roteador mux.
	for _, item := range users {
		//Itera sobre o slice de usuários e verifica se há algum usuário com o ID
		//correspondente ao ID fornecido na solicitação.
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusOK) // 200 OK
			//Se um usuário com o ID correspondente for encontrado, ele é codificado em JSON
			// e escrito na resposta HTTP.
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
	//Se nenhum usuário for encontrado, um objeto vazio User{} é codificado em JSON
	//e retornado como resposta.
}

// Função para criar um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//Declaração de uma variável do tipo User para armazenar o novo usuário
	_ = json.NewDecoder(r.Body).Decode(&user)
	//Decodifica o corpo da solicitação HTTP em JSON e preenche a estrutura
	// User com os dados fornecidos
	user.ID = strconv.Itoa(len(users) + 1)

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	//O novo usuário é adicionado ao slice de usuários e, em seguida, é
	//codificado em JSON e escrito na resposta HTTP.
}

// Função para atualizar um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			//estamos efetivamente removendo o elemento de users localizado no índice index
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			w.WriteHeader(http.StatusOK) // 200 OK
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

// Função para excluir um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			for i := range users {
				users[i].ID = strconv.Itoa(i + 1)
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(users)
}
