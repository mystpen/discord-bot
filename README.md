# Discord bot

Discord bot written in golang programming language.

## About
- <b>Unique Features:</b>
    - <b>Weather Updates:</b> On command, the bot fetches and displays current weather information for a specified location.
    - <b>Custom Game:</b> A simple text-based game that users can play directly in the chat.
- <b>Commands:</b>
    - ```!help``` all existing commands
    - ```!weather <location>```   current weather of the location
    - ```!rolldice```  mini-game: throws the dice and shows what came up

## Usage

- Clone repo on your work repository:

```
git clone git@github.com:mystpen/discord-bot.git
```
- Go to ```discord-bot``` directory
- The following environment variables were used for the project:
    - DISCORD_TOKEN  (api token for discord bot)
    - WEATHER_API_KEY   (openweathermap api key)

- To run project:
```
go run ./cmd/bot/.
```