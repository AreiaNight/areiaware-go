package modules

// Common shared constants and types for modules

const (
    TARGET_DIR    = "./archivos_privados"
    MAPPING_FILE  = ".secret_mapping.enc"
    HIDDEN_FOLDER = ".hidden_data"
)

// FileMapping guarda la relación entre archivo original y oculto
type FileMapping struct {
    OriginalPath string `json:"original"`
    HiddenPath   string `json:"hidden"`
}

// Lista de personajes (usada para generar contraseñas de recuperación)
var CastlevaniaCharacters = []string{
    "Carmilla", "Lenore", "Morana", "Tera", "Maria", "Annette", "Drolta", "Striga",
}

// Master/Recovery passwords pueden residir en main pero se exponen aquí si se desea compartir
var MasterPassword = "MadameYuDoberman"
var RecoveryPassword = ""
