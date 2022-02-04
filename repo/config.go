package repo

import "os"

type Config struct {
	From string
	To []string
	Host string
	Password string
	Port string
	Addres string
}

func NewConfig(to string) Config{
	 cfg :=Config{
		 From: os.Getenv("from"),
		 Host: os.Getenv("host"),
		 Port: os.Getenv("port"),
		 Password: os.Getenv("password"),
		 To : []string{to},
	 }
	 cfg.Addres= cfg.Host+":"+cfg.Port
	 return cfg
}