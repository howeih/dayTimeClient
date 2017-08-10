package daytimeClient

import (
	"fmt"
	"net"
)

const (
	tcp string = "tcp"
)
const daytimePort = 13
const maxline int = 4096

type DayTimer struct {
	ipaddr string
}

func (dt *DayTimer) GetTime() (timeString string, err error) {
	conn, err := net.Dial(tcp, fmt.Sprintf("%s:13", dt.ipaddr))
	if err != nil {
		return "", err
	}
	buffer := make([]byte, maxline)
	byteRead, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	res := buffer[0:byteRead]
	timeString = string(res[:])
	return
}

// NewDayTimer returns a new Pinger struct pointer
func NewDayTimer(ipaddr string) *DayTimer {
	return &DayTimer{
		ipaddr: ipaddr,
	}
}
