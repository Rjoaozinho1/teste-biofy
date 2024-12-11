package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Rjoaozinho1/teste-biofy/auth"
	"github.com/Rjoaozinho1/teste-biofy/repo"
	"github.com/Rjoaozinho1/teste-biofy/storage"
	"github.com/google/uuid"
)

type APIServer struct {
	Port    string
	Storage storage.Storage
}

func NewApiServer(port string, storage storage.Storage) *APIServer {
	return &APIServer{
		Port:    port,
		Storage: storage,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/mensagem", auth.ValidaJWTHandler(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			s.handleListarItens(w, r)
		case "POST":
			s.handleCriaItens(w, r)
		default:
			http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/mensagem/", auth.ValidaJWTHandler(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/mensagem/"):]
		if id == "" {
			http.Error(w, "ID necessario", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "GET":
			s.handleGetItem(w, r, id)
		case "PUT":
			s.handleAtualizaItem(w, r, id)
		case "DELETE":
			s.handleDeletaItem(w, r, id)
		default:
			http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
		}
		s.handleLogin(w, r)
	})
	mux.HandleFunc("/validate-token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
		}
		s.handleValidaToken(w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("../app/public")))

	log.Printf("Servidor iniciado na porta %s...", s.Port)
	if err := http.ListenAndServe(s.Port, mux); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}

func (s *APIServer) handleCriaItens(w http.ResponseWriter, r *http.Request) {
	var item repo.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}
	if item.Nome == "" || item.Mensagem <= "" {
		http.Error(w, "Campo de nome e mensagem são obrigatorios", http.StatusBadRequest)
		return
	}
	item.SessionID = uuid.New()
	item.CreatedAt = time.Now()
	if err := s.Storage.CreateItem(item); err != nil {
		log.Println(err)
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (s *APIServer) handleListarItens(w http.ResponseWriter, _ *http.Request) {
	itens, err := s.Storage.GetAllItens()
	if err != nil {
		http.Error(w, "Erro ao trazer todos os itens", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itens)
}

func (s *APIServer) handleGetItem(w http.ResponseWriter, _ *http.Request, id string) {
	item, err := s.Storage.GetItem(id)
	if err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (s *APIServer) handleAtualizaItem(w http.ResponseWriter, r *http.Request, id string) {
	var item repo.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}
	if item.Nome == "" || item.Mensagem == "" {
		http.Error(w, "Campo de nome e mensagem são obrigatorios", http.StatusBadRequest)
		return
	}
	item, err := s.Storage.GetItem(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Erro ao achar o item no banco de dados", http.StatusInternalServerError)
		return
	}
	if err := s.Storage.UpdateItem(id, item); err != nil {
		log.Println(err)
		http.Error(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}
	item2, err := s.Storage.GetItem(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Erro ao recuperar o item atualizado", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item2)
}

func (s *APIServer) handleDeletaItem(w http.ResponseWriter, _ *http.Request, id string) {
	if err := s.Storage.DeleteItem(id); err != nil {
		log.Println(err)
		http.Error(w, "Erro ao excluir o item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req repo.ReqLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}
	token, err := auth.CriaJWTToken(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Credenciais inválidas", http.StatusInternalServerError)
		return
	}
	resp := repo.ResLogin{
		Token: token,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *APIServer) handleValidaToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	// log.Println(authHeader)
	if authHeader == "" {
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return
	}
	// Remove o prefixo "Bearer " do cabeçalho, se presente
	token := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}
	// Valida o token
	_, err := auth.ValidaJWTToken(token)
	if err != nil {
		log.Println("Erro ao validar o token:", err)
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}
	// Se o token for válido, retorna um status 200 com uma mensagem de sucesso
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Token válido"})
}
