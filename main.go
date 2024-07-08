package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

type Response struct {
    Message string `json:"message"`
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func firstAPI(w http.ResponseWriter, r *http.Request) {
    log.Printf("First API hit at %s from %s", time.Now().Format(time.RFC3339), r.RemoteAddr)
    
    response := Response{Message: "Hello from the first API!"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func secondAPI(w http.ResponseWriter, r *http.Request) {
    log.Printf("Second API hit at %s from %s", time.Now().Format(time.RFC3339), r.RemoteAddr)

    user := User{
        ID:    1,
        Name:  "John Doe",
        Email: "john.doe@example.com",
    }
    jsonResponse, err := json.Marshal(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func main() {
    log.Println("Starting server on :5000")
    
    http.HandleFunc("/api/first", firstAPI)
    http.HandleFunc("/api/second", secondAPI)
    log.Fatal(http.ListenAndServe(":5000", nil))
}
