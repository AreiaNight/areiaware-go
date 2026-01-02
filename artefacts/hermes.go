package atenea

import (
    "fmt"
	"time"
	"strings"
    "github.com/bwmarrin/discordgo"
)

var discordSession *discordgo.Session // Crea una nueva sesion de discord

// Inicializa la conexion con discord
func InitDiscord() error {

	// Verificacion basica de notificaciones de discord
	if !GlobalConfig.Discord.Enabled {
		fmt.Println("Discord notifications are disabled")
		return nil
	}
	
	// Ya funciona, no tocar plox
	token := GlobalConfig.Discord.BotToken
	
	 // Verificar que no tenga espacios
    if strings.Contains(token, " ") {
        fmt.Println("[!]  WARNING: Token contiene espacios!")
    }
    
    fmt.Println("[!] Creating Discord session...")


	// Crea una nueva sesion de discord con el token
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Errorf("Error creating Discord session: %v", err)
	}

	// Abre al conexion con discord
	err = session.Open()
	if err != nil {
		return fmt.Errorf("Error opening Discord session: %v", err)
	}

	discordSession = session // Se guarda la sesion en la variable global 

	fmt.Println("[!] Discord bot initialized successfully")
	fmt.Println("[!] Discord bot is now ON")
	return nil

}

// SendSimpleMessage envía un mensaje de texto simple a Discord
func SendSimpleMessage(message string) error {
    
	// Si Discord está deshabilitado, no hacer nada
    if !GlobalConfig.Discord.Enabled {
        return nil
    }

    // Si la sesión no está inicializada, error
    if discordSession == nil {
        return fmt.Errorf("Discord session not initialized")
    }

    // Enviar el mensaje al canal configurado
    _, err := discordSession.ChannelMessageSend(
        GlobalConfig.Discord.ChannelID,  // A qué canal
        message,                          // Qué mensaje
    )
    
    if err != nil {
        return fmt.Errorf("error sending message to Discord: %v", err)
    }

    return nil
}

// SendNotificationEmbed envía una notificación con formato bonito (embed)

func SendNotificationEmbed(title, description string, color int) error {
   
	// Si Discord está deshabilitado, no hacer nada
    if !GlobalConfig.Discord.Enabled {
        return nil
    }

    // Si la sesión no está inicializada, error
    if discordSession == nil {
        return fmt.Errorf("Discord session not initialized")
    }

    // Crear el embed (mensaje bonito)
    embed := &discordgo.MessageEmbed{
        Title:       title,                    // Título del mensaje
        Description: description,               // Contenido
        Color:       color,                     // Color del borde (en decimal)
        Timestamp:   time.Now().Format(time.RFC3339), // Hora actual
    }

    // Enviar el embed
    _, err := discordSession.ChannelMessageSendEmbed(
        GlobalConfig.Discord.ChannelID,
        embed,
    )

    if err != nil {
        return fmt.Errorf("error sending embed to Discord: %v", err)
    }

    return nil
}