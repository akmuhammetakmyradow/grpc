package main

import (
	"context"
	"fmt"
	"grpc/api"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("Starting server...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "tm_mr", CityName: "Mary"},
			&api.CityEntry{CityCode: "tm_ag", CityName: "Ashgabat"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest, res api.WeatherService_QueryWeatherServer) error {
	for {
		err := res.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 10})

		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}
