package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"git.nothr.in/nothr/g-electra/pkg/database"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// create koanf
	k := koanf.New(".")

	cfgPath := flag.String("config", "config.toml", "config file for the app")
	flag.Parse()

	// Load config file
	if err := k.Load(file.Provider(*cfgPath), toml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	//Initilize logger
	lgrOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if k.Bool("app.debug") {
		lgrOptions.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, lgrOptions))
	logger.Info("Loaded Config file")

	// Initialise DB
	err := database.ConnectDB(k, logger)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialise Echo
	e := echo.New()
	e.HideBanner = true

	// serve static files
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "./ui/dist/",
		Index: "index.html",
		HTML5: true,
	}))

	e.Start(":3000")

}
