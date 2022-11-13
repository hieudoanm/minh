package forex

import (
	"chatbot-webhook/clients/fixer"
	"chatbot-webhook/utils"
	"fmt"
	"strings"

	"github.com/leekchan/accounting"
)

var API_KEY_FIXER = utils.Getenv("API_KEY_FIXER", "")

var PERSONAL_CURRENCIES = []string{
	"AUD",
	"CAD",
	"CHF",
	"CNH",
	"EUR",
	"GBP",
	"JPY",
	"KRW",
	"SGD",
	"USD",
	"VND",
}

func GetForexMessage() string {
	latestResponse, latestError := fixer.GetLatest(API_KEY_FIXER)
	if latestError != nil {
		return "Forex: N/A"
	}

	var base string = latestResponse.Base
	var rates map[string]float64 = latestResponse.Rates
	if base == "" {
		return "Forex: N/A"
	}

	var personalSymbols = []string{}
	for key := range rates {
		if utils.Contains(PERSONAL_CURRENCIES, key) {
			personalSymbols = append(personalSymbols, key)
		}
	}

	var personalRates = map[string]float64{}
	for _, key := range personalSymbols {
		personalRates[key] = rates[key]
	}

	var ac = accounting.Accounting{Symbol: "VND", Precision: 0, Thousand: ",", Decimal: ""}

	var items = []string{}
	for symbol, symbolRate := range personalRates {
		var baseRate = personalRates[base]
		fmt.Println(symbol, symbolRate)
		var vndRate = personalRates["VND"]
		var rate = baseRate * vndRate / symbolRate
		var item = fmt.Sprintf("- `%s` - `%s`", symbol, ac.FormatMoney(rate))
		items = append(items, item)
	}

	var message string = strings.Join(items, "\n")

	return message
}
