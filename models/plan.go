package data

import (
	"context"
	"fmt"
	"time"
)

// Plan is the type for subscription plans
type Plan struct {
	ID         int    `json:"id"`
	PlanName   string `json:"plan_name"`
	PlanAmount int    `json:"plan_amount"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p *Plan) GetAll() ([]*Plan, error) {
	ctx := context.Background()

	var plans []*Plan

	query := db.DB.From(planTable).Select("*")

	if err := query.Execute(ctx, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

// GetOne returns one plan by id
func (p *Plan) GetOne(id string) (*Plan, error) {
	ctx := context.Background()

	var plan Plan

	query := db.DB.From(planTable).Select("*").Eq("id", id)
	if err := query.Execute(ctx, &plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

// SubscribeUserToPlan subscribes a user to one plan by insert
// values into user_plans table
func (p *Plan) SubscribeUserToPlan(user User, plan Plan) error {

	// _, _, err := db.From(userPlanTable).Insert(map[string]interface{}{

	// }).Execute()
	// if err != nil {
	// 	return err
	// }
	return nil

}

// AmountForDisplay formats the price we have in the DB as a currency string
func (p *Plan) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
