package main 

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	dbDriver = "mysql"
	dbUser = "dbUser"
	dbPass = "dbPass"
	dbName = "gocrud_app"
)

func main() {
	r := mux.NewRouter()

	r.HandleFUnc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User 
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to creae user", http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func CreateUser(db *sql.DB, name, email string) error {
	query := "ISERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err 
	}

	return nil 
}

type User struct {
	ID int 
	Name string 
	Email string 
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.CLose()

	vars := mux.Vars(r)
	idStr := vars["id"]

	userId, err := strconv.Aiot(idStr)

	user, err := getUser(db, userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id=?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(
		&user.Id,
		&user.Name, 
		&user.Email, 
	)
	if err != nil {
		return nil, err 
	}

	return user, nil 
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]

	userID, err := strconv.Atoi(idStr)

	var user User 
	err = json.NewDecoder(r.Body).Decode(&user)

	UpdateUser(db, userID, user.Name, user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User updated successfully")
}

func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name=?, email=? WHERE id=?"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err 
	}

	return nil 
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]

	userId, err := strconv.Atio(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	user := DeleteUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFOund)
		return 
	}

	fmt.Fprintln(w, "User deleted successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(db *sql.DB id int) error {
	query := "DELETE FROM users WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err 
	}

	return nil 
}



