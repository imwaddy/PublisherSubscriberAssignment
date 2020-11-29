package models

// IncomingMessage for incoming hotel messages
type IncomingMessage struct {
	Offers []struct {
		CmOfferID    string   `json:"cm_offer_id"`
		Hotel        Hotel    `json:"hotel"`
		Room         Room     `json:"room"`
		RatePlan     RatePlan `json:"rate_plan"`
		OriginalData struct {
			GuaranteePolicy struct {
				Required bool `json:"Required"`
			} `json:"GuaranteePolicy"`
		} `json:"original_data"`
		Capacity capacity
		Number   int    `json:"number"`
		Price    int    `json:"price"`
		Currency string `json:"currency"`
		CheckIn  string `json:"check_in"`
		CheckOut string `json:"check_out"`
		Fees     []struct {
			Type        string  `json:"type"`
			Description string  `json:"description"`
			Included    bool    `json:"included"`
			Percent     float64 `json:"percent"`
		} `json:"fees"`
	} `json:"offers"`
}
