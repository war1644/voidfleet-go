package game

import "fmt"

type Event struct {
	MsgList    []Msg
	MsgTypeMap map[string]string
}

func NewEvent() *Event {
	return &Event{
		MsgList:    []Msg{},
		MsgTypeMap: map[string]string{"info": "信息", "primary": "新闻", "success": "任务"},
	}
}

type Msg struct {
	MsgType     string
	MsgTypePill string
	MsgText     string
}

func (s *Event) Add(msg Msg) {
	s.MsgList = append(s.MsgList, msg)
}

func (s *Event) Get(msgLen int) []Msg {
	if len(s.MsgList) < msgLen {
		fmt.Println("MsgList not Length!")
	}
	tmpMsgList := s.MsgList[:msgLen]
	s.MsgList = s.MsgList[msgLen:]
	return tmpMsgList
}

func (s *Event) NewMsg(msgType, msgText string) {
	newMsg := Msg{
		MsgType:     msgType,
		MsgTypePill: s.MsgTypeMap[msgType],
		MsgText:     msgText,
	}
	s.Add(newMsg)
	//{MsgType:"info",MsgTypePill:"信息",Msg:`你已进入${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`},
	//{MsgType:"primary",MsgTypePill:"新闻",Msg:`海盗正在袭击${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`},
	//{MsgType:"primary",MsgTypePill:"新闻",Msg:`星区警卫队正在前往${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`},
	//{MsgType:"success",MsgTypePill:"任务",Msg:`test任务,坐标:0,0`}
}
