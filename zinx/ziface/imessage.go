package ziface

type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetDataLen(uint32)
	SetData([]byte)
}
