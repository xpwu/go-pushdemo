package main

import (
	"context"
	"github.com/xpwu/go-cmd/arg"
	"github.com/xpwu/go-cmd/cmd"
	"github.com/xpwu/go-log/log"
	"github.com/xpwu/go-streamclient/pushc"
	"time"
)

func main() {
	cmd.RegisterCmd("push", "push data to client", func(args *arg.Arg) {
		pushUrl := ""
		args.String(&pushUrl, "u", "push url")
		data := ""
		args.String(&data, "d", "data")

		args.Parse()

		err := pushc.PushData(context.TODO(), pushUrl, []byte(data), 10*time.Second)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("push ok")
		}
	})

	cmd.RegisterCmd("close", "close client", func(args *arg.Arg) {
		pushUrl := ""
		args.String(&pushUrl, "u", "push url")

		args.Parse()

		err := pushc.Close(context.TODO(), pushUrl)
		if err != nil {
			log.Error(err)
		} else {
			log.Info("close ok")
		}
	})

	cmd.Run()
}
