package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "damv",
		Short: "Moves and renames files",
		// Long: ` `,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// func init() {
// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

// rootCmd.AddCommand(addCmd)
// rootCmd.AddCommand(initCmd)
// }
