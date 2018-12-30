package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Println("BOOM!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
} /*
---------------------
作者：turkeycock
来源：CSDN
原文：https://blog.csdn.net/TurkeyCock/article/details/80359654
版权声明：本文为博主原创文章，转载请附上博文链接！*/
