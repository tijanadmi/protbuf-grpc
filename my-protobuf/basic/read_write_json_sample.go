package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln("Can not open file", err)
	}
	defer f.Close()

	// Write some bytes to the file
	_, err = f.Write(b)
	if err != nil {
		log.Fatalln("Can not write to file", err)
	}

	// Close the file
	err = f.Close()
	if err != nil {
		log.Fatalln("Can not close the file", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	// Open a file for reading

	// Read the file's contents
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can not open file", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}

func ReadFromJSONSample() {
	dest := basic.User{}

	ReadProtoFromJSON("superman_file.json", &dest)

	log.Println(&dest)
}
