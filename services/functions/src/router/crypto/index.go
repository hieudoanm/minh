package crypto_router

import (
	crypto_controller "chatbot-functions/src/controller/crypto"

	"github.com/julienschmidt/httprouter"
)

func CryptoRouter(router *httprouter.Router) {
	router.GET("/crypto/coins", crypto_controller.GetCryptoCoins)
	router.GET("/crypto/coins/:id", crypto_controller.GetCryptoCoin)
}
