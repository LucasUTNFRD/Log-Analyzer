package main

import (
	"TP2-Analisis-Logs/src/log-analyzer/parser"
	diccionario "TP2-Analisis-Logs/src/log-analyzer/tdas/hash"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if err := ejecutarInterfaz(command); err != nil {
			fmt.Fprintf(os.Stderr, "Error en comando : %v\n", err)
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

// command to execute

//agregar_archivo <nombre_archivo>: procesa de forma completa un archivo de log.
//ver_visitantes <desde> <hasta>: muestra todas las IPs que solicitaron algún recurso en el servidor, dentro del rango de IPs determinado.
//ver_mas_visitados <n>: muestra los n recursos más solicitados.

func ejecutarInterfaz(command string) error {
	cmd := strings.Fields(command)
	log.Println(cmd)
	if len(cmd) == 0 {
		return fmt.Errorf("comando vacio")
	}
	switch cmd[0] {
	case "agregar_archivo":
		if len(cmd) != 2 {
			return fmt.Errorf("numero invalido de argumentos")
		}
		nombreArchivo := cmd[1]
		return agregarArchivo(nombreArchivo)
	case "ver_visitantes":
		if len(cmd) != 3 {
			return fmt.Errorf("numero invalido de argumentos")
		}
		desde, hasta := cmd[1], cmd[2]
		return verVisitantes(desde, hasta)
	case "ver_mas_visitados":
		if len(cmd) != 2 {
			return fmt.Errorf("numero invalido de argumentos")
		}
		n := cmd[1]
		return verMasVisitado(n)
	default:
		return fmt.Errorf(cmd[0])
	}
}

func agregarArchivo(nombreArchivo string) error {
	logLines, err := parser.ParseLogFile(nombreArchivo)
	if err != nil {
		return err
	}
	ipDict := diccionario.CrearHash[string, []time.Time]()
	parser.ProcessLogFile(logLines, ipDict)
	fmt.Printf("OK\n")
	return nil
}

func verMasVisitado(n string) error {
	return nil
}

func verVisitantes(desde string, hasta string) error {
	return nil
}
