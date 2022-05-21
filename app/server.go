package app

import (
	"shopping-bot/pkg/config"
	"shopping-bot/pkg/telegram"
)

type Server struct {
	b *telegram.Bot
	c *config.Config
}

func NewServer() *Server {
	c := config.NewConfig()

	return &Server{
		b: telegram.NewBot(c.Bot.Token),
		c: c,
	}
}

func (s *Server) Run() error {
	return s.b.Run()
}
