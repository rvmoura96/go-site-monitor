package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 5
const delay = 5

func main() {
	showHeader()
	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			startMonitor()
		case 2:
			readLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida.")
			os.Exit(-1)
		}

	}

}

func showHeader() {
	name := "Renan" //:= atribuição de var
	version := 1.2
	fmt.Println("Olá, sr(a)", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido é ", command)
	fmt.Println("")
	return command
}

func startMonitor() {
	fmt.Println("Monitorando")
	sites := readSites()
	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, "endereço:", site)
			getSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func getSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro ao acessar: ", site, "Codigo do erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso :D")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "está com probelama. Status Code:", resp.StatusCode)
		logRegister(site, false)
	}
}

func readSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	file.WriteString(time.Now().Format("02/01/2006  15:04:05 - ") + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()

}

func readLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println(string(file))
}
