package requests

import "final_project/business/transactions"

type Payment struct {
	// ID          int     `json:"id"`
	Total_Price float64 `json:"total_price"`
}

func (payment *Payment) ToDomain() transactions.TransactionDomain {
	return transactions.TransactionDomain{
		// ID:          payment.ID,
		Total_Price: payment.Total_Price,
	}
}
