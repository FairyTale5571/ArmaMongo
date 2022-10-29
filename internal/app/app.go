package app

import (
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/fairytale5571/mongo/pkg/logger"
	"github.com/fairytale5571/mongo/pkg/models"
	"github.com/fairytale5571/mongo/pkg/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type App struct {
	Logger *logger.Wrapper
	DB     *mongo.Mongo
	Config models.Config
}

var app *App

func NewApp() (*App, error) {
	a := &App{}
	a.Logger = logger.New("mongo_app")
	if err := env.Parse(&a.Config); err != nil {
		a.Logger.Errorf("error parse config: %v", err)
		return nil, err
	}

	db, err := mongo.New(a.Config.MongoURL)
	if err != nil {
		a.Logger.Fatalf("error create database: %v", err)
		return nil, err
	}
	a.DB = db
	a.Logger.Infof("mongo started %s", a.DB.Version())
	return a, nil
}

func Version(a []string) string {
	return "0.0.1"
}

func Setup(a []string) string {
	if app == nil {
		a, err := NewApp()
		if err != nil {
			return err.Error()
		}
		app = a
	}
	return "done"
}

func Write(a []string) string {
	if app == nil {
		return "setup not done"
	}
	if len(a) < 2 {
		return "not enough arguments"
	}

	payload := bson.M{
		"player_uid": a[1],
		"type_log":   a[2],
		"created_at": time.Now().UnixNano(),
		"log":        bson.M{},
	}

	log := strings.Split(a[3], "~~")
	for _, v := range log {
		s := strings.Split(v, "^^")
		if len(s) == 2 {
			payload["log"].(bson.M)[s[0]] = s[1]
		}
	}

	err := app.DB.Write(app.Config.Database, a[0], payload)
	if err != nil {
		return err.Error()
	}
	return "done"
}
