package models

import "testing"

func TestStockModel(t *testing.T) {
	stock := Stock{
		Code:  "AAPL",
		Name:  "Apple Inc.",
		Price: 150.0,
	}

	if stock.Code != "AAPL" {
		t.Errorf("Expected AAPL, got %s", stock.Code)
	}
	if stock.Name != "Apple Inc." {
		t.Errorf("Expected Apple Inc., got %s", stock.Name)
	}
	if stock.Price != 150.0 {
		t.Errorf("Expected 150.0, got %f", stock.Price)
	}
}
