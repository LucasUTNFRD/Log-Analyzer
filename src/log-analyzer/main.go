package main

import (
	"TP2-Analisis-Logs/src/log-analyzer/LogProccessing"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	CommandLoop()
}

func CommandLoop() {
	lp := LogProccessing.NewLogProcessor()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "agregar_archivo":
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "Error en comando agregar_archivo\n")
				return
			}
			lp.ProcessLogFile(parts[1])

		case "ver_visitantes":
			if len(parts) != 3 {
				fmt.Fprintf(os.Stderr, "Error en comando ver_visitantes\n")
				return
			}
			lp.ListVisitors(parts[1], parts[2])

		case "ver_mas_visitados":
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "Error en comando ver_mas_visitados\n")
				return
			}
			lp.ListMostVisited(parts[1]) // Placeholder

		default:
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", command)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando\n")
		return
	}
}
