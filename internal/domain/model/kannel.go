package model

type (
	KannelRequest struct {
		From      string `json:"from" form:"from" query:"from" validate:"required" `
		To        string `json:"to" form:"to" query:"to" validate:"required"`
		Content   string `json:"content" form:"content" query:"content" validate:"required"`
		IpAddress string `json:"ip_address" form:"ip_address" query:"ip_address"`
		TrxId     string `json:"trx_id" form:"trx_id" query:"trx_id"`
	}

	KannelResponse struct {
		RequestId    string `json:"requestId"`
		ErrorCode    string `json:"errorCode"`
		ErrorMessage string `json:"errorMessage"`
		Message      string `json:"Message"`
	}
)
