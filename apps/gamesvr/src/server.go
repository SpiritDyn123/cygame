package src

import (
	"github.com/SpiritDyn123/gocygame/libs/net/tcp"
	"github.com/SpiritDyn123/gocygame/libs/utils"
	"github.com/SpiritDyn123/gocygame/libs/timer"
	"github.com/SpiritDyn123/gocygame/libs/chanrpc"
	"github.com/SpiritDyn123/gocygame/libs/go"
	"fmt"
	"github.com/SpiritDyn123/gocygame/apps/gamesvr/src/etc"
	"github.com/SpiritDyn123/gocygame/apps/common"
	"github.com/SpiritDyn123/gocygame/apps/common/net/codec"
	"github.com/SpiritDyn123/gocygame/apps/common/net"
	"github.com/SpiritDyn123/gocygame/apps/gamesvr/src/global"
	"github.com/SpiritDyn123/gocygame/apps/common/tools"
	"github.com/SpiritDyn123/gocygame/apps/gamesvr/src/session"
)


type GameSvrGlobal struct {
	utils.Pooller
	net_ser tcp.INetServer

	wheel_timer_ timer.WheelTimer

	msg_dispatcher_ tools.IMsgDispatcher


	session_mgr_ *session.SessionMgr
}

func (svr *GameSvrGlobal) GetName() string {
	return fmt.Sprintf("%s_%d_%d", etc.Game_Config.System_.Svr_name_,
		etc.Game_Config.System_.Svr_group_id_, etc.Game_Config.System_.Svr_id_)
}

func (svr *GameSvrGlobal) Start() (err error) {

	global.GameSvrGlobal = svr

	svr.ChanServer = chanrpc.NewServer(common.Default_Chan_Server_Len)
	svr.ChanServer.Register(common.Chanrpc_key_tcp_accept, svr.onTcpAccept)
	svr.ChanServer.Register(common.Chanrpc_key_tcp_recv, svr.onTcpRecv)
	svr.ChanServer.Register(common.Chanrpc_key_tcp_close, svr.onTcpClose)

	svr.GoServer = g.New(common.Default_Go_Server_Len)

	//定时器
	svr.wheel_timer_ = timer.CreateWheelTimer()
	svr.TimerServer = timer.NewDispatcher(10)
	svr.TimerServer.AfterFunc(common.Default_Svr_Logic_time, svr.onTimer)

	//消息管理器
	svr.msg_dispatcher_ = tools.CreateMsgDispatcher()

	//session管理器
	svr.session_mgr_ = &session.SessionMgr{}

	//启动socket
	protocol := &codec.ProtoInnerProtocol{
		Endian_: common.Default_Net_Endian,
	}

	svr.net_ser, err = net.CreateTcpServer(&etc.Game_Config.System_, nil, protocol, common.Default_Net_Head_Len,
		common.Default_Net_Endian, common.Default_Send_Chan_Len, svr.ChanServer)
	if err != nil {
		return
	}
	if !svr.net_ser.Start() {
		return fmt.Errorf("netserver start error")
	}

	return nil
}

func (svr *GameSvrGlobal) Close() {
	svr.net_ser.Stop()
}

func (svr *GameSvrGlobal) Pool(cs chan bool) {
	svr.Pooller.Pool(cs)
}

func (svr *GameSvrGlobal) GetPriority() int {
	return 0
}

func (svr *GameSvrGlobal) GetMsgDispatcher() tools.IMsgDispatcher {
	return svr.msg_dispatcher_
}

func (svr *GameSvrGlobal) GetWheelTimer() timer.WheelTimer {
	return svr.wheel_timer_
}

func (svr *GameSvrGlobal) onTimer() {
	svr.wheel_timer_.Step()
	svr.TimerServer.AfterFunc(common.Default_Svr_Logic_time, svr.onTimer)
}

func (svr *GameSvrGlobal) onTcpAccept(args []interface{}) {
	tcp_session := args[0].(*tcp.Session)
	svr.session_mgr_.OnAccept(tcp_session)

}

func (svr *GameSvrGlobal) onTcpRecv(args []interface{}) {
	tcp_session := args[0].(*tcp.Session)
	svr.session_mgr_.OnRecv(tcp_session, args[1])
}

func (svr *GameSvrGlobal) onTcpClose(args []interface{}) {
	tcp_session := args[0].(*tcp.Session)
	svr.session_mgr_.OnClose(tcp_session)
}