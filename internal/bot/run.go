package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	disToken := os.Getenv("DISCORD_TOKEN")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	// Bot Service 
	botService := NewBotService(disToken, weatherApiKey)

	botService.session.AddHandler(botService.messageCreate)

	// Open session
	err = botService.session.Open()
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Println("Bot is running")

	// Shutdown
	defer botService.session.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
