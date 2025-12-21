package atenea

// Constantes
const (
    VERSION = "2.0"
    TARGET_DIR    = "./archivos_privados"  // Carpeta objetivo
    MAPPING_FILE  = ".secret_mapping.enc"  // Archivo con el mapeo
    HIDDEN_FOLDER = ".hidden_data"         // Carpeta donde ocultar
)

// Variables globales, para la contraseña 
var (
    MasterPassword   = "MadameYuDoberman"  // Master password - ← Mayúscula
    RecoveryPassword = ""                   // Se genera dinámicamente - ← Mayúscula
)

// FileMapping guarda la relación entre archivo original y oculto, los mapping son una especie de matrices en dónde se genera un "mapa" y guarda cosas en cada casilla
type FileMapping struct {
	// En originalPath se guardan los paths de los files originales para después re-ingresarlos
    OriginalPath string `json:"original"` //Se usan json porque es para guardar datos en un formato entendible y legible
	// El hiidingPath es lo mismo, se crea una carpeta en dónde se guardan todos los files y se guarda el path en este map
    HiddenPath   string `json:"hidden"`
}

// Lista de personajes 
var CastlevaniaCharacters = []string{
	"Carmilla", "Lenore", "Morana", "Tera", "Maria", "Annette", "Drolta", "Striga",
}

func SetRecoveryPassword(pass string) { // Exporta el recovery password
    RecoveryPassword = pass
}