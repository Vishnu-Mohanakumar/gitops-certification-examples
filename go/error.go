package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
		db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database")
		if err != nil {
			http.Error(w, "Error in database connection", 500)
			return
		}
		defer db.Close()

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Error in query execution", 500)
			return
		}
		defer rows.Close()
	})
	http.ListenAndServe(":8080", nil)
}
