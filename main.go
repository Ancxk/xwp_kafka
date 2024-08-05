package main

import (
	"time"
	"xwp_kafka/kafka_"
)

func main() {
	Ka()
}

type Tes struct {
	Num  int    `json:"num,omitempty"`
	Name string `json:"name,omitempty"`
}

func Te() {

}

func Ka() {
	kafka_.CreateTopic()
	for i := 0; i < 10; i++ {
		go func() {
			for {
				//time.Sleep(10 * time.Millisecond)
				kafka_.WriteMsg2()
			}
		}()
	}
	for i := 0; i < 10; i++ {
		go kafka_.ReadMsg()
	}
	time.Sleep(10 * time.Minute)
}
