package LogProccessing

import (
	diccionario2 "TP2-Analisis-Logs/src/log-analyzer/tdas/BST"
	diccionario "TP2-Analisis-Logs/src/log-analyzer/tdas/hash"
	colaPrioridad "TP2-Analisis-Logs/src/log-analyzer/tdas/heap"
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

type LogProcessor struct {
	log         []LogLine
	mostVisited colaPrioridad.Heap[string]
	visitors    diccionario2.DiccionarioOrdenado[string, string]
}

func NewLogProcessor() *LogProcessor {
	return &LogProcessor{
		log: []LogLine{},
	}
}

func compareIP(ip1, ip2 net.IP) int {
	ip1 = ip1.To4()
	ip2 = ip2.To4()
	if ip1 == nil || ip2 == nil {
		return 0
	}
	for i := 0; i < 4; i++ {
		if ip1[i] < ip2[i] {
			return -1
		} else if ip1[i] > ip2[i] {
			return 1
		}
	}
	return 0
}

const layout = "2006-01-02T15:04:05-07:00"

func (lp *LogProcessor) ProcessLogFile(logFile string) {

	file, err := os.Open(logFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando agregar_archivo\n")
		return
	}

	defer file.Close()
	//var logLines []LogLine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		logFields := strings.Split(lines, "\t")
		timestamp, _ := time.Parse(layout, logFields[1])
		logEntry := LogLine{
			ip:        logFields[0],
			Timestamp: timestamp,
			Method:    logFields[2],
			Url:       logFields[3],
		}
		lp.log = append(lp.log, logEntry)
	}
	fmt.Println("OK")
	lp.checkDoS()
	//TODO : encolar mas visitados, guardar ip visitantes
	//
}

func (lp *LogProcessor) ListMostVisited(n string) {
	fmt.Print("Sitios más visitados:")
}
func (lp *LogProcessor) ListVisitors(desde, hasta string) {
	fmt.Print("Sitios más visitados:")
}

func (lp *LogProcessor) checkDoS() {
	//var suspectIps []string
	ipDict := diccionario.CrearHash[string, []time.Time]()
	var suspectIps []net.IP
	for i := 0; i < len(lp.log); i++ { //iterar todas las lineas es O(n)
		ip := lp.log[i].ip
		if !ipDict.Pertenece(ip) {
			ipDict.Guardar(ip, []time.Time{lp.log[i].Timestamp})
		} else {
			timestamp := ipDict.Obtener(ip)
			timestamp = append(timestamp, lp.log[i].Timestamp)
			ipDict.Guardar(ip, timestamp)
			if len(timestamp) >= 5 {
				if timestamp[len(timestamp)-1].Sub(timestamp[len(timestamp)-5]) < 2*time.Second {
					ipParsed := net.ParseIP(ip)
					suspectIps = append(suspectIps, ipParsed)
				}
			}
		}
	}

	colaPrioridad.HeapSort[net.IP](suspectIps, compareIP) //toma O(klog(k)) donde k es la cantidad de ip suspechosas y k << n
	for _, ip := range suspectIps {
		fmt.Println("DoS: ", ip)
	}

}
