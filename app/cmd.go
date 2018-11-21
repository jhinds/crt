package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "crt",
	Short: "crt is a tool to query the Certificate Transparency Logs",
	Long: `crt is a tool to query the Certificate Transparency Logs
				  it does so by querying https://crt.sh
				  Complete documentation is available at https://github.com/jhinds/crt`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var domain string

var between string

var days string

func init() {
	cmd.PersistentFlags().StringVar(&domain, "domain", "", "Domain to find certificates for. % is a wildcard")
	cmd.PersistentFlags().StringVar(&days, "days", "", "How many days back to query")
	cmd.PersistentFlags().StringVar(&between, "between", "", "The dates to run the query for in the format start-date:end-date.  The dates should have the format YYYY-MM-DD")
}
