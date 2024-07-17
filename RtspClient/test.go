package main

import (
	"log"
	"os"

	"github.com/VKCOM/joy4/av/avutil"
	"github.com/VKCOM/joy4/av/pubsub"
	"github.com/VKCOM/joy4/format"
	"github.com/VKCOM/joy4/format/rtsp"
	"github.com/nareix/joy4/av/pubsub"
)

func main() {
	// Register RTSP and MP4 formats
	format.RegisterAll()
	rtsp.DebugRtsp = true

	// Create a new RTSP client
	s, err := rtsp.Dial("rtsp://admin:admin@192.168.90.102:8554/live")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// Create a new MP4 file for writing
	out, err := os.Create("output.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Create a new pubsub to subscribe to the RTSP stream
	ps := pubsub.NewSimpleSubscriber(s.Streams())
	// Create an MP4 Muxer for writing the stream to an MP4 file
	muxer, err := avutil.NewMuxerForFormat("mp4", out)
	if err != nil {
		log.Fatal(err)
	}
	defer muxer.Close()

	// Write the header to the MP4 file
	err = muxer.WriteHeader(ps.Streams())
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the RTSP stream packets and write them to the MP4 file
	for {
		pkt, err := ps.ReadPacket()
		if err != nil {
			log.Fatal(err)
		}
		if pkt == nil {
			break
		}
		err = muxer.WritePacket(pkt)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Close the MP4 muxer
	err = muxer.WriteTrailer()
	if err != nil {
		log.Fatal(err)
	}
}
