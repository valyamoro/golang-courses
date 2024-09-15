package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

import redis "github.com/go-redis/redis/v8"
import jwt "github.com/golang-jwt/jwt/v4"

var redisClient *redis.Client
var jwtSecret = []byte("supersecretkey")

type User struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

type TokenClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

var users = map[int]User{
	1: {ID: 1, Name: "Maxim", Items: []string{"Sword", "Shield"}},
	2: {ID: 2, Name: "Alice", Items: []string{"Bow", "Arrow"}},
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	http.HandleFunc("/login", login)
	http.HandleFunc("/users/", authMiddleware(getUsers))
	http.HandleFunc("/items", authMiddleware(getItems))

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		claims := &TokenClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	userID := 1
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaim{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour())),
		},
	})

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	user, ok := users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	cacheKey := fmt.Sprintf("user:%d", userID)
	if cachedUser, err := redisClient.Get(r.Context(), cacheKey).Result(); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedUser))
		return
	}

	userJSON, _ := json.Marshal(user)
	redisClient.Set(r.Context(), cacheKey, userJSON, 5*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	user, ok := users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var items []string
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		if i > 0 {
			fmt.Printf("Retrying...(%d/%d)\n", i, maxRetries)
		}

		if i == 2 {
			items = user.Items
			break
		}

		time.Sleep(1 * time.Second)
	}

	if items == nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	batchSize := 2
	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}

		w.Write([]byte(fmt.Sprintf("Processing batch: %v\n", items[i:end])))
	}
}

func getIDFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]

	return strconv.Atoi(idStr)
}
