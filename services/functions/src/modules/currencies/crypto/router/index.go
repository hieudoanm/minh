package crypto_router

import (
	crypto_controller "chatbot-functions/src/modules/currencies/crypto/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/crypto/coins", crypto_controller.GetCryptoCoins)
	router.GET("/crypto/coins/:id", crypto_controller.GetCryptoCoin)
}
