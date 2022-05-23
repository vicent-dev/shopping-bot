package app

import (
	"github.com/joho/godotenv"
	"shopping-bot/pkg/config"
	"shopping-bot/pkg/telegram"
)

type Server struct {
	b *telegram.Bot
	c *config.Config
}

func NewServer() *Server {
	godotenv.Load()

	c := config.NewConfig()

	return &Server{
		b: telegram.NewBot(c.Bot.Token),
		c: c,
	}
}

func (s *Server) Run() error {
	return s.b.Run()
}
