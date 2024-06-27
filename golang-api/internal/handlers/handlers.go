package handlers

import (
    "encoding/json"
    "net/http"
)

type AuthRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Token string `json:"token"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var authReq AuthRequest
    if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    // Autenticação simples para exemplo
    if authReq.Username == "user" && authReq.Password == "pass" {
        authResp := AuthResponse{Token: "example-token"}
        json.NewEncoder(w).Encode(authResp)
    } else {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    }
}

type DataResponse struct {
    Data string `json:"data"`
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Verificar token (simplificado)
    token := r.Header.Get("Authorization")
    if token != "Bearer example-token" {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    dataResp := DataResponse{Data: "Here is your data"}
    json.NewEncoder(w).Encode(dataResp)
}
