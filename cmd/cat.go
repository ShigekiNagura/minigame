/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"gacha/aa"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print the ascii art of cat",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := aa.Aa.Open("cat.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		// ファイルを読み込んで出力
		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		fmt.Println(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
