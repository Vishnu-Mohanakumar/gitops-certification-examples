package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/runcommand", func(w http.ResponseWriter, r *http.Request) {
		command := r.FormValue("command")

		cmd := exec.Command("sh", "-c", command)
		out, err := cmd.CombinedOutput()

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprintf(w, "%s", out)
	})

	http.ListenAndServe(":8080", nil)
}
