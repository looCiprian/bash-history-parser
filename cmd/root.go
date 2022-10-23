package cmd

import (
	"fmt"
	"os"

	"bash-history-parser/internal"

	"github.com/spf13/cobra"
)

var history_file string
var starting_dir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bashHistoryParser",
	Short: "Parser for detect full path of .bash_history file",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		err := internal.Run(history_file, starting_dir)
		if err != nil {
			fmt.Println(err)
		}

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

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bashHistoryParser.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&history_file, "file", "f", "", ".bash_history file to parse (required)")
	rootCmd.MarkFlagRequired("file")

	rootCmd.Flags().StringVarP(&starting_dir, "dir", "d", "home/parser", "Starting dir, usually the user home directory (required)")
	rootCmd.MarkFlagRequired("file")
}
