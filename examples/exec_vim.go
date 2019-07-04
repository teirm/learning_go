// Exec vim from go
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    
    cmd := exec.Command("vim")
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
    cmd.Stderr = os.Stderr
    cmd.Run()

    fmt.Printf("All done")
}

