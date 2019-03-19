package game

type Event struct {
	MsgList []Msg
}

type Msg struct {
	MsgType     string
	MsgTypePill string
	MsgText     string
}

func (s *Event) Add(msg Msg) {
	s.MsgList = append(s.MsgList, msg)
}

func NewEvent() *Event {
	return &Event{
		MsgList: []Msg{},
	}
}

func (s *Event) NewMsg(msgType, msgTypePill, msgText string) Msg {
	return Msg{
		MsgType:     "info",
		MsgTypePill: "信息",
		MsgText:     `你已进入${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`,
	}
	//{MsgType:"info",MsgTypePill:"信息",Msg:`你已进入${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`},
	//{MsgType:"primary",MsgTypePill:"新闻",Msg:`海盗正在袭击${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`},
	//{MsgType:"primary",MsgTypePill:"新闻",Msg:`星区警卫队正在前往${data.Galaxy}-${data.Name},坐标:${data.X},${data.Y}`},
	//{MsgType:"success",MsgTypePill:"任务",Msg:`test任务,坐标:0,0`}
}
