package twitbot

import "github.com/kurrik/twittergo"

type TwitterBot struct {
	client *twittergo.Client
	yapper Yapper
}

func NewTwitterBot() *TwitterBot {

	return nil
}

func (this *TwitterBot) Execute() {

	msg := this.yapper.Next()

	if len(msg) > 144 {
		return
	}

}
