package discord

import "github.com/bwmarrin/discordgo"

func helpCommandHandler(session *discordgo.Session, message *discordgo.MessageCreate, helpSubject string) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	switch helpSubject {
	case "ping":
		title = "Ping help!"
		description = "You get the pong!"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "This is only a test",
				Value: "`!untitled ping`: Pong!",
			},
		}
	case "pong":
		title = "Pong help!"
		description = "You get the ping!"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "This is only a test",
				Value: "`!untitled pong`: Ping!",
			},
		}
	case "hello":
		title = "Hello help!"
		description = "World!"
		fields = []*discordgo.MessageEmbedField{}
	case "plagueis":
		title = "The dark side this is"
		description = "Not a story the jedi would tell"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Darth Plagueis",
				Value: "A bad dood",
			},
			{
				Name:  "Grogu",
				Value: "p cute, kinda nice",
			},
		}
	default:
		title = "General help!"
		description = "Pick a topic below to get help"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "ping",
				Value: "`!untitled help ping` ",
			},
			{
				Name:  "pong",
				Value: "`!untitled help pong`",
			},
			{
				Name:  "hello",
				Value: "`!untitled help hello`",
			},
			{
				Name:  "plagueis",
				Value: "`!untitled help plagueis`",
			},
		}
	}

	session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       15105570,
	})
}