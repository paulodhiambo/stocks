package dao

type Stock struct {
	TradeId       int     `json:"trade_id"`
	StockSymbol   string  `json:"stock_symbol"`
	TradeDate     string  `json:"trade_date"`
	TradeTime     string  `json:"trade_time"`
	TradePrice    float64 `json:"trade_price"`
	TradeVolume   int     `json:"trade_volume"`
	BuyerId       int     `json:"buyer_id"`
	SellerId      int     `json:"seller_id"`
	BuyerName     string  `json:"buyer_name"`
	SellerName    string  `json:"seller_name"`
	BuyerCountry  string  `json:"buyer_country"`
	SellerCountry string  `json:"seller_country"`
	BuyerEmail    string  `json:"buyer_email"`
	SellerEmail   string  `json:"seller_email"`
	BuyerPhone    string  `json:"buyer_phone"`
	SellerPhone   string  `json:"seller_phone"`
	BuyerAddress  string  `json:"buyer_address"`
	SellerAddress string  `json:"seller_address"`
	BuyerCity     string  `json:"buyer_city"`
	SellerCity    string  `json:"seller_city"`
}
