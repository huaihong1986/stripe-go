package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestPrice_Unmarshal(t *testing.T) {
	priceData := map[string]interface{}{
		"id":                  "price_123",
		"object":              "price",
		"unit_amount":         0,
		"unit_amount_decimal": "0.0123456789",
	}

	bytes, err := json.Marshal(&priceData)
	assert.NoError(t, err)

	var price Price
	err = json.Unmarshal(bytes, &price)
	assert.NoError(t, err)

	assert.Equal(t, 0.0123456789, price.UnitAmountDecimal)
}

func TestPriceTierParams_AppendTo(t *testing.T) {
	params := &PriceParams{
		Tiers: []*PriceTierParams{
			{UnitAmount: Int64(500), UpTo: Int64(5)},
			{UnitAmount: Int64(100), UpToInf: Bool(true)},
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	t.Logf("body = %+v", body)
	assert.Equal(t, []string{"inf"}, body)
}
