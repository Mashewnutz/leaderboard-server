package leaderboard

import (
	"fmt"

	redis "gopkg.in/redis.v5"
)

var redisClient *redis.Client

// Init initialises the redis connection
func Init() error {
	redisClient = newClient()
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func getTopScores(count int64) []redis.Z {
	fmt.Println("getTopScores:", count)
	scores, err := redisClient.ZRevRangeWithScores("leaderboard", 0, count-1).Result()
	if err != nil {
		panic(err)
	}
	return scores
}

func getRank(name string) int64 {
	fmt.Println("getRank:", name)
	rank, err := redisClient.ZRevRank("leaderboard", name).Result()
	if err != nil {
		panic(err)
	}
	return rank
}

func getScore(name string) float64 {
	fmt.Println("getScore:", name)
	score, err := redisClient.ZScore("leaderboard", name).Result()
	if err != nil {
		panic(err)
	}
	return score
}

func postScore(entry redis.Z) {
	fmt.Println("postScore:", entry)
	redisClient.ZAdd("leaderboard", entry)
}
