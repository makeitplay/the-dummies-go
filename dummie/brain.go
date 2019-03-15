package dummie

import (
	"fmt"
	"github.com/makeitplay/arena"
	"github.com/makeitplay/arena/orders"
	"github.com/makeitplay/client-player-go"
	"github.com/makeitplay/the-dummies-go/strategy"
	"github.com/sirupsen/logrus"
)

var ClientResponser client.Responser
var TeamPlace arena.TeamPlace
var PlayerNumber arena.PlayerNumber
var MyRule strategy.PlayerRule

// TeamBallPossession stores the team's name that has touched on the ball for the last time
var TeamBallPossession arena.TeamPlace

type Dummie struct {
	GameMsg     *client.GameMessage
	Player      *client.Player
	TeamState   strategy.TeamState
	PlayerState strategy.PlayerState
	Logger      *logrus.Entry
}

func (d *Dummie) React() {
	var ordersSet []orders.Order
	var msg string
	//if d.Player.IsGoalkeeper() {
	//	msg, ordersSet = d.orderForGoalkeeper()
	//} else {
	switch d.PlayerState {
	case strategy.DisputingTheBall:
		msg, ordersSet = d.orderForDisputingTheBall()
		ordersSet = append(ordersSet, d.Player.CreateCatchOrder())
	case strategy.Supporting:
		msg, ordersSet = d.orderForSupporting()
	case strategy.HoldingTheBall:
		msg, ordersSet = d.orderForHoldingTheBall()
		//case strategy.Defending:
		//	msg, ordersSet = b.orderForDefending(turn)
		//	ordersSet = append(ordersSet, turn.Player().CreateCatchOrder())
		//}
	}

	ClientResponser.SendOrders(fmt.Sprintf("%s %s", d.Player.ID(), msg), ordersSet...)
}