package main

import (
	"fmt"
	"log"
	"net"
	// "netpackets4go/utils"
)

const (
	HOST = "localhost"
	PORT = "8023"
	TYPE = "tcp"
)

// an individual number used to represent each packet. Global in order to iterate and not start over
var GlobalPacketNumber uint = 0

// each packet contains 256 characters, which is roughly a paragraph depending on the length of the sentences 
type Packets struct {
	Number []uint
	Value []string;
};

// replacing the vehemently religious, comically inept, disgustingly abundant copy-paste of this statement: "if err != nil"
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {	
	log.Fatal(ServerFunc("sup bruh, it's me, yo mama, this is a bunch of random filler text. Have a wonderful day!!!!!!!!!!!!!!!!!!!!! Yesterday I went out to eat food and it was like literally the greatest food on earth like I was so happy omg it was like the best food in the entire world like wow omg I'm going crazy it's like so amazing omg this food is just like omg like wow I'm absolute going crazy lmao :) and this is likke super cool and I'm like super happy and this is just the best thing ever becuase I have this really awesome TCP connection and it's just really amazing because I'm just really happy and stuff like whatever I'm just so like really just like so amazing and like whatever because everything is so like just so amazing like omg it's just so amazing like omg I'm just so happy like whatever it's just so amazing I just love life because everything in the world is just so amazing and I'm just so happy :)"))
}




// turns big pieces of data into packets of 1024 bytes each (I think that's 256 characters)
func (packets Packets) Packetizer(str string) Packets {
	dataAmountPerPacket := 256
	// purePacketsPerString := math.floor(len(str)/dataAmountPerPacket)
	for i := 0; i<len(str); i++ {
		GlobalPacketNumber += 1
		packets.Number = append(packets.Number, GlobalPacketNumber)

		
		// fmt.Println(i % dataAmountPerPacket)
		// adding individual packets
		if i % dataAmountPerPacket == 0 && i >= 256 {
			fmt.Println("NO REMAINDER")
			packets.Value = append(packets.Value, str[i-dataAmountPerPacket : i ])
		} else {
		// checking if there are no more full packets left
			if len(str[i : len(str)]) <= dataAmountPerPacket && len(str[0 : i-1]) % dataAmountPerPacket == 0 {
				fmt.Println("well?")
				// fmt.Println(str[i-1 : len(str)])	
				packets.Value = append(packets.Value, str[i-1 : len(str)])
				
			}
		}
		/*else {
			fmt.Println("REMAINDER")
			packets.Value = append(packets.Value, str[i : len(str)])
		}*/ 
				
	}	
	return packets
}


func ServerFunc(str string) error {
	packets := Packets{}
	packet := packets.Packetizer(str)
	fmt.Println("starting server")
	
	ln, err := net.Listen(TYPE, ":"+PORT)
	check(err)

	conn, err := ln.Accept()
	check(err)
	
	fmt.Println(packet.Value[0])
	fmt.Println(packet.Value[1])
	// server sends individual packets
	for i := 0; i<len(packet.Value); i++ {
		fmt.Println(packet.Value[1] + "\n")
		_, _ = conn.Write([]byte(packet.Value[i]))

	}

	return err
}
