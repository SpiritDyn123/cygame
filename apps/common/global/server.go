package global

import (
	"github.com/SpiritDyn123/gocygame/apps/common/proto"
	"github.com/SpiritDyn123/gocygame/apps/common/tools"
	"github.com/SpiritDyn123/gocygame/libs/timer"
	"github.com/SpiritDyn123/gocygame/libs/utils"
)

type IServerGlobal interface {
	utils.IPooller
	GetMsgDispatcher() tools.IMsgDispatcher
	GetWheelTimer() timer.WheelTimer
}

type IServerGlobal_Publish interface {
	GetPublishSvrs() []ProtoMsg.EmSvrType
	GetSvrBaseInfo() *ProtoMsg.PbSvrBaseInfo
}