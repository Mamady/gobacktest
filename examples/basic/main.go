package main

import (
	"github.com/mamady/gobacktest/pkg/backtest"
	"github.com/mamady/gobacktest/pkg/data"
	"github.com/mamady/gobacktest/pkg/strategy"
)

func main() {
	// define symbols
	symbols := []string{"TEST.DE"}

	// initiate new backtester and load symbols
	test := backtest.New()
	test.SetSymbols(symbols)

	// create data provider and load data into the backtest
	data := &data.BarEventFromCSVFile{FileDir: "../testdata/test/"}
	data.Load(symbols)
	test.SetData(data)

	// set portfolio with initial cash and default size and risk manager
	portfolio := &backtest.Portfolio{}
	portfolio.SetInitialCash(10000)

	sizeManager := &backtest.Size{DefaultSize: 100, DefaultValue: 1000}
	portfolio.SetSizeManager(sizeManager)

	riskManager := &backtest.Risk{}
	portfolio.SetRiskManager(riskManager)

	test.SetPortfolio(portfolio)

	// create strategy provider and load into the backtest
	strategy := &strategy.Basic{}
	test.SetStrategy(strategy)

	// create execution provider and load into the backtest
	exchange := &backtest.Exchange{Symbol: "TEST", ExchangeFee: 1.00}
	test.SetExchange(exchange)

	// choose a statistic and load into the backtest
	statistic := &backtest.Statistic{}
	test.SetStatistic(statistic)

	// run the backtest
	test.Run()

	// print the result of the test
	test.Stats().PrintResult()
}
