package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var token, port string

func minit() {
	fmt.Println("in init main.go")
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error initislizing the local environment", err)
	}
	token = os.Getenv("TOKEN")
	if token == "" {
		log.Panic("Set TOKEN in your environment variables")
	}
	port = os.Getenv("PORT")
	if token == "" {
		log.Panic("Set PORT in your environment variables")
	}
}

var (
	update Update
	msg    *Message
	err    error
)

// HandleRequests endpoint that'll receive updates from the telegram api via webhook
// TODO: make goroutines to perform the actions for the deserialized body
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		fmt.Println("Could not read update Body")
	}

	fmt.Println("text:", update.Message.Text, "username:", update.Message.From.Username, "firstname:", update.Message.From.FirstName, "lastname:", update.Message.From.LastName)
	// fmt.Println("text:", update.Message.Text, "username:", update.Message.ForwardFrom.Username, "firstname:", update.Message.ForwardFrom.FirstName, "lastname:", update.Message.ForwardFrom.LastName)
	if update.Message.Entities[0].Type == "bot_command" {
		txt := update.Message.Text
		if txt == "/sayhello" {
			firstname := update.Message.From.FirstName
			msg, err = bot.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Hello, %s.\nHow are you today? Or would you prefer I called you %s", update.Message.From.Username, firstname))
			if err != nil {
				fmt.Println("error in sending reply:", err)
			}

		} else if txt == "/start" {
			msg, err = bot.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Hello, I am %s and I am here to keep you company!", bot.Name))
			if err != nil {
				fmt.Println("error in sending reply:", err)
			}

		} else if txt == "/sendlocation" {
			_, err = bot.SendLocation(update.Message.Chat.ID, 7.418574, 3.88426)
			if err != nil {
				fmt.Println("error in sending reply:", err)
			}
		}
	} else {
		msg, err = bot.SendMessage(update.Message.Chat.ID, "Hello, please send a supported command /playgame not working right now")
		if err != nil {
			fmt.Println("error in sending reply:", err)
		}
	}

}

func main() {
	http.HandleFunc("/", HandleRequests)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("port already taken", err)
	}
	fmt.Println("Running")
}
