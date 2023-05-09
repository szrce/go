package main

import (
	"encoding/base64"
	"fmt"
	"github.com/nareix/joy4/format/rtsp"
)

func main() {

	uri := "rtsp://admin:admin@192.168.90.102:8554/live"

	rtsp.ErrCodecDataChange = fmt.Errorf("rtsp: codec data change, please call HandleCodecDataChange()")

	rtsp.DebugRtp = true
	rtsp.DebugRtsp = true
	rtsp.SkipErrRtpBlock = true

	username := "admin"
	password := "admin"

	req := rtsp.Request{
		Header: []string{
			fmt.Sprintf(`Authorization: Basic %s`, base64.StdEncoding.EncodeToString([]byte(username+":"+password))),
		},
	}

	session, err := rtsp.Dial(uri)
	session.Headers = append(req.Header)
	session.Options()
	session.Streams()

	if err != nil {
		panic(err)
		return
	}

	for {
		packet, err := session.ReadPacket()
		if err != nil {
			panic(err)
		}
		fmt.Println(packet.Data)
	}

}
