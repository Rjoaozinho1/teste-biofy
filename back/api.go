package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type APIServer struct {
	Port    string
	Storage Storage
}

func NewApiServer(port string, storage Storage) *APIServer {
	return &APIServer{
		Port:    port,
		Storage: storage,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/itens", ValidaJWTHandler(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			s.handleListaritens(w, r)
		case "POST":
			s.handleCriaitens(w, r)
		default:
			http.Error(w, "metodo nao permitido", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/itens/", ValidaJWTHandler(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/itens/"):]
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
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	log.Printf("Servidor iniciado na porta %s...", s.Port)
	if err := http.ListenAndServe(s.Port, mux); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}

func (s *APIServer) handleCriaitens(w http.ResponseWriter, r *http.Request) {
	var item Item
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

func (s *APIServer) handleListaritens(w http.ResponseWriter, _ *http.Request) {
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
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}
	if item.Nome == "" || item.Mensagem == "" {
		http.Error(w, "Campo de nome e mensagem são obrigatorios", http.StatusBadRequest)
		return
	}
	if err := s.Storage.UpdateItem(id, item); err != nil {
		log.Println(err)
		http.Error(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}
	item, err := s.Storage.GetItem(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Erro ao recuperar o item atualizado", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
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
	var req ReqLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}
	token, err := CriaJWTToken(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Credenciais inválidas", http.StatusInternalServerError)
		return
	}
	resp := ResLogin{token}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
