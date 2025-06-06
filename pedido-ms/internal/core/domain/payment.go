package domain

import (
	"pedido-ms/internal/core/enum"
)

type (
	PaymentParams struct {
		CardId, Bin, NumToken, CardholderName, SecurityCode, ExpirationMonth, ExpirationYear string
		Brand                                                                                enum.Brand
	}

	Payment struct {
		cardId, bin, numToken, cardholderName, securityCode, expirationMonth, expirationYear string
		brand                                                                                enum.Brand
	}
)

func CreatePayment(input PaymentParams) *Payment {
	payment := Payment{
		cardId:          input.CardId,
		bin:             input.Bin,
		numToken:        input.NumToken,
		cardholderName:  input.CardholderName,
		securityCode:    input.SecurityCode,
		expirationMonth: input.ExpirationMonth,
		expirationYear:  input.ExpirationYear,
		brand:           input.Brand,
	}

	return &payment
}

func RecoverPayment(input PaymentParams) *Payment {
	payment := Payment{
		brand:           input.Brand,
		cardId:          input.CardId,
		bin:             input.Bin,
		numToken:        input.NumToken,
		cardholderName:  input.CardholderName,
		securityCode:    input.SecurityCode,
		expirationMonth: input.ExpirationMonth,
		expirationYear:  input.ExpirationYear,
	}

	return &payment
}

func (p *Payment) Map() *PaymentParams {
	pp := PaymentParams{}
	pp.CardId = p.cardId
	pp.Bin = p.bin
	pp.NumToken = p.numToken
	pp.CardholderName = p.cardholderName
	pp.SecurityCode = p.securityCode
	pp.ExpirationMonth = p.expirationMonth
	pp.ExpirationYear = p.expirationYear
	pp.Brand = p.brand

	return &pp
}
