package graph

import (
	"context"
	"testing"

	"starbucks/graph/model"
)

func TestResolver_Coffees(t *testing.T) {
	resolver := NewResolver()
	
	// Add test data
	coffee1 := &model.Coffee{
		ID:                "1",
		Name:              "Test Coffee",
		Size:              "Grande",
		Price:             4.50,
		QuantityAvailable: 10,
		Description:       strPtr("Test description"),
	}
	resolver.AddCoffee(coffee1)

	// Test query
	ctx := context.Background()
	queryResolver := resolver.Query()
	
	coffees, err := queryResolver.Coffees(ctx)
	if err != nil {
		t.Fatalf("Failed to get coffees: %v", err)
	}

	if len(coffees) != 1 {
		t.Errorf("Expected 1 coffee, got %d", len(coffees))
	}

	if coffees[0].Name != "Test Coffee" {
		t.Errorf("Expected coffee name 'Test Coffee', got '%s'", coffees[0].Name)
	}
}

func TestResolver_Coffee(t *testing.T) {
	resolver := NewResolver()
	
	coffee1 := &model.Coffee{
		ID:                "1",
		Name:              "Test Coffee",
		Size:              "Grande",
		Price:             4.50,
		QuantityAvailable: 10,
	}
	resolver.AddCoffee(coffee1)

	ctx := context.Background()
	queryResolver := resolver.Query()
	
	// Test getting existing coffee
	coffee, err := queryResolver.Coffee(ctx, "1")
	if err != nil {
		t.Fatalf("Failed to get coffee: %v", err)
	}

	if coffee.Name != "Test Coffee" {
		t.Errorf("Expected coffee name 'Test Coffee', got '%s'", coffee.Name)
	}

	// Test getting non-existent coffee
	_, err = queryResolver.Coffee(ctx, "999")
	if err == nil {
		t.Error("Expected error for non-existent coffee, got nil")
	}
}

func TestResolver_PurchaseCoffee(t *testing.T) {
	resolver := NewResolver()
	
	coffee1 := &model.Coffee{
		ID:                "1",
		Name:              "Test Coffee",
		Size:              "Grande",
		Price:             4.50,
		QuantityAvailable: 10,
	}
	resolver.AddCoffee(coffee1)

	ctx := context.Background()
	mutationResolver := resolver.Mutation()

	// Test successful purchase
	result, err := mutationResolver.PurchaseCoffee(ctx, "1", 2)
	if err != nil {
		t.Fatalf("Failed to purchase coffee: %v", err)
	}

	if !result.Success {
		t.Error("Expected purchase to succeed")
	}

	if result.RemainingQuantity != 8 {
		t.Errorf("Expected remaining quantity 8, got %d", result.RemainingQuantity)
	}

	// Test purchasing more than available
	result, err = mutationResolver.PurchaseCoffee(ctx, "1", 20)
	if err != nil {
		t.Fatalf("Failed to attempt purchase: %v", err)
	}

	if result.Success {
		t.Error("Expected purchase to fail due to insufficient stock")
	}

	// Test purchasing non-existent coffee
	result, err = mutationResolver.PurchaseCoffee(ctx, "999", 1)
	if err != nil {
		t.Fatalf("Failed to attempt purchase: %v", err)
	}

	if result.Success {
		t.Error("Expected purchase to fail for non-existent coffee")
	}
}

func TestResolver_ConcurrentPurchases(t *testing.T) {
	resolver := NewResolver()
	
	coffee1 := &model.Coffee{
		ID:                "1",
		Name:              "Test Coffee",
		Size:              "Grande",
		Price:             4.50,
		QuantityAvailable: 100,
	}
	resolver.AddCoffee(coffee1)

	ctx := context.Background()
	mutationResolver := resolver.Mutation()

	// Simulate concurrent purchases
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			_, err := mutationResolver.PurchaseCoffee(ctx, "1", 5)
			if err != nil {
				t.Errorf("Failed to purchase coffee: %v", err)
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Check final quantity
	queryResolver := resolver.Query()
	coffee, err := queryResolver.Coffee(ctx, "1")
	if err != nil {
		t.Fatalf("Failed to get coffee: %v", err)
	}

	expectedQuantity := 50 // 100 - (10 * 5)
	if coffee.QuantityAvailable != expectedQuantity {
		t.Errorf("Expected quantity %d, got %d", expectedQuantity, coffee.QuantityAvailable)
	}
}

func strPtr(s string) *string {
	return &s
}
