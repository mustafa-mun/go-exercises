package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

// // printPrompt displays the repl prompt at the start of each loop
// func printPrompt() {
//     fmt.Print(cliName, "> ")
// }
 
// // printUnkown informs the user about invalid commands
// func printUnknown(text string) {
//     fmt.Println(text, ": command not found")
// }
 
// // displayHelp informs the user about our hardcoded functions
// func displayHelp() {
//     fmt.Printf(
//         "Welcome to %v! These are the available commands: \n",
//         cliName,
//     )
//     fmt.Println("-h / -help  - Show available commands")
//     fmt.Println("clear   - Clear the terminal screen")
// }
 
// // clearScreen clears the terminal screen
// func clearScreen() {
//     cmd := exec.Command("clear")
//     cmd.Stdout = os.Stdout
//     cmd.Run()
// }
 
// // handleInvalidCmd attempts to recover from a bad command
// func handleInvalidCmd(text string) {
//     defer printUnknown(text)
// }
 
// // handleCmd parses the given commands
// func handleCmd(text string) {
//     handleInvalidCmd(text)
// }
 
// // cleanInput preprocesses input to the db repl
// func cleanInput(text string) string {
//     output := strings.TrimSpace(text)
//     output = strings.ToLower(output)
//     return output
// }
 
