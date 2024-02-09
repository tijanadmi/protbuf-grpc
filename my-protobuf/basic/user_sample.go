package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

// new instance of protobuf message
func BasicUser() {

	addr := basic.Address{
		Street:  "Daily Planet",
		City:    "Metropolis",
		Country: "US",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  40.70797893425118,
			Longitude: -74.01163838107261,
		},
	}

	h := basic.User{
		Id:       99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	log.Println(&h)
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       98,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
			"id":97,
			"username":"batman",
			"is_active":true,
			"password":"YmF0bWFucGFzc3dvcmQ=",
			"emails":[
				 "batman@movie.com",
				 "batman@dc.com"
			],
			"gender":"GENDER_MALE"
	 }
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Got error : ", err)
	}

	log.Println(&p)
}
