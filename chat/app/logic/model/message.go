package model

const (
	MessageChatTypeUser  = iota + 1 // 用户聊天类型
	MessageChatTypeGroup            // 群组聊天类型
)

const (
	MessageTypeSystem   = iota + 1 //系统消息
	MessageTypeText                //文本
	MessageTypeImage               //图片
	MessageTypeVideo               //视频
	MessageTYpeAudio               //音频
	MessageTypeEmoticon            //表情
	MessageTypeCard                //名片
)

type MessageModel struct {
	PriID
	Uid
	ToId     int    `gorm:"column:to_id;not null;type:int(11) unsigned;index;comment:发送者" json:"to_id"`
	ChatType int8   `gorm:"column:chat_type;not null;default:1;comment:目标类型，1=用户，2=群组" json:"chat_type"`
	Type     int8   `gorm:"column:type;not null;default:1;comment:消息类型" json:"type"`
	Content  string `gorm:"column:content;not null;type:varchar(5000);comment:内容" json:"content"`
	Options  string `gorm:"column:options;type:varchar(255);not null;default:'';comment:选项" json:"options"`
	Create
}

// TableName 表名
func (u *MessageModel) TableName() string {
	return "message"
}
