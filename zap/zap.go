package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type user struct {
	Name      string
	Email     string
	CreatedAt time.Time
}

func (u user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Name)
	enc.AddString("email", u.Email)
	enc.AddInt64("created_at", u.CreatedAt.UnixNano())
	return nil
}

func main() {
	logger, _ := zap.NewDevelopment()
	user := &user{
		Name: "Zap",
		Email: "zap@sample.com",
		CreatedAt: time.Now(),
	}
	logger.Info("object sample", zap.Object("userObj", user))
}
