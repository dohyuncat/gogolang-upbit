package upbit

import (
	"fmt"
	"net/http"
)

const (
	FuncGetAccounts                = "GetAccounts"
	FuncGetOrderChance             = "GetOrderChance"
	FuncGetOrder                   = "GetOrder"
	FuncGetOrders                  = "GetOrders"
	FuncPurchaseOrder              = "PurchaseOrder"
	FuncSellOrder                  = "SellOrder"
	FuncCancelOrder                = "CancelOrder"
	FuncGetWithdraws               = "GetWithdraws"
	FuncGetWithdraw                = "GetWithdraw"
	FuncGetWithdrawChance          = "GetWithdrawChance"
	FuncWithdrawCoin               = "WithdrawCoin"
	FuncWithdrawKrw                = "WithdrawKrw"
	FuncGetDeposits                = "GetDeposits"
	FuncGetDeposit                 = "GetDeposit"
	FuncGenerateDepositCoinAddress = "GenerateDepositCoinAddress"
	FuncGetDepositCoinAddresses    = "GetDepositCoinAddresses"
	FuncGetDepositCoinAddress      = "GetDepositCoinAddress"
	FuncDepositKrw                 = "DepositKrw"
	FuncGetWalletStatus            = "GetWalletStatus"
	FuncGetApiKeys                 = "GetApiKeys"
	FuncGetMarkets                 = "GetMarkets"
	FuncGetMinuteCandles           = "GetMinuteCandles"
	FuncGetDayCandles              = "GetDayCandles"
	FuncGetWeekCandles             = "GetWeekCandles"
	FuncGetMonthCandles            = "GetMonthCandles"
	FuncGetTradeTicks              = "GetTradeTicks"
	FuncGetTickers                 = "GetTickers"
	FuncGetOrderbooks              = "GetOrderbooks"

	RouteAccounts                    = "/accounts"
	RouteOrderChance                 = "/orders/chance"
	RouteOrder                       = "/order"
	RouteOrders                      = "/orders"
	RouteWithdraws                   = "/withdraws"
	RouteWithdraw                    = "/withdraw"
	RouteWithdrawsChance             = "/withdraws/chance"
	RouteWithdrawsCoin               = "/withdraws/coin"
	RouteWithdrawsKrw                = "/withdraws/krw"
	RouteDeposits                    = "/deposits"
	RouteDeposit                     = "/deposit"
	RouteDepositsGenerateCoinAddress = "/deposits/generate_coin_address"
	RouteDepositsCoinAddresses       = "/deposits/coin_addresses"
	RouteDepositsCoinAddress         = "/deposits/coin_address"
	RouteDepositsKrw                 = "/deposits/krw"
	RouteStatusWallet                = "/status/wallet"
	RouteApiKeys                     = "/api_keys"
	RouteMarketAll                   = "/market/all"
	RouteCandlesMinutes              = "/candles/minutes/"
	RouteCandlesDays                 = "/candles/days"
	RouteCandlesWeeks                = "/candles/weeks"
	RouteCandlesMonths               = "/candles/months"
	RouteTradesTicks                 = "/trades/ticks"
	RouteTicker                      = "/ticker"
	RouteOrderbook                   = "/orderbook"

	ApiSectionExchange  = "exchange"
	ApiSectionQuotation = "quotation"

	ApiGroupDefault      = "default"
	ApiGroupOrder        = "order"
	ApiGroupStatusWallet = "status-wallet"
	ApiGroupMarket       = "market"
	ApiGroupCandles      = "candles"
	ApiGroupCrixTrades   = "crix-trades"
	ApiGroupTicker       = "ticker"
	ApiGroupOrderbook    = "orderbook"
)

type ApiInfo struct {
	Method, Url, Section, Group string
}

func GetApiInfo(funcName string) (*ApiInfo, error) {
	switch funcName {
	case FuncGetAccounts: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteAccounts,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetOrderChance: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrderChance,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetOrder: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrder,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetOrders: //?????? ????????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrders,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncPurchaseOrder: //????????????
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteOrders,
			Section: ApiSectionExchange, Group: ApiGroupOrder,
		}, nil
	case FuncSellOrder: //????????????
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteOrders,
			Section: ApiSectionExchange, Group: ApiGroupOrder,
		}, nil
	case FuncCancelOrder: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodDelete, Url: RouteOrder,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWithdraws: //?????? ????????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteWithdraws,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWithdraw: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteWithdraw,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWithdrawChance: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteWithdrawsChance,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncWithdrawCoin: //?????? ????????????
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteWithdrawsCoin,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncWithdrawKrw: //?????? ????????????
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteWithdrawsKrw,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDeposits: //?????? ????????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDeposits,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDeposit: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDeposit,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGenerateDepositCoinAddress: //?????? ?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteDepositsGenerateCoinAddress,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDepositCoinAddresses: //?????? ?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDepositsCoinAddresses,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDepositCoinAddress: //?????? ?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDepositsCoinAddress,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncDepositKrw: //?????? ????????????
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteDepositsKrw,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWalletStatus: //????????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteStatusWallet,
			Section: ApiSectionExchange, Group: ApiGroupStatusWallet,
		}, nil
	case FuncGetApiKeys: //API ??? ????????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteApiKeys,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetMarkets: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteMarketAll,
			Section: ApiSectionQuotation, Group: ApiGroupMarket,
		}, nil
	case FuncGetMinuteCandles: //??? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesMinutes,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetDayCandles: //??? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesDays,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetWeekCandles: //??? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesWeeks,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetMonthCandles: //??? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesMonths,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetTradeTicks: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteTradesTicks,
			Section: ApiSectionQuotation, Group: ApiGroupCrixTrades,
		}, nil
	case FuncGetTickers: //????????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteTicker,
			Section: ApiSectionQuotation, Group: ApiGroupTicker,
		}, nil
	case FuncGetOrderbooks: //?????? ?????? ??????
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrderbook,
			Section: ApiSectionQuotation, Group: ApiGroupOrderbook,
		}, nil
	default:
		return nil, fmt.Errorf("function is not defined")
	}
}
