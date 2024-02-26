package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
}

func Run() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ldate|log.Lshortfile)

	// app := &application{
	// 	errorLog: errorLog,
	// 	infoLog: infoLog,
	// }

	//logging to a file/////////////////////////////////
	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// infoLog = log.New(f, "INFO\t", log.Ldate|log.Ltime)

	// srv := &http.Server{
	// 	Addr: *addr,
	// 	ErrorLog: errorLog,
	// 	Handler: mux,
	// }

	//application///////////////////////////////


	err := godotenv.Load()
	if err != nil {
		errorLog.Fatal("Error loading .env file")
	}
	disToken := os.Getenv("DISCORD_TOKEN")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	// Bot Service
	botService := NewBotService(disToken, weatherApiKey)

	botService.session.AddHandler(botService.messageCreate)

	// Open session
	err = botService.session.Open()
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Println("Bot is running")

	// Shutdown
	defer botService.session.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
