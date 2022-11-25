package crypto_controller

import (
	crypto_client "chatbot-functions/src/clients/crypto"
	"chatbot-functions/src/utils"
	"fmt"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetCryptoCoins(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// Query Parameters
	limit, limitError := strconv.Atoi(utils.GetQueryParameter(request, "limit", "100"))
	if limitError != nil {
		fmt.Println(limitError)
	}
	offset, offsetError := strconv.Atoi(utils.GetQueryParameter(request, "offset", "100"))
	if offsetError != nil {
		fmt.Println(offsetError)
	}
	var orderBy = utils.GetQueryParameter(request, "orderBy", crypto_client.ORDER_BY["MARKET_CAP"])
	var orderDirection = utils.GetQueryParameter(request, "orderDirection", crypto_client.ORDER_DIRECTION["DESC"])
	var timePeriod = utils.GetQueryParameter(request, "timePeriod", crypto_client.TIME_PERIOD["24h"])
	var tier = utils.GetQueryParameter(request, "tier", "")
	var tag = utils.GetQueryParameter(request, "tag", "")
	// Service
	var coinsResponseBody crypto_client.CoinsResponseBody = crypto_client.GetCryptoCoins(crypto_client.CoinsRequest{
		Limit:          limit,
		Offset:         offset,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
		TimePeriod:     timePeriod,
		Tier:           tier,
		Tag:            tag,
	})

	json.NewEncoder(writer).Encode(coinsResponseBody)
}

func GetCryptoCoin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var id string = params.ByName("id")
	// Service
	var coinResponseBody crypto_client.CoinResponseBody = crypto_client.GetCryptoCoin(id)

	json.NewEncoder(writer).Encode(coinResponseBody)
}
