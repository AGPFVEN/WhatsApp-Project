package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/get-message", messageHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mi Página Web</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
    </style>
    <script>
        window.addEventListener('load', function() {
            fetch('/get-message')
                .then(response => response.json())
                .then(data => {
                    document.getElementById('message').innerText = data.message;
                })
                .catch(error => console.error('Error al cargar el mensaje:', error));
        });
    </script>
</head>
<body>
    <h1>Bienvenido a mi página web</h1>
    <p id="message">Cargando mensaje...</p>
</body>
</html>`)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "¡Este es un mensaje desde el servidor!"}
	json.NewEncoder(w).Encode(response)
}
