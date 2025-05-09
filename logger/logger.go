package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

var LogFile *os.File

const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	CYAN   = "\033[36m"
	RESET  = "\033[0m"
)

const (
	SUCCESS = "🟢"
	INFO    = "🔵"
	WARN    = "🟡"
	ERROR   = "🔴"
	FATAL   = "💀"
)

// ==================== fonction logger ==================== //

func callerInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "???"
	}
	fn := runtime.FuncForPC(pc)
	funcName := filepath.Base(fn.Name())
	fileName := filepath.Base(file)
	return fmt.Sprintf("%s:%s:%d", fileName, funcName, line)
}

func defineTypeLogger(level string) (string, string) {
	switch level {
	case "SUCCESS":
		return SUCCESS, GREEN
	case "INFO":
		return INFO, CYAN
	case "WARN":
		return WARN, YELLOW
	case "ERROR":
		return ERROR, RED
	case "FATAL":
		return FATAL, RED
	default:
		return "", ""
	}
}

func logWithLocation(level string, msg string) {
	location := callerInfo(3)
	emoji, color := defineTypeLogger(level)
	stringFormat := fmt.Sprintf("%s%s [%s] [%s]: %v%s", color, emoji, level, location, RESET, msg)
	log.Println(stringFormat)
}

func logfWithLocation(level string, format string, args ...any) {
	location := callerInfo(3)
	emoji, color := defineTypeLogger(level)
	log.Printf("%s%s [%s] [%s]: %v%s", color, emoji, level, location, RESET, fmt.Sprintf(format, args...))
}

func S(msg string) {
	logWithLocation("SUCCESS", msg)
}

func I(msg string) {
	logWithLocation("INFO", msg)
}

func W(msg string) {
	logWithLocation("WARN", msg)
}

func E(msg string) {
	logWithLocation("ERROR", msg)
}

func F(msg string) {
	location := callerInfo(3)
	emoji, color := defineTypeLogger("FATAL")
	log.Fatalf("%s%s [FATAL] [%s] %s %v", color, emoji, location, msg, RESET)
}

func Sf(format string, args ...any) {
	logfWithLocation("SUCCESS", format, args...)
}

func If(format string, args ...any) {
	logfWithLocation("INFO", format, args...)
}

func Wf(format string, args ...any) {
	logfWithLocation("WARN", format, args...)
}

func Ef(format string, args ...any) {
	logfWithLocation("ERROR", format, args...)
}

func Ff(format string, args ...any) {
	location := callerInfo(3)
	emoji, color := defineTypeLogger("FATAL")
	log.Fatalf("%s%s [FATAL] [%s] %s %v", color, emoji, location, fmt.Sprintf(format, args...), RESET)
}

// ==================== couleur & writer ==================== //

type dualLogger struct {
	stdout    io.Writer
	file      io.Writer
	withColor bool
}

func (d *dualLogger) Write(p []byte) (n int, err error) {
	if d.withColor {
		_, _ = d.stdout.Write(p)
	} else {
		_, _ = d.stdout.Write(stripColor(p))
	}

	return d.file.Write(stripColor(p))
}

var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripColor(input []byte) []byte {
	return ansiRegexp.ReplaceAll(input, []byte(""))
}

// ==================== init ==================== //

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	log.Fatal("Impossible de trouver la racine du projet")
	return ""
}

func InitFromEnv(env string) {
	projectRoot := findProjectRoot()
	log.Printf("%s [INFO] Projet racine trouvé : %s %v", CYAN, projectRoot, RESET)

	logDir := filepath.Join(projectRoot, "tmp", "log", env)
	log.Printf("%s [INFO] Log Dir : %s", CYAN, logDir)

	logPath := filepath.Join(logDir, "log-"+time.Now().Format("2006-01-02")+".log")
	log.Printf("%s [INFO] Log Path : %s", CYAN, logPath)

	emojiS, colorS := defineTypeLogger("SUCCESS")
	emojiI, colorI := defineTypeLogger("INFO")
	emojiF, colorF := defineTypeLogger("FATAL")

	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatalf("%s%s [FATAL] Impossible de créer le chemin de log requis : %v %v", colorF, emojiF, err, RESET)
	}

	log.Printf("%s%s [INFO] Création du fichier de log à l’emplacement : %s %v", colorI, emojiI, logPath, RESET)

	LogFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%s%s [FATAL] Impossible d’accéder au fichier de log : %v %v", colorF, emojiF, err, RESET)
	}

	log.Printf("%s%s [INFO] Fichier de log ouvert à l’emplacement : %s %v", colorI, emojiI, logPath, RESET)

	dualWriter := &dualLogger{
		stdout:    os.Stdout,
		file:      LogFile,
		withColor: true,
	}
	log.SetOutput(dualWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("%s%s [SUCCESS] Logger initialisé avec succès. Fichier : %s %v", colorS, emojiS, logPath, RESET)
}

func Init(env string) {
	if env == "" {
		env = "dev"
	}

	InitFromEnv(env)
}

func Close() {
	if LogFile != nil {
		LogFile.Close()
	}
}
