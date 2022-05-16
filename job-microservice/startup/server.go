package startup

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-TIM1-2022/dislinkt/util/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"io"
	"job-microservice/application"
	"job-microservice/infrastructure/api"
	"job-microservice/infrastructure/persistance"
	"job-microservice/model"
	"job-microservice/startup/config"
	"log"
	"net"
)

type Server struct {
	config      *config.Config
	tracer      otgo.Tracer
	closer      io.Closer
	mongoClient *mongo.Client
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init(config.JobServiceName)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
	}
}

func (server *Server) GetTracer() otgo.Tracer {
	return server.tracer
}

func (server *Server) GetCloser() io.Closer {
	return server.closer
}

func (server *Server) Start() {
	server.mongoClient = server.initMongoClient()
	jobStore := server.initStoreStore(server.mongoClient)
	jobService := server.initJobService(jobStore, server.config)

	jobHandler := server.initJobHandler(jobService)

	server.startGrpcServer(jobHandler)
}

func (server *Server) Stop() {
	log.Println("stopping server")
	server.mongoClient.Disconnect(context.TODO())
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistance.GetClient(server.config.JobDBHost, server.config.JobDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) startGrpcServer(jobHandler *api.JobHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	log.Println(fmt.Sprintf("started grpc server on localhost:%s", server.config.Port))
	//////////////userService.RegisterUserServiceServer(grpcServer, jobHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initStoreStore(client *mongo.Client) model.JobStore {
	store := persistance.NewJobMongoDBStore(client)
	return store
}

func (server *Server) initJobService(store model.JobStore, config *config.Config) *application.JobService {
	return application.NewJobService(store, config)
}

func (server *Server) initJobHandler(
	service *application.JobService) *api.JobHandler {
	return api.NewJobHandler(service)
}
