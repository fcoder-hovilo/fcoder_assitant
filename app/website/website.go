// website.go

package website

import (
	"fcoder_assitant/config"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

// StartWebServer starts the web server
func StartWebServer(dg *discordgo.Session, config config.Config) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Welcome to the fcoder_assitant's website!</h1>")
	})

	fmt.Println("Starting web server on :5000...")
	if err := http.ListenAndServe(":5000", router); err != nil {
		fmt.Println("Error starting web server:", err)
	}
}
