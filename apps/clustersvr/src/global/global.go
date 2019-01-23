package global

import (
	"github.com/SpiritDyn123/gocygame/apps/common/tools"
	"github.com/SpiritDyn123/gocygame/libs/utils"
	"github.com/SpiritDyn123/gocygame/libs/timer"
)

type IClusterSvrGlobal interface {
	utils.IPooller
	GetMsgDispatcher() tools.IMsgDispatcher
	GetWheelTimer() timer.WheelTimer
	GetSvrsMgr() ISvrsMgr
}

var ClusterSvrGlobal IClusterSvrGlobal