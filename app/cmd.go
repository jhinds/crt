package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "gcrt",
	Short: "gcrt is a tool to query the Certificate Transparency Logs",
	Long: `gcrt is a tool to query the Certificate Transparency Logs
				  it does so by querying https://crt.sh
				  Complete documentation is available at https://github.com/jhinds/gcrt`,
	Run: func(cmd *cobra.Command, args []string) {
		GetCerts()
	},
}

// Execute runs the application
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const gcrtURL = "https://crt.sh"

type cliOptions struct {
	Domain  string
	Between string
	Days    int
	Count   string
	Output  string
	Offset  string
	Limit   string
}

var opts cliOptions

func init() {
	cmd.PersistentFlags().StringVar(&opts.Between, "between", "", "The dates to run the query for in the format start-date:end-date.  The dates should have the format YYYY-MM-DD")
	cmd.PersistentFlags().StringVarP(&opts.Count, "count", "c", "", "Don't return the results just the count")
	cmd.PersistentFlags().IntVar(&opts.Days, "days", -1, "How many days back to query")
	cmd.PersistentFlags().StringVarP(&opts.Domain, "domain", "d", "", "Domain to find certificates for. % is a wildcard")
	cmd.PersistentFlags().StringVarP(&opts.Output, "output", "o", "json", "The type of output for the certificates")
	cmd.MarkPersistentFlagRequired("domain")
}

// GetCerts will query the Certificate logs and return the result
func GetCerts() {
	validCommand()
	cleanDomain := strings.Replace(opts.Domain, "%", "%25", -1)
	url := fmt.Sprintf("%s/?q=%s&output=json", gcrtURL, cleanDomain)
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	resp, err := client.Get(url)
	if err != nil {
		errors.Wrap(err, "Error Getting Response")
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Wrap(err, "Error Reading Body")
	}
	var certs []CertResponse
	err = json.Unmarshal(contents, &certs)
	if err != nil {
		log.Fatal(err)
		fmt.Print("Error Unmarshalling JSON")
	}

	var filteredCerts []CertResponse
	if opts.Days > -1 {
		dateLookback := time.Now().AddDate(0, 0, -opts.Days)
		for _, cert := range certs {
			if cert.EntryTimestamp.Time.After(dateLookback) {
				filteredCerts = append(filteredCerts, cert)
			}
		}
	} else {
		filteredCerts = certs
	}

	if opts.Output == "text" {
		printTextOutput(filteredCerts)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func printTextOutput(certs []CertResponse) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"COMMON_NAME", "NAME_VALUE", "ENTRY_TIMESTAMP", "ISSUER_NAME", "NOT_AFTER", "NOT_BEFORE"})
	for _, cert := range certs {
		table.Append(cert.ToArray())
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.Render() // Send output

}

func printJSONOutput(certs []CertResponse) {
	fmt.Printf("%v\n", certs)
}

func validCommand() {
	var validOutput bool
	switch opts.Output {
	case "text",
		"json":
		validOutput = true
	default:
		validOutput = false

	}
	if !validOutput {
		log.Fatalf("%s is not a valid option for output. valid options: text,json", opts.Output)
	}
}
