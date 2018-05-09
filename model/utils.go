package model

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	client = &twitterClient{}
)

// SetUp twitterクライアントを起動する。今のところは自分用
func SetUp() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_KEY_SECRET"))
	_api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	client = &twitterClient{
		api: _api,
	}
}

// GetTwitterClient ツイッタークライアントを取得する
func GetTwitterClient() TwitterClient {
	return client
}
