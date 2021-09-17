package cache

import (
	"log"

	"github.com/impact-eintr/WebKits/encoding"
	"github.com/impact-eintr/ecache/client"
)

func CacheUpdate(key string, value string) {
	b := encoding.Str2bytes(value)
	cmd := &client.Cmd{
		Name:  "set",
		Key:   key,
		Value: b,
	}
	cli, err := client.New("tcp", "127.0.0.1:6430")
	if err != nil {
		log.Println(err)
		return
	}
	cli.Run(cmd)

}

func CacheCheck(key string) []byte {

	cmd := &client.Cmd{
		Name: "get",
		Key:  key,
	}
	cli, err := client.New("tcp", "127.0.0.1:6430")
	if err != nil {
		log.Println(err)
		return nil
	}
	cli.Run(cmd)

	return cmd.Value

}
