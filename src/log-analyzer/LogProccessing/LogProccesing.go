package LogProccessing

import (
	diccionario2 "TP2-Analisis-Logs/src/log-analyzer/tdas/BST"
	diccionario "TP2-Analisis-Logs/src/log-analyzer/tdas/hash"
	colaPrioridad "TP2-Analisis-Logs/src/log-analyzer/tdas/heap"
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//TODO utilizar como estructurar un hash de ip (no permanente), abbIPS y un heap de visitados

type LogLine struct {
	ip        string
	Timestamp time.Time
	Method    string
	Url       string
}

type LogProcessor struct {
	log         []LogLine
	mostVisited colaPrioridad.ColaPrioridad[VisitedURL]
	visitors    diccionario2.DiccionarioOrdenado[string, string]
}

type VisitedURL struct {
	Url      string
	Cantidad int
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

func ABBcompareIP(ip1, ip2 string) int {
	parsedIp1 := net.ParseIP(ip1)
	parsedIp2 := net.ParseIP(ip2)
	parsedIp1 = parsedIp1.To4()
	parsedIp2 = parsedIp2.To4()
	return compareIP(parsedIp1, parsedIp2)
}

func compareCantidad(url1, url2 VisitedURL) int {
	cantidadURL1, cantidadURL2 := url1.Cantidad, url2.Cantidad
	if cantidadURL1 > cantidadURL2 {
		return 1
	} else if cantidadURL1 < cantidadURL2 {
		return -1
	} else {
		return 0
	}
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
	lp.processVisitedURL()
	lp.processIps()
}

func (lp *LogProcessor) ListMostVisited(n string) {
	k, err := strconv.Atoi(n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando ver_mas_visitados\n")
		return
	}
	fmt.Printf("se pidio n = %d elementos\n", k)
	//TODO:Copiar heap para que desencolar no modifique la estrucutra de datos
	// Controlar que n este en el rango de cantidad de logLines
	fmt.Println("Sitios más visitados:")
	for i := 0; i < k; i++ {
		elem := lp.mostVisited.Desencolar()
		url, cant := elem.Url, elem.Cantidad
		fmt.Printf("\t%s - %d\n", url, cant)
	}
	fmt.Println("OK")
}

func (lp *LogProcessor) ListVisitors(desde, hasta string) {
	iter := lp.visitors.IteradorRango(&desde, &hasta)
	fmt.Println("Visitantes: ")
	for iter.HaySiguiente() {
		elem, _ := iter.VerActual()
		fmt.Printf("\t%s\n", elem)
		iter.Siguiente()
	}
}

func (lp *LogProcessor) processIps() {
	orderedIpDict := diccionario2.CrearABB[string, string](ABBcompareIP)
	for i := 0; i < len(lp.log); i++ {
		ip := lp.log[i].ip
		if !orderedIpDict.Pertenece(ip) {
			orderedIpDict.Guardar(ip, ip)
		}
	}
	lp.visitors = orderedIpDict
}

func (lp *LogProcessor) processVisitedURL() {
	urlDict := diccionario.CrearHash[string, int]()
	for i := 0; i < len(lp.log); i++ {
		url := lp.log[i].Url
		if !urlDict.Pertenece(url) {
			urlDict.Guardar(url, 1)
		} else {
			cantidad := urlDict.Obtener(url)
			cantidad++
			urlDict.Guardar(url, cantidad)
		}
	}
	iter := urlDict.Iterador()
	var urlVisited []VisitedURL
	for iter.HaySiguiente() {
		url, cantidad := iter.VerActual()
		fmt.Println(url, cantidad)
		urlVisited = append(urlVisited, VisitedURL{url, cantidad})
		iter.Siguiente()
	}
	lp.mostVisited = colaPrioridad.CrearHeapArr[VisitedURL](urlVisited, compareCantidad)
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
