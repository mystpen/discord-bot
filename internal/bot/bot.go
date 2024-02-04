package bot

import (
	"discord-bot/internal/usecase"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type BotService struct {
	session *discordgo.Session
	weather *usecase.WeatherUsecase
	game    *usecase.GameUsecase
}

func NewBotService(botToken, weatherApiKey string) *BotService {
	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &BotService{
		session: session,
		weather: usecase.NewWeatherUsecase(weatherApiKey),
		game:    usecase.NewGameUsecase(),
	}
}

// commands implementation
func (b *BotService) messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	content := strings.Split(message.Content, " ")
	switch {
	case content[0] == "!help":
		information := "Command list:\n"
		information += "•	!weather <location>  // shows the current weather of the location\n"
		information += "•	!rolldice  // throws the dice and shows what came up"
		session.ChannelMessageSend(message.ChannelID, information)

	case content[0] == "!weather" && len(content) == 2:
		weatherInfo, err := b.weather.GetWeather(content[1])
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, err.Error())
			return
		}
		session.ChannelMessageSend(message.ChannelID, weatherInfo)

	case content[0] == "!rolldice" && len(content) == 1:
		result := b.game.RollDice()
		session.ChannelMessageSend(message.ChannelID, result)

	default:
		return
	}
}
