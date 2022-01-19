package main

import (
	"fmt"
	"golog"
)

func main() {
	golog.LogLevel(6)
	golog.LogTimeFormat(golog.ANSIC)
	print()
	golog.LogFormat(1)
	print()
	golog.LogFormat(2)
	print()
	golog.LogFormat(3)
	print()
	golog.LogFormat(4)
	print()
	golog.LogTimeFormat(golog.RFC822Z)
	golog.LogFormat(3)
	print()
	golog.LogInfoFormat(4)
	print()

}

func print() {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	golog.Fatal("Fatal message")
	golog.Error("Error message")
	golog.Warning("Warning message")
	golog.Info("Info message")
	golog.Debug("Debug message")
	golog.Trace("Trace message")
}
