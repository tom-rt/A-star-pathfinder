package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	var line string = ""
	var maze []string
	if err != nil {
		log.Fatal(fileName, " File does not exist:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		maze = append(maze, line)
	}
	return maze
}

func getFileName() string {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file name.")
		return ""
	}
	fileName := os.Args[1]
	return fileName
}

func main() {
	var fileName string = getFileName()
	if len(fileName) == 0 {
		return
	}
	var maze []string = readFile(fileName)
	fmt.Println(maze)
	return
}