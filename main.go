package main

import (
	"fmt"
	// "goupie-tracker/handler"
	"goupie-tracker/helper"
	"goupie-tracker/routes"
	"net/http"
	"os"
	"strconv"
	"strings"
)


func main() {
	PORT := envPortOr("3000")
	args := os.Args[1:]
	if len(args) > 0 {
		tab := strings.Split(args[0], "=")
		if len(tab) > 1 && helper.IsInt(tab[1]) && tab[1] != "" {
			t, _ := strconv.Atoi(tab[1])
			if t >= 1024 && t <= 65535 {
				PORT = ":" + tab[1]
			}
		}
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.Route()
	fmt.Println("Listening in http://localhost" + PORT)

	http.ListenAndServe(PORT, nil)
}
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
	  return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
  }