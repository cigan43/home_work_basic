package main

import (
	"github.com/labstack/gommon/log"
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
	Host      string `cfg_name:"host" cfg_desc:"Server host"`
	Port      int    `cfg_name:"port" cfg_desc:"Server port"`
	SearchAPI string `cfg_name:"apis.search" cfg_desc:"Search API endpoint"`
}

// func config()
var (
	logAnalyzerFile   string
	logAnalyzerLevel  string
	logAnalyzerOutput string
	showHelp          bool
)

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

	log.Infof("configPath == %s", logAnalyzerFile)
}
