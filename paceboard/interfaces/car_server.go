package interfaces

import (
	"backend/paceboard"
	"time"
)

type CarServer interface {
	Open(port string, ttl time.Duration) error
	Close() error

	Accept() (*paceboard.CarConnection, error)
	Read(connection *paceboard.CarConnection) (paceboard.CarPacket, error)
	Write(connection *paceboard.CarConnection, packet paceboard.CarPacket) error
}
