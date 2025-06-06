package enum

import "fmt"

type (
	OrderState int
	Brand      int
)

const (
	Recebido OrderState = iota
	Cancelado
	Confirmado
	Em_preparacao
	Enviado
	Entregue
)

var stateName = map[OrderState]string{
	Recebido:      "recebido",
	Cancelado:     "cancelado",
	Confirmado:    "confirmado",
	Em_preparacao: "em preparacao",
	Enviado:       "enviado",
	Entregue:      "entrege",
}

func (os OrderState) String() string {
	return stateName[os]
}

const (
	Mastercard Brand = iota
	Visa
	Amex
	Elo
	Hipercard
)

var BrandTypeName = map[Brand]string{
	Mastercard: "Mastercard",
	Visa:       "Visa",
	Amex:       "Amex",
	Elo:        "Elo",
	Hipercard:  "Hipercard",
}

func (pt Brand) String() string {
	return BrandTypeName[pt]
}

func GetBrandBy(value string) Brand {
	switch value {
	case "Mastercard":
		return Mastercard
	case "Visa":
		return Visa
	case "Amex":
		return Amex
	case "Elo":
		return Elo
	case "Hipercard":
		return Hipercard
	default:
		panic(fmt.Errorf("unknown brand: %s", value))
	}
}
