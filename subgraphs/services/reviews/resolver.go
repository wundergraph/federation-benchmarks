package reviews

import (
	"context"
	"fmt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func init() {
	for i := 12; i < 100; i++ {
		reviews = append(reviews, &Review{
			ID:      fmt.Sprintf("%d", i),
			Author:  &User{ID: "3", Username: str("@shopper")},
			Product: &Product{Upc: fmt.Sprintf("%d", i)},
			Body:    str("It's ok."),
		})
	}
}

type Resolver struct{}

func (r *Resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

type entityResolver struct{ *Resolver }

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*User, error) {
	u := &User{ID: id}
	for _, user := range usernames {
		if id == user.ID {
			u.Username = user.Username
		}
	}
	for _, r := range reviews {
		if r.Author.ID == id {
			u.Reviews = append(u.Reviews, r)
		}
	}
	return u, nil
}

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*Product, error) {
	p := &Product{Upc: upc}
	for _, r := range reviews {
		if r.Product.Upc == upc {
			p.Reviews = append(p.Reviews, r)
		}
	}
	return p, nil
}

func (r *entityResolver) FindReviewByID(ctx context.Context, id string) (*Review, error) {
	for _, r := range reviews {
		if r.ID == id {
			return r, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

var reviews = []*Review{
	{
		ID:      "1",
		Body:    str("Love it!"),
		Author:  &User{ID: "1", Username: str("@ada")},
		Product: &Product{Upc: "1"},
	},
	{
		ID:      "2",
		Author:  &User{ID: "1", Username: str("@ada")},
		Product: &Product{Upc: "2"},
		Body:    str("Too expensive."),
	},
	{
		ID:      "3",
		Author:  &User{ID: "2", Username: str("@complete")},
		Product: &Product{Upc: "3"},
		Body:    str("Could be better."),
	},
	{
		ID:      "4",
		Author:  &User{ID: "2", Username: str("@complete")},
		Product: &Product{Upc: "1"},
		Body:    str("Prefer something else."),
	},
	{
		ID:      "5",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "4"},
		Body:    str("Awesome!"),
	},
	{
		ID:      "6",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "5"},
		Body:    str("Could be better."),
	},
	{
		ID:      "7",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "6"},
		Body:    str("Ok."),
	},
	{
		ID:      "8",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "7"},
		Body:    str("Liked it!"),
	},
	{
		ID:      "9",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "8"},
		Body:    str("Not for me."),
	},
	{
		ID:      "10",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "9"},
		Body:    str("Liked it!"),
	},
	{
		ID:      "11",
		Author:  &User{ID: "3", Username: str("@shopper")},
		Product: &Product{Upc: "10"},
		Body:    str("A bit pricey."),
	},
}

var usernames = []*User{
	{ID: "1", Username: str("@ada")},
	{ID: "2", Username: str("@complete")},
	{ID: "3", Username: str("@shopper")},
}

func str(s string) *string {
	return &s
}
