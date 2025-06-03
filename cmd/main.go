package main

import (
	"fmt"
	"goslog"
	"log/slog"
	"os"
	"time"
)

func main() {
	fmt.Println("Default Logger:")
	logTest(goslog.DefaultLogger(os.Stdout))

	fmt.Println("\nCustom Logger:")
	logTest(goslog.CustomLogger(os.Stdout))
}

type address struct {
	Street string
	City   string
	State  string
}

type data struct {
	Name  string
	Age   int
	Email string
	Addr  *address
}

var datas = []data{
	{Name: "Alice", Age: 30, Email: "alice@email", Addr: &address{
		Street: "123 Main St",
		City:   "Anytown",
		State:  "CA",
	}},
	{Name: "Bob", Age: 25, Email: "bob@email", Addr: &address{
		Street: "456 Elm St",
		City:   "Othertown",
		State:  "TX",
	}},
	{Name: "Charlie", Age: 35, Email: "charlie@email", Addr: nil},
}

func logTest(l *slog.Logger) {
	l.Debug("Debug message", slog.String("key", "value"))
	l.Info("Info message", slog.String("key", "value"))
	l.Warn("Warning message", slog.String("key", "value"))
	l.Error("Error message", slog.String("key", "value"))

	w := l.With(slog.String("key", "value"))
	w.Info("Info message and with")
	w.Info("Info message and with attr", slog.String("key", "value2"))

	g := l.WithGroup("group1")
	g.Info("Info message in group", slog.String("key", "value3"))

	t := time.Now()

	l.Info("Info message and \"any\" data",
		"bool", true,
		"time", t,
		slog.Any("time2", t),
		slog.Any("key", "value4"),
		slog.Any("array", []string{
			"January",
			"February",
			"March",
			"April",
			"May",
			"June",
			"July",
			"August",
			"September",
			"October",
			"November",
			"December",
		}),
		slog.Any("primes", []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31}),
		"primes2", []int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31},
		slog.Any("users", datas),
	)
}
