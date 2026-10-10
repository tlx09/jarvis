package jarvis

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var jarvisCmd = &cobra.Command{
	Use:   "ask",
	Short: "ask Jarvis",
	RunE: func(cmd *cobra.Command, args []string) error {
		question := viper.GetString("question")
		fmt.Printf("The answer to %s is 42", question)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(jarvisCmd)
	jarvisCmd.Flags().String("question", "What's the answer to life", "User question")
}


