package main

import (
	//"github.com/labstack/gommon/log"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

// Вы разрабатываете командную утилиту для анализа лог-файлов. Утилита должна принимать на вход путь к лог-файлу, анализировать его содержимое и выводить статистику по различным параметрам логов.

// Требования:

// Утилита должна поддерживать следующие флаги командной строки:

// -file <путь_к_лог_файлу>: указывает путь к анализируемому лог-файлу (обязательный флаг).
// -level <уровень_логов>: указывает уровень логов для анализа (необязательный флаг, значение по умолчанию - ""info"").
// -output <путь_к_файлу>: указывает путь к файлу, в который будет записана статистика (необязательный флаг, если не указан, статистика выводится в стандартный поток вывода).
// Утилита должна обрабатывать переменные окружения:

// LOG_ANALYZER_FILE: путь к анализируемому лог-файлу (если не указан через флаг -file).
// LOG_ANALYZER_LEVEL: уровень логов для анализа (если не указан через флаг -level).
// LOG_ANALYZER_OUTPUT: путь к файлу для записи статистики (если не указан через флаг -output).
// Утилита должна анализировать лог-файл и собирать статистику по указанному уровню логов (или по уровню по умолчанию). Формат и содержание статистики определяйте на свое усмотрение.

// Утилита должна выводить статистику либо в указанный файл (если указан флаг -output), либо в стандартный поток вывода.
// Напишите юнит тесты на реализованные функции;

type appConfig struct {
	File   string
	Level  string
	Output string
}

type status struct {
	ip     string
	status string
	method string
	engine string
}

func (cfg *appConfig) ConfigFile(value string) {
	// fmt.Println(value)
	switch {
	case value == "":
		cfg.File = os.Getenv("LOG_ANALYZER_FILE")
	default:
		cfg.File = value
	}
}

func (cfg *appConfig) ConfigLevel(value string) {
	switch value {
	case "":
		cfg.Level = os.Getenv("LOG_ANALYZER_LEVEL")
	default:
		cfg.Level = value
	}
}

func (cfg *appConfig) ConfigOutput(value string) {
	switch value {
	case "":
		cfg.Output = os.Getenv("LOG_ANALYZER_OUTPUT")
	default:
		cfg.Output = value
	}
}

func ReadFile(logfile string) ([]string, error) {
	words := []string{}
	file, err := os.Open(logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return words, nil
}

func (s *status) Add(ip, method, status, engine string) {
	s.ip = ip
	s.method = clearstring(method)
	s.status = status
	s.engine = clearstring(engine)
}

func clearstring(st string) string {
	return strings.ReplaceAll(st, "\"", "")
}

func sort(sl []status) map[string]map[string]int64 {
	m := make(map[string]map[string]int64)
	m["ip"] = make(map[string]int64)
	m["method"] = make(map[string]int64)
	m["status"] = make(map[string]int64)
	m["engine"] = make(map[string]int64)
	for s := range sl {
		_, ok := m["ip"][sl[s].ip]
		if ok == true {
			m["ip"][sl[s].ip]++
		} else {
			m["ip"][sl[s].ip] = 1
		}
		_, ok_method := m["method"][sl[s].method]

		if ok_method == true {
			m["method"][sl[s].method]++
		} else {
			m["method"][sl[s].method] = 1
		}
		_, ok_engine := m["engine"][sl[s].engine]
		if ok_engine == true {
			m["engine"][sl[s].engine]++
		} else {
			m["engine"][sl[s].engine] = 1
		}

		_, ok_status := m["status"][sl[s].status]
		if ok_status == true {
			m["status"][sl[s].status]++
		} else {
			m["status"][sl[s].status] = 1
		}

	}
	return m
}

func main() {
	var (
		logAnalyzerFile   string
		logAnalyzerLevel  string
		logAnalyzerOutput string
		showHelp          bool
	)

	pflag.StringVarP(&logAnalyzerFile, "file", "i", "",
		"Input file")
	pflag.StringVarP(&logAnalyzerLevel, "level", "l", "",
		"log level")
	pflag.StringVarP(&logAnalyzerOutput, "output", "o", "",
		"Out file")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
	}

	c := appConfig{}
	c.ConfigFile(logAnalyzerFile)
	c.ConfigLevel(logAnalyzerLevel)
	c.ConfigOutput(logAnalyzerOutput)

	if c.File == "" {
		panic("not log file")
	}

	if c.Level == "" {
		c.Level = "info"
	}

	if c.Output == "" {
		c.Output = "logout/"
	}

	// var st []status
	// struc := status{}
	stringlog, err := ReadFile(c.File)
	if err != nil {
		fmt.Println(err)
	}
	for s := range stringlog {
		fmt.Println(stringlog[s])
		// fmt.Println(strings.Split(stringlog[s], " "))
		// struc.Add(strings.Split(stringlog[s], " ")[0],
		// strings.Split(stringlog[s], " ")[5],
		// strings.Split(stringlog[s], " ")[8],
		// strings.Split(stringlog[s], " ")[11])
		// st == append(st, struc)
	}
	// rr := sort(st)
	// fmt.Println(rr)
	// for r := range rr["ip"] {
	// 	fmt.Println(r, rr["ip"][r])
	// }
}
