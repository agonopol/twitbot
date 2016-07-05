package lib

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"log"
)

type TwitterBot struct {
	client *anaconda.TwitterApi
	yapper Yapper
	logger *log.Logger
}

func NewTwitterBot() *TwitterBot {
	return nil
}

func (this *TwitterBot) Execute() {

	msg := this.yapper.Next()
	if len(msg) > 144 {
		this.logger.Println(fmt.Sprintf("[%s] exceeds the maximum tweet length of 144 chars", msg))
		return
	}

	this.client.PostTweet(status, v)
}
