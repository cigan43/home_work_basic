package main

import (
	//"github.com/labstack/gommon/log"
	"bufio"
	"fmt"
	"log"
	"os"

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

// func config()
var (
	logAnalyzerFile   string
	logAnalyzerLevel  string
	logAnalyzerOutput string
	showHelp          bool
)

func (cfg *appConfig) ConfigFile(value string) {
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
		cfg.File = os.Getenv("LOG_ANALYZER_LEVEL")
	default:
		cfg.File = value
	}
}

func (cfg *appConfig) ConfigOutput(value string) {
	switch value {
	case "":
		cfg.File = os.Getenv("LOG_ANALYZER_OUTPUT")
	default:
		cfg.File = value
	}
}

func ReadFile(logfile string) {
	// fmt.Println(logfile)
	file, err := os.Open(logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
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

	// fmt.Println(logAnalyzerFile)
	c := appConfig{}
	c.ConfigFile(logAnalyzerFile)
	c.ConfigLevel(logAnalyzerLevel)
	c.ConfigOutput(logAnalyzerOutput)
	fmt.Print(c.File)
	// if c.File == "" {
	// 	panic("not log file")
	// }

	// if c.Level == "" {
	// 	c.Level = "info"
	// }

	// if c.Output == "" {
	// 	c.Output = "logout/"
	// }

	// fmt.Println(c.File)
	// ReadFile(c.File)
}
