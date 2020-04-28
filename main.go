package main

import (
	"log"
	"os"

	"github.com/landoop/tableprinter"
	"github.com/mishrasunny174/Corona-Stats/api"
	"github.com/olekukonko/tablewriter"
)

type printableStats struct {
	Rank        int    `header:"Rank"`
	CountryName string `header:"Country Name"`
	TotalCases  int    `header:"Total Cases"`
	ActiveCases int    `header:"Active Cases"`
	Deaths      int    `header:"Deaths"`
	Recoverd    int    `header:"Recovered"`
	TotalTests  int    `header:"Total Tests"`
}

func getPrintableStats() ([]printableStats, error) {
	countries, err := api.GetAllCountryStats()
	if err != nil {
		return nil, err
	}
	pStats := make([]printableStats, 0)
	for i, v := range countries {
		var p = printableStats{
			Rank:        i + 1,
			CountryName: v.CountryName,
			TotalCases:  v.TotalCases,
			Deaths:      v.Deaths,
			Recoverd:    v.Recovered,
			ActiveCases: v.Active,
			TotalTests:  v.Tests}
		pStats = append(pStats, p)
	}
	return pStats, nil
}

func main() {
	printableStats, err := getPrintableStats()
	printer := tableprinter.New(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderAlignment = tablewriter.ALIGN_CENTER
	printer.HeaderFgColor = tablewriter.FgGreenColor
	printer.DefaultAlignment = tablewriter.ALIGN_CENTER
	printer.Print(printableStats)
}
