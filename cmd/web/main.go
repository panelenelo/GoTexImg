package main

import (
	"GoTexImg/internal/models"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const port string = ":8181"

type application struct {
	Logger *slog.Logger
	Client *http.Client
	TModel *models.TeximgModel
}

func main() {

	dsn, addr := parseFlags()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		Client: &http.Client{Timeout: 5 * time.Second},
		TModel: &models.TeximgModel{DB: db},
		Logger: logger,
	}

	app.Logger.Info("Server starting on localhost", "port", addr)
	err = http.ListenAndServe(addr, app.routes())
	if err != nil {
		app.Logger.Error(err.Error())
	}
	os.Exit(1)
}

func parseFlags() (string, string) {
	addr := flag.String("addr", port, "HTTP network address")

	dbName := flag.String("dbName", os.Getenv("POSTGRES_DB"), "DB name")
	dbPort := flag.String("dbPort", os.Getenv("PGRS_PORT"), "DB port")
	dbHost := flag.String("dbHost", os.Getenv("PGRS_HOST"), "DB host address")
	dbUser := flag.String("dbUser", os.Getenv("POSTGRES_USER"), "DB User")
	dbPass := flag.String("dbPass", os.Getenv("POSTGRES_PASSWORD"), "DB password")
	flag.Parse()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5", *dbHost, *dbPort, *dbUser, *dbPass, *dbName)

	return dsn, *addr
}
