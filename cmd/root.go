package cmd

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var (
	// Command line variables
	mileage   int
	from      int
	auctionID int
	workers   int
)

// Vehicle is the meta data for a scraped vehicle
type Vehicle struct {
	Manufacturer   string
	Model          string
	Variant        string
	Mileage        string
	Colour         string
	Transmission   string
	Fuel           string
	BodyType       string
	V5             string
	VAT            string
	PreviousOwners int
	Keys           int
	Url            string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "city-auction-group-scraper",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := colly.NewCollector(colly.AllowedDomains("https://www.cityauctiongroup.com"))
		c.Visit(fmt.Sprintf("https://https://www.cityauctiongroup.com/auctions/%s", auctionID))
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
	rootCmd.Flags().IntVar(&auctionID, "--auction-id", 0, "The auction id to scrape")
	rootCmd.Flags().IntVar(&mileage, "--under-mileage", 200_000, "Exclude vehicles exceeding mileage")
	rootCmd.Flags().IntVar(&from, "--from", 1990, "Exclude vehicles older than this year")
	rootCmd.Flags().IntVar(&workers, "--workers", 0, "The number of asynchronous workers to use")
}
