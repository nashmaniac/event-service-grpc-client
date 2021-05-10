package main

import (
	"context"
	"encoding/json"
	"event-service-grpc-client/pb"
	"event-service-grpc-client/serializer"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"io/ioutil"
	"log"
	"os"
)

func ConfigureClient() pb.EventClient {
	connectionString := "[::]:8000"
	cc, err := grpc.Dial(connectionString, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	client := pb.NewEventClient(cc)
	return client
}

func mockCreateRequest(client pb.EventClient)  {
	structData := readJsonFile("samples/data.json")
	input := &pb.EventInputParams{
		Email:         "raju@gmail.com",
		Environment:   "staging",
		Component:     "analytics",
		MessageString: "The order is placed successfully",
		Data:          structData,
	}
	res, err := client.CreateEvent(context.Background(), input)
	if err != nil {
		log.Println(err)
	}
	log.Println(serializer.ConvertProtobufToJson(res))
}

func readJsonFile(fileName string) *structpb.Struct {
	current_dir, _ := os.Getwd()
	filePath := current_dir+"/" + fileName
	bytes, _ := ioutil.ReadFile(filePath)
	var b map[string]interface{}
	_ = json.Unmarshal(bytes, &b)
	structBody, _ := structpb.NewStruct(b)
	return structBody
}

func mockFilterRequest(client pb.EventClient)  {
	queryParams := map[string]string {
		//"email": "shetu2153@gmail.com",
		//"component": "orders",
		//"from": "2021-05-10",
		//"message": "successfully",
	}
	eventFilterInput := &pb.EventFilterInput{Params: queryParams}
	res, err := client.FilterEvents(context.Background(), eventFilterInput)
	if err != nil {
		log.Println(err.Error());
	}
	log.Println(len(res.Events))
	log.Println(serializer.ConvertProtobufToJson(res))
}

func main() {
	log.Println("Initializing Client")
	client := ConfigureClient()
	mockFilterRequest(client)
	//mockCreateRequest(client)
}