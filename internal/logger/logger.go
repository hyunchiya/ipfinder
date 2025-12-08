package logger

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	Silent  bool
	Verbose bool
	NoColor bool
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.Silent {
		return
	}
	timestamp := time.Now().Format("15:04:05")
	msg := fmt.Sprintf(format, args...)
	if l.NoColor {
		fmt.Printf("[%s] %s\n", timestamp, msg)
	} else {
		fmt.Printf("\033[90m[%s]\033[0m \033[97m%s\033[0m\n", timestamp, msg)
	}
}

func (l *Logger) Success(source, ip string, count int) {
	if l.Silent {
		return
	}
	timestamp := time.Now().Format("15:04:05")
	if l.NoColor {
		fmt.Printf("[%s] %s %s %d\n", timestamp, source, ip, count)
	} else {
		fmt.Printf("\033[90m[%s]\033[0m \033[92m%s\033[0m \033[97m%s\033[0m \033[36m%d\033[0m\n",
			timestamp, source, ip, count)
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	msg := fmt.Sprintf(format, args...)
	if l.NoColor {
		fmt.Printf("[%s] %s\n", timestamp, msg)
	} else {
		fmt.Printf("\033[90m[%s]\033[0m \033[91m%s\033[0m\n", timestamp, msg)
	}
}

func (l *Logger) Warning(format string, args ...interface{}) {
	if l.Silent {
		return
	}
	timestamp := time.Now().Format("15:04:05")
	msg := fmt.Sprintf(format, args...)
	if l.NoColor {
		fmt.Printf("[%s] %s\n", timestamp, msg)
	} else {
		fmt.Printf("\033[90m[%s]\033[0m \033[93m%s\033[0m\n", timestamp, msg)
	}
}

func (l *Logger) Verbosef(format string, args ...interface{}) {
	if !l.Verbose || l.Silent {
		return
	}
	timestamp := time.Now().Format("15:04:05")
	msg := fmt.Sprintf(format, args...)
	if l.NoColor {
		fmt.Printf("[%s] %s\n", timestamp, msg)
	} else {
		fmt.Printf("\033[90m[%s]\033[0m \033[90m%s\033[0m\n", timestamp, msg)
	}
}

func (l *Logger) Line() {
	if l.Silent {
		return
	}
	fmt.Println()
}

func (l *Logger) Stat(char string, length int) {
	if l.Silent {
		return
	}
	line := strings.Repeat(char, length)
	if l.NoColor {
		fmt.Println(line)
	} else {
		fmt.Printf("\033[90m%s\033[0m\n", line)
	}
}
