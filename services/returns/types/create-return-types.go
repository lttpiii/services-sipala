package types

type (
	DTOCreateReturns struct {
		BorrowTransactionID string `json:"borrow_transaction_id" binding:"required"`
	}
)

type (
	ReqCreateReturns struct {
		BorrowTransactionID string
	}

	ResCreateReturns Return
)
