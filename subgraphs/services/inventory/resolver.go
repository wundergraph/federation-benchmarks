package inventory

import (
	"context"
	"fmt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func init() {
	for i := 11; i < 100; i++ {
		inventory = append(inventory, &Product{
			Upc:     fmt.Sprintf("%d", i),
			InStock: boolean(true),
		})
	}
}

type Resolver struct{}

func (r *Resolver) Entity() EntityResolver {
	return &entityResolver{r}
}
func (r *Resolver) Product() ProductResolver {
	return &productResolver{r}
}

type ProductResolver interface {
	ShippingEstimate(ctx context.Context, obj *Product) (*int, error)
}

type entityResolver struct{ *Resolver }

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*Product, error) {
	for _, p := range inventory {
		if p.Upc == upc {
			return p, nil
		}
	}
	return nil, fmt.Errorf("%v not found", upc)
}

var inventory = []*Product{
	{
		Upc:     "1",
		InStock: boolean(true),
	},
	{
		Upc:     "2",
		InStock: boolean(false),
	},
	{
		Upc:     "3",
		InStock: boolean(true),
	},
	{
		Upc:     "4",
		InStock: boolean(false),
	},
	{
		Upc:     "5",
		InStock: boolean(true),
	},
	{
		Upc:     "6",
		InStock: boolean(true),
	},
	{
		Upc:     "7",
		InStock: boolean(false),
	},
	{
		Upc:     "8",
		InStock: boolean(true),
	},
	{
		Upc:     "9",
		InStock: boolean(true),
	},
	{
		Upc:     "10",
		InStock: boolean(true),
	},
}

func boolean(b bool) *bool {
	return &b
}

type productResolver struct{ *Resolver }

func (r *productResolver) ShippingEstimate(ctx context.Context, obj *Product) (*int, error) {
	// free for expensive items
	if obj.Price != nil && *(obj.Price) > 1000 {
		return integer(0), nil
	}
	if obj.Weight == nil {
		return nil, nil
	}
	return integer(*(obj.Weight) / 2), nil
}

func integer(i int) *int {
	return &i
}
