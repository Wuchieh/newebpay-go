package ebpay

import (
	"errors"
	"fmt"
)

type ItemType int

const (
	ItemTypeProduct ItemType = iota + 1
	ItemTypeTickets
	ItemTypeReserve
)

type OrderDetail struct {
	ItemName    string   `validate:"required"`
	ItemAmt     int      `validate:"required"`
	ItemType    ItemType `validate:"required,item_type"`
	ItemOrderNo string   `validate:"required"`
}

func (o *OrderDetail) validate() error {
	err := validate.Struct(o)
	if err != nil {
		if vErr, ok := isValidatorErr(err); ok {
			var errs []error
			for _, fieldError := range vErr {
				switch fieldError.Field() {
				case "ItemName":
					switch fieldError.Tag() {
					case "required":
						errs = append(errs, ErrOrderDetailItemNameEmpty)
					}
				case "ItemAmt":
					switch fieldError.Tag() {
					case "required":
						errs = append(errs, ErrOrderDetailItemAmtEmpty)
					}
				case "ItemType":
					switch fieldError.Tag() {
					case "required":
						errs = append(errs, ErrOrderDetailItemTypeEmpty)
					case "item_type":
						errs = append(errs, ErrOrderDetailItemOrderTypeExist)
					}
				case "ItemOrderNo":
					switch fieldError.Tag() {
					case "required":
						errs = append(errs, ErrOrderDetailItemOrderNoEmpty)
					}
				}
			}

			fmt.Println(errs)

			if errs == nil {
				return err
			} else {
				return errors.Join(errs...)
			}
		}
	}
	return nil
}
