package cmd

import (
	"github.com/MelkoV/go-learn-api/api"
	"github.com/MelkoV/go-learn-common/app"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run API JsonRPC server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("api.port")
		l := logger.NewCategoryLogger("jsonrpc/api", logger.NewStreamLog())
		l.Format("init", app.SYSTEM_UUID, "starting API server on port %d", port).Info()
		api.Serve(port, l)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
