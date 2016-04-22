package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
)

func main() {
    fmt.Println("Building executable server file for Windows...")
    
    cmd := exec.Command("go", "build", "-v", "main.go")
    
    if err := cmd.Run(); err != nil { 
        log.Fatal("Building executable server file for Windows - failed.") 
    } 
    fmt.Println("Building executable server file for Windows - finished.")  
    
    fmt.Println("Building executable server file for Linux...")
    
    cmd = exec.Command("go", "build", "-v", "main.go")
    cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=linux"))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=amd64"))
    
    if err := cmd.Run(); err != nil { 
        log.Fatal("Building executable server file for Linux - failed.") 
    } 
    fmt.Println("Building executable server file for Linux - finished.")
    
    
    fmt.Println("Building client js & css for production...")
    exec.Command("gulp", "--gulpfile", "client/gulpfile.js", "js:bundle").Output()
    exec.Command("gulp", "--gulpfile", "client/gulpfile.js", "css:bundle").Output()
    fmt.Println("Building client js & css for production - finished.")
}