package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mattn/go-mastodon"
	"github.com/mopeneko/donhaialert/api/database"
	"github.com/mopeneko/donhaialert/api/domain"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	c := cron.New(cron.WithLocation(jst))
	c.AddFunc("0 0 * * *", task)
	c.Run()
}

func task() {
	database.Init()

	users := []domain.User{}
	err := database.DB.
		Preload("AccessToken").
		Preload("Credential").
		Find(&users).
		Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		panic(fmt.Sprintf("an unexpected error was occurred: %+v\n", err))
	}

	ctx := context.Background()

	for _, user := range users {
		config := mastodon.Config{
			Server: "https://" + user.Credential.Host,
			ClientID: user.Credential.ClientID,
			ClientSecret: user.Credential.ClientSecret,
			AccessToken: user.AccessToken.AccessToken,
		}
		client := mastodon.NewClient(&config)

		account, err := client.GetAccountCurrentUser(ctx)
		if err != nil {
			log.Println("failed to get account info:", user.MastodonID, "on", user.Credential.Host)
			continue
		}

		diff := fmt.Sprintf(
			"トゥート: %d(%+d)\n" +
				"フォロー: %d(%+d)\n" +
				"フォロワー: %d(%+d)\n" +
				"\n" +
				"https://donhaialert.com\n" +
				"#donhaialert",
				account.StatusesCount, account.StatusesCount - user.TootsCount,
				account.FollowingCount, account.FollowingCount - user.FollowsCount,
				account.FollowersCount, account.FollowersCount - user.FollowersCount,
			)

		toot := mastodon.Toot{
			Status: diff,
			Visibility: "unlisted",
		}

		_, err = client.PostStatus(ctx, &toot)
		if err != nil {
			log.Println("failed to toot:", user.MastodonID, "on", user.Credential.Host)
		}
	}
}
