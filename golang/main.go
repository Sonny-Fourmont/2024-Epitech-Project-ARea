package main

import (
	"demo/nasa"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron"
)

func main() {
	apiKey := os.Getenv("NASA_API_KEY")
	if apiKey == "" {
		log.Fatal("NASA_API_KEY introuvable dans le fichier .env")
	}

	c := cron.New()

	c.AddFunc("1 * * * *", func() {
		apod, err := nasa.GetNasaAPOD(apiKey)
		if err != nil {
			log.Printf("Erreur lors de la récupération de l'image APOD: %v", err)
			return
		}
		fmt.Println("Nouvelle requête APOD à", time.Now())
		fmt.Println("Titre :", apod.Title)
		fmt.Println("Date :", apod.Date)
		fmt.Println("URL :", apod.URL)
		fmt.Println("Description :", apod.Explanation)
	})

	c.Start()

	select {}
}
