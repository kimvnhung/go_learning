package main

import (
	"fmt"
	"log"

	"github.com/kimvnhung/golio"
	"github.com/kimvnhung/golio/api"
	"github.com/sirupsen/logrus"
)

func main() {
	client := golio.NewClient("RGAPI-348e79d5-07c7-4d94-99b4-d59d544c6d40",
		golio.WithRegion(api.RegionVietnam),
		golio.WithLogger(logrus.New().WithField("method", "endpoint")))
	account, err := client.Riot.Account.GetByRiotID("HeroKimGaming", "#1222")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Account: %s\n", account.Puuid)
}
