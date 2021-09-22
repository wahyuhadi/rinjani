package main

import (
	"flag"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/rinjani/directory"
	"github.com/wahyuhadi/rinjani/models"
	"github.com/wahyuhadi/rinjani/scan"
)

var (
	location = flag.String("l", ".", "location file to scan , example -l /tmp/folderscan ")
	ext      = flag.String("e", "", "extension  , example -e .go ")
	rules    = flag.String("r", "sast-rules", "example -r /home/rules")
)

func parseOpts() *models.Options {
	flag.Parse()
	// Return Options models
	return &models.Options{
		// Locations file paths
		Location: *location,
		// Extension File
		Ext: *ext,
		// Location Rule Files
		Rules: *rules,
	}
}

func main() {
	// Parse Options and send to struct
	options := parseOpts()

	files, err := directory.GetFileInDir(*options)
	if err != nil {
		gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
		return
	}

	rules, err := directory.GetFileInRuleDir(*options)
	if err != nil {
		gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
		return
	}
	rulesData := directory.RulesToStruct(rules)

	found := scan.RunnerScan(files, rulesData)
	if len(found) == 0 {
		gologger.Info().Msg("Good luck no issue found !!")
		return
	}

	gologger.Info().Msgf("Oops found %v issue", len(found))
	gologger.Info().Msgf("Before push your code, please fixing this issue .")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Location", "index", "Severity", "Remediation"})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	)
	table.SetBorder(true)   // Set Border to false
	table.AppendBulk(found) // Add Bulk Data
	table.Render()
}
