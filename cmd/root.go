/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bombar",
	Short: "put the command u wanna run and fire",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		r, _ := strconv.Atoi(cmd.Flag("repeat").Value.String())
		if r <= 0 {
			fmt.Println("WARNING: repeat flag cannot be 0 or less, changed to 50 ", r)
			r = 50
		}
		cm := args[0]
		arg := args[1:]
		bombar(r, cm, arg...)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bombar.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().Int32P("repeat", "r", 50, "Amount of times the command will run")
}

func bombar(r int, cmd string, args ...string) {
	now := time.Now()
	for i := 0; i < r; i++ {
		c := exec.Command(cmd, args...)
		_, err := c.StderrPipe()
		if err != nil {
			panic(err)
		}
		_, err = c.StdinPipe()
		if err != nil {
			panic(err)
		}
		err = c.Start()
		if err != nil {
			panic(err)
		}
		c.Wait()
	}
	fmt.Printf("Spent: %v", time.Since(now))
}
