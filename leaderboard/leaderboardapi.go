package leaderboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	redis "gopkg.in/redis.v5"
)

const getTopScoresPath = "/leaderboard/gettopscores"
const getRankPath = "/leaderboard/getrank"
const getScorePath = "/leaderboard/getscore"
const postScorePath = "/leaderboard/postscore"

type topScores struct {
	Scores []redis.Z
}

// Bind wires up the end points for the api
func Bind() {
	http.HandleFunc(getTopScoresPath, getTopScoresHandler)
	http.HandleFunc(getRankPath, getRankHandler)
	http.HandleFunc(getScorePath, getScoreHandler)
	http.HandleFunc(postScorePath, postScoreHandler)
}

func getTopScoresHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %q: %q\n", strings.ToUpper(r.Method), r.URL.Path)
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		panic(err)
	}
	scores := topScores{getTopScores(int64(count))}
	bytes, err := json.Marshal(scores)

	fmt.Fprintf(w, "%s", string(bytes))
}

func getRankHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %q: %q\n", strings.ToUpper(r.Method), r.URL.Path)
	userName := r.URL.Query().Get("user")
	fmt.Fprintf(w, "%d", getRank(userName))
}

func getScoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %q: %q\n", strings.ToUpper(r.Method), r.URL.Path)
	userName := r.URL.Query().Get("user")
	fmt.Fprintf(w, "%d", getRank(userName))
}

func postScoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %q: %q\n", strings.ToUpper(r.Method), r.URL.Path)
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var scoreEntry redis.Z
	err = json.Unmarshal(bytes, &scoreEntry)
	if err != nil {
		panic(err)
	}
	postScore(scoreEntry)
}
