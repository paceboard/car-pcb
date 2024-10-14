package paceboard

type CarPacket struct {
	Seq           uint32
	MessageType   uint8
	PayloadLength uint16
	Payload       []byte
}
