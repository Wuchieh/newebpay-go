package ebpay

import "errors"

var (
	ErrOrderDetailItemNameEmpty      = errors.New("ItemName is empty")
	ErrOrderDetailItemAmtEmpty       = errors.New("ItemAmt is empty")
	ErrOrderDetailItemTypeEmpty      = errors.New("ItemType is empty")
	ErrOrderDetailItemOrderNoEmpty   = errors.New("ItemOrderNo is empty")
	ErrOrderDetailItemOrderTypeExist = errors.New("ItemType is not exist")

	ErrMPGAesEncryption = errors.New("aes encryption error")
)
