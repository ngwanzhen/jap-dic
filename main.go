package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func save(i string) {
	// v := convertToStringSlice(i)
	// v := convertToStruct(i)
	v := convertToMap(i)
	v.convertToString()
	// print(v)
	saveToFile(v)

	// shapeMain()
	// file_reader()
	// checkAnimals()
}

func main() {

	// app := &cli.App{
	// 	Name:  "JapDic",
	// 	Usage: "saves a jap vocab",
	// 	Action: func(c *cli.Context) error {
	// 		// fmt.Printf("Hello %q", c.Args().Get(0))
	// 		i := c.Args().Get(1)
	// 		// fmt.Println(i)
	// 		save(i)
	// 		return nil
	// 	},
	// }

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "save",
				Aliases: []string{"s"},
				Usage:   "saves a jap vocab",
				Action: func(cCtx *cli.Context) error {
					s := ""
					for i := 0; i < cCtx.NArg(); i++ {
						s += cCtx.Args().Get(i) + " "
					}
					save(strings.TrimSpace(s))
					fmt.Println("saved", s)
					return nil
				},
			},
			{
				Name:    "find",
				Aliases: []string{"f"},
				Usage:   "find a word from the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Finding word: ", cCtx.Args().First())
					line := find(cCtx.Args().First())
					if line == "" {
						fmt.Println("Nothing found")
					} else {
						fmt.Println("Found word: ", line)
					}
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read all",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("showing all...")
					lines := readFromFile()
					for _, l := range lines {
						fmt.Println(l)
					}
					// fmt.Println(s)
					return nil
				},
			},
			{
				Name:    "translate",
				Aliases: []string{"tr"},
				Usage:   "translate",
				Action: func(cCtx *cli.Context) error {
					a := cCtx.Args().First()
					fmt.Println("Translate word: ", a)
					if isJap(a) {
						s := translate(a, "ja")
						fmt.Println(s)
					} else {
						s := translate(a, "en")
						fmt.Println(s)
					}
					return nil
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "randomly show 1 jap or eng word",
				Action: func(cCtx *cli.Context) error {
					s := readFromFile()
					r1, r2 := randomVocab(s)
					fmt.Println(r1)
					time.Sleep(8 * time.Second)
					fmt.Println(r2)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
