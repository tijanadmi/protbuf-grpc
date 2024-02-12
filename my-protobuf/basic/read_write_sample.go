package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(msg proto.Message, fname string) {
	b, err := proto.Marshal(msg)

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

func ReadProtoFromFile(fname string, dest proto.Message) {
	// Open a file for reading

	// Read the file's contents
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can not open file", err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal", err)
	}
}

func dummyUser() basic.User {
	addr := basic.Address{
		Street:  "Daily Planet",
		City:    "Metropolis",
		Country: "US",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  40.70797893425118,
			Longitude: -74.01163838107261,
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{"fly": 5, "speed": 5, "durability": 4}

	return basic.User{
		Id:                   99,
		Username:             "superman",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Address:              &addr,
		CommunicationChannel: &comm,
		SkillRating:          sr,
	}
}

func WriteToFileSample() {
	u := dummyUser()
	WriteProtoToFile(&u, "superman_file.bin")
}

func ReadFromFileSample() {
	dest := basic.User{}

	ReadProtoFromFile("superman_file.bin", &dest)

	log.Println(&dest)
}
