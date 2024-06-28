package parser

import (
	diccionario2 "TP2-Analisis-Logs/src/log-analyzer/tdas/BST"
	diccionario "TP2-Analisis-Logs/src/log-analyzer/tdas/hash"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type LogLine struct {
	ip        string
	Timestamp time.Time
	Method    string
	Url       string
}

const layout = "2006-01-02T15:04:05-07:00"

func ParseLogFile(logFile string) ([]LogLine, error) {
	file, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var logLines []LogLine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		logFields := strings.Split(lines, "\t")
		//ip := net.ParseIP(logFields[0])
		timestamp, _ := time.Parse(layout, logFields[1])
		logEntry := LogLine{
			ip:        logFields[0],
			Timestamp: timestamp,
			Method:    logFields[2],
			Url:       logFields[3],
		}
		logLines = append(logLines, logEntry)
	}
	return logLines, nil
}

func ProcessLogFile(logLines []LogLine, ipDict diccionario.Diccionario[string, []time.Time]) {
	//var suspectIps []string
	for i := 0; i < len(logLines); i++ {
		ip := logLines[i].ip
		if !ipDict.Pertenece(ip) {
			ipDict.Guardar(ip, []time.Time{logLines[i].Timestamp})
		} else {
			timestamp := ipDict.Obtener(ip)
			timestamp = append(timestamp, logLines[i].Timestamp)
			ipDict.Guardar(ip, timestamp)
			if len(timestamp) >= 5 {
				if timestamp[len(timestamp)-1].Sub(timestamp[len(timestamp)-5]) < 2*time.Second {
					fmt.Printf("DoS: %s\n", ip)
					//suspectIps = append(suspectIps, ip)
				}
			}
		}
}
