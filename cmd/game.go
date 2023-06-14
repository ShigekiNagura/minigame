/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gacha/game"

	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("どのゲームで遊びますか？")
		fmt.Println("1: ○×ゲーム")
		fmt.Println("2: ガチャ")

		var n int
		fmt.Scan(&n)

		switch n {
		case 1:
			game.Marubatsu()
		case 2:
			game.Gacha()
		}
	},
}

func init() {
	rootCmd.AddCommand(gameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
