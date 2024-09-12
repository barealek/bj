package client

import (
	t "github.com/barealek/bj/types"
)

type BJModel struct {
	player *t.Player
	dealer *t.Dealer
}

func CreateBJModel() *BJModel {
	bjModel := &BJModel{
		player: t.CreatePlayer(),
		dealer: t.CreateDealer(),
	}
	bjModel.dealer.DealSelf()
	bjModel.player.Hit(bjModel.dealer.Deal())
	bjModel.dealer.DealSelf()
	bjModel.player.Hit(bjModel.dealer.Deal())

	return bjModel
}
