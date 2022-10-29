package models

type Config struct {
	MongoURL string `env:"MONGO_URL,required"`
	Database string `env:"DATABASE_MONGO,required" envDefault:"logs"`
}

type Document struct {
	Log     map[string]string `json:"log"`
	UID     string            `json:"uid"`
	Service string            `json:"service"`
}
