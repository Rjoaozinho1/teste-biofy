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

func newApiServer(port string, storage Storage) *APIServer {
	return &APIServer{
		Port:    port,
		Storage: storage,
	}
}

func (s *APIServer) Run() {

	http.HandleFunc("/items", s.handleListarItems)
	http.HandleFunc("/items", s.handleCriaItems)
	http.HandleFunc("/items/{id}", s.handleGetItem)
	http.HandleFunc("/items/{id}", s.handleAtualizaItem)
	http.HandleFunc("/items/{id}", s.handleDeletaItem)

	if err := http.ListenAndServe(s.Port, nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
	log.Printf("Servidor iniciado na porta %s...", s.Port)
}

func (s *APIServer) handleCriaItems(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}

	item.SessionID = uuid.New()
	item.CreatedAt = time.Now()

	if err := s.Storage.CreateItem(item); err != nil {
		http.Error(w, "Erro ao criar o item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (s *APIServer) handleListarItems(w http.ResponseWriter, r *http.Request) {
	items, err := s.Storage.GetAllItems()
	if err != nil {
		http.Error(w, "Erro ao trazer todos os items", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (s *APIServer) handleGetItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	item, err := s.Storage.GetItem(id)
	if err != nil {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (s *APIServer) handleAtualizaItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Corpo Invalido", http.StatusBadRequest)
		return
	}

	if err := s.Storage.UpdateItem(id, item); err != nil {
		http.Error(w, "Erro ao atualizar o item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (s *APIServer) handleDeletaItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := s.Storage.DeleteItem(id); err != nil {
		http.Error(w, "Erro ao excluir o item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
