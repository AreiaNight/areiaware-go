package atenea

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// Configuracion para el json
type Config struct {
    Discord   DiscordConfig   `json:"discord"`    // Busca la configuracion de discord
    Areiaware AreiawareConfig `json:"areiaware"`  // Busca la configuracion de areiaware
    // Estructura > Campo TipoDeDato
}

type DiscordConfig struct {
    BotToken  string `json:"bot_token"`  // Token del bot
    ChannelID string `json:"channel_id"` // ID del canal
    Enabled   bool   `json:"enabled"`    // Habilitado o no
}

type AreiawareConfig struct {
    MasterPassword string `json:"master_password"`  // Contraseña maestra
    TargetDirectory string `json:"target_directory"` // Directorio objetivo
}

type CommandsConfig struct {
    Name     string `yaml:"nombre"`
    Handler  string `yaml:"handler"`
    Template string `yaml:"template"`
}

// Constantes
const (
    VERSION      = "2.5"
    MAPPING_FILE = ".secret_mapping.enc" // Archivo con el mapeo
    HIDDEN_FOLDER = ".hidden_data"       // Carpeta donde ocultar
)

var (
    MasterPassword   = "" // Se carga desde config.json
    RecoveryPassword = "" // Se genera dinámicamente
    TARGET_DIR       = "" // Se carga desde config.json
    GlobalConfig     Config // Variable global para almacenar la configuración
)

// FileMapping guarda la relación entre archivo original y oculto
type FileMapping struct {
    OriginalPath string `json:"original"` // Se usan json porque es para guardar datos en un formato entendible y legible
    HiddenPath   string `json:"hidden"`
}

// Lista de personajes
var CastlevaniaCharacters = []string{
    "Carmilla", "Lenore", "Morana", "Tera", "Maria", "Annette", "Drolta", "Striga",
}

func SetRecoveryPassword(pass string) { // Exporta el recovery password
    RecoveryPassword = pass
}

// LoadConfig para cargar en json
func LoadConfig() error {
    // Leer el archivo
    file, err := os.Open("config.json")
    if err != nil {
        return fmt.Errorf("error opening config file: %v", err)
    }
    defer file.Close()

    // Leer todo el contenido
    data, err := ioutil.ReadAll(file)
    if err != nil {
        return fmt.Errorf("error reading config file: %v", err)
    }

    // Pasar el json, usamos unmarshal para transformar el json en una estructura para go
    var config Config
    err = json.Unmarshal(data, &config)
    if err != nil {
        return fmt.Errorf("error parsing config: %v", err)
    }

    // Asignar a la variable global
    GlobalConfig = config

    // Validar Discord configuraciones
    if config.Discord.BotToken == "" || config.Discord.BotToken == "YOUR_BOT_TOKEN_HERE" {
        return fmt.Errorf("Discord bot token is not configured in the json config file")
    }


    if config.Discord.ChannelID == "" || config.Discord.ChannelID == "YOUR_CHANNEL_ID_HERE" {
        return fmt.Errorf("Discord channel ID is not configured in the json config file")
    }

    // Validamos las configuraciones de Areiaware
    if config.Areiaware.MasterPassword == "" || config.Areiaware.MasterPassword == "YOUR_MASTER_PASSWORD_HERE" {
        return fmt.Errorf("Master password is not configured in the json config file")
    }

    if config.Areiaware.TargetDirectory == "" || config.Areiaware.TargetDirectory == "YOUR_TARGET_DIRECTORY_HERE" {
        return fmt.Errorf("Target directory is not configured in the json config file")
    }

    // Cargamos las variables
    MasterPassword = config.Areiaware.MasterPassword
    TARGET_DIR = config.Areiaware.TargetDirectory

    fmt.Println("[!] Configuration loaded successfully")
    if config.Discord.Enabled {
        fmt.Println("Discord notifications: ENABLED")
    } else {
        fmt.Println("Discord notifications: DISABLED")
    }
    fmt.Println()

    return nil
}