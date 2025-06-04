package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "FileProcessor",
		Usage: "Lê um arquivo JSON e extrai dados",
		Action: func(c *cli.Context) error {
			fs := afero.NewOsFs()
			filename := c.Args().First()
			if filename == "" {
				return fmt.Errorf("forneça o caminho do arquivo")
			}

			data, err := afero.ReadFile(fs, filename)
			if err != nil {
				return err
			}

			result := gjson.Get(string(data), "nome")
			fmt.Println("Nome encontrado no JSON:", result.String())
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}