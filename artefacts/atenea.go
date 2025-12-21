package atenea

import (
    "crypto/sha256"
    cryptoRand "crypto/rand"
    "fmt"
    mathRand "math/rand"
    "os"
    "strings"
    "syscall"
    "time"
    
    "golang.org/x/term"
)

// Genera el Password aleatorio usando los nombres 
func GeneratePasswd() string{  // ← Mayúscula
	mathRand.Seed(time.Now().UnixNano()) // Generar semilla random basado en el tiempo actual y nano segundos

	var selected[]string // variable tipo slice para guardar los personajes, el var puede actuar como un :=
	usedIndices := make(map[int]bool) // Mapa para asegurarnos que no se repitan personajes, es como una matriz (mapa) 
	
	for len(selected) < 3 {
		idx := mathRand.Intn(len(CastlevaniaCharacters))  // ← Mayúscula

		if !usedIndices[idx] { 
			selected = append(selected, CastlevaniaCharacters[idx])  // ← Mayúscula
			usedIndices[idx] = true 
		}
	}

	numero := mathRand.Intn(100)
	recovery := fmt.Sprintf("%s-%s-%s-%d", selected[0], selected[1], selected[2], numero)
	return recovery
}

// Función para leer el password
func GetPsswd() []byte{  // ← Mayúscula
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

	userInput := strings.TrimSpace(string(passwd))

	// Verificar contra master password o recovery password
	if userInput == MasterPassword || userInput == RecoveryPassword {  // ← Mayúscula
		// Siempre usar masterPassword para el hash (consistencia)
		hash := sha256.Sum256([]byte(MasterPassword)) // Convertimos el password a hash - ← Mayúscula
		fmt.Println("[✓] Password accepted!")
		return hash[:]
	} else {
		fmt.Println("[x] Wrong password")
		os.Exit(1)
	}

	return nil // Nunca llegará aquí
}

// Funcion para generar nombre aleatorio 
func generateRandomName() string {  // ← minúscula (función interna)
    b := make([]byte, 16)
    cryptoRand.Read(b)  // ← Usar cryptoRand
    return fmt.Sprintf("%x.dat", b)
}