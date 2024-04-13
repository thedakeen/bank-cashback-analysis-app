package main

import (
	mongoDB "bank-cashback-analysis/backend/pkg/models/mongodb"
	"context"
	"flag"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	users    *mongoDB.UserModel
	otps     *mongoDB.OtpModel
	promos   *mongoDB.PromoModel
	cards    *mongoDB.CardModel
}

func main() {
	addr := flag.String("addr", ":7777", "HTTP networks address")
	mongoURI := flag.String("mongoURI", "mongodb+srv://user:<password>@cluster1.wdbaku4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1", "MongoDB URI")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*mongoURI))
	if err != nil {
		errorLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			errorLog.Fatal(err)
		}
	}()

	if err := client.Ping(context.TODO(), nil); err != nil {
		errorLog.Fatal(err)
	}

	db := client.Database("BCAapp")

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		otps:     mongoDB.NewOtpModel(db.Collection("otps")),
		users:    mongoDB.NewUserModel(db.Collection("users")),
		promos:   mongoDB.NewPromoModel(db.Collection("promos")),
		cards:    mongoDB.NewCardModel(db.Collection("cards")),
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),

		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		c := cron.New()

		app.promos.DropCollection()
		app.kaspiParser()
		app.insertHalyk(app.promos)

		c.AddFunc("0 0 * * *", func() {
			app.promos.DropCollection()
			app.kaspiParser()
			app.insertHalyk(app.promos)
		})
		c.Start()

		select {}
	}()

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
