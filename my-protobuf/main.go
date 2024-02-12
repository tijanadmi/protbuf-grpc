package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//basic.BasicHello()
	basic.BasicUser()
	//basic.BasicUnmarshalAnyKnown()
	//basic.BasicUnmarshalAnyNotKnown()
	//basic.BasicUnmarshalAnyIs()
	//basic.BasicOneof()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
	//basic.BasicUserGroup()
	//jobsearch.JobSearchCandidate()
	//jobsearch.JobSearchSoftware()
	//basic.WriteToFileSample()
	//basic.ReadFromFileSample()
	//basic.WriteToJSONSample()
	//basic.ReadFromJSONSample()
	basic.BasicWriteUserContentV1()
	basic.BasicReadUserContentV1()
}
