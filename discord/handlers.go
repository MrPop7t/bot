package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CommandsHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	
	if message.Author.ID == session.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "!Untitled")  {
		return
	}

	cmdArgs := strings.Split(message.Content, " ")[1:]

	if len(cmdArgs) == 0 {
		session.ChannelMessageSend(message.ChannelID, errorMessage("Missing comand!", "For a help typed !untitled help"))
		return 
	}

	switch cmdArgs[0] {
	case "ping":
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	case "pong":
		session.ChannelMessageSens(message.ChannelID, "Ping!")
	case "plagueis":
		session.ChannelMessageSend(message.ChannelID, "Did you ever hear the Tragedy of Darth")
	case "help":
		if len(cmdArgs) > 1 {
			helpCommandHandler(session, message, cmdArgsp[1])
		} else {
			helpCommandHandler(session, message, "")
		}
	default:
		session.ChannelMessageSend(message.ChannelID, errorMessage("Invalid command!", "For a list of commands/help, type `!unititled help`"))
	}

}

func errorMessage(title string, message string) string {
	return "X **" + title + "**\n" + message
}