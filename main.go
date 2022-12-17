package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	bye             = "[Info] Bye!"
	notepadEmpty    = "[Info] Notepad is empty"
	notepadCleared  = "[OK] All notes were successfully deleted"
	noteCreated     = "[OK] The note was successfully created"
	noteUpdated     = "[OK] The note at position %d was successfully updated\n"
	noteDeleted     = "[OK] The note at position %d was successfully deleted\n"
	unknownCommand  = "[Error] Unknown command"
	notepadFull     = "[Error] Notepad is full"
	missingNote     = "[Error] Missing note argument"
	missingPosition = "[Error] Missing position argument"
	invalidPosition = "[Error] Invalid position: %s\n"
	nothingToUpdate = "[Error] There is nothing to update"
	nothingToDelete = "[Error] There is nothing to delete"
	outOfBoundaries = "[Error] Position %d is out of the boundary [1, %d]\n"
)

func main() {
	var (
		maximumNumberNotes int
		notepad            []string
		note               string
	)

	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&maximumNumberNotes)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command and data: ")
		scan := scanner.Scan()
		if scan {
			userInput := scanner.Text()
			userInput = strings.Trim(userInput, " ")
			commands := strings.SplitN(userInput, " ", 2)
			switch commands[0] {
			case "exit":
				fmt.Println(bye)
				os.Exit(1)
			case "create":
				if len(commands) < 2 {
					fmt.Println(missingNote)
					continue
				}
				note = commands[1]
				if len(notepad) < maximumNumberNotes {
					notepad = append(notepad, note)
					fmt.Println(noteCreated)
				} else {
					fmt.Println(notepadFull)
				}
			case "list":
				if len(notepad) == 0 {
					fmt.Println(notepadEmpty)
				} else {
					for i, note := range notepad {
						fmt.Printf("[Info] %d: %s\n", i+1, note)
					}
				}
			case "clear":
				notepad = nil
				fmt.Println(notepadCleared)
			case "update":
				if len(commands) < 2 {
					fmt.Println(missingPosition)
				} else {
					arguments := strings.SplitN(commands[1], " ", 2)
					if len(arguments) < 2 {
						fmt.Println(missingNote)
					} else {
						val, err := strconv.Atoi(arguments[0])
						if err != nil {
							fmt.Printf(invalidPosition, arguments[0])
						} else if val < 1 || maximumNumberNotes < val {
							fmt.Printf(outOfBoundaries, val, maximumNumberNotes)
						} else if len(notepad) < val {
							fmt.Println(nothingToUpdate)
						} else {
							notepad[val-1] = arguments[1]
							fmt.Printf(noteUpdated, val)
						}
					}
				}
			case "delete":
				if len(commands) < 2 {
					fmt.Println(missingPosition)
				} else {
					val, err := strconv.Atoi(commands[1])
					if err != nil {
						fmt.Printf(invalidPosition, commands[1])
					} else if val < 1 || maximumNumberNotes < val {
						fmt.Printf(outOfBoundaries, val, maximumNumberNotes)
					} else if len(notepad) < val {
						fmt.Println(nothingToDelete)
					} else {
						notepad = append(notepad[:val-1], notepad[val:]...)
						fmt.Printf(noteDeleted, val)
					}
				}
			default:
				fmt.Println(unknownCommand)
			}
		}
	}
}
