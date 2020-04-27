package main

import (
	"log"
	"os"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
	"github.com/mishrasunny174/Corona-Stats/api"
)

type printableStats struct {
	CountryName string `header:"Country Name"`
	TotalCases  int    `header:"Total Cases"`
	ActiveCases int    `header:"ActiveCases"`
	Deaths      int    `header:"Deaths"`
	Recoverd    int    `header:"Recovered"`
}

func getPrintableStats() ([]printableStats, error) {
	countries, err := api.GetAllCountryStats()
	if err != nil {
		return nil, err
	}
	pStats := make([]printableStats, 0)
	for _, v := range countries {
		var p = printableStats{
			CountryName: v.CountryName,
			TotalCases:  v.Stats.Confirmed,
			Deaths:      v.Stats.Deaths,
			Recoverd:    v.Stats.Recovered,
			ActiveCases: v.Stats.Confirmed - v.Stats.Deaths - v.Stats.Recovered}
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
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor
	printer.Print(printableStats)
}
