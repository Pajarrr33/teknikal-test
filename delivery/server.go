package delivery

import (
	"database/sql"
	"fmt"
	"os"

	"teknikal-test/config"
	"teknikal-test/delivery/controller"
	"teknikal-test/delivery/middleware"
	"teknikal-test/repository"
	"teknikal-test/service"
	"teknikal-test/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Server struct {
	jwtService service.JWTService
	authUc     usecase.AuthUsecase
	payUc      usecase.PaymentUsecase
	authMid    middleware.AuthMiddleware
	router     *gin.Engine
	host       string
}

func (s *Server) initRoutes() {
	routerGroup := s.router.Group(config.API_V1)

	controller.NewAuthController(s.authUc, routerGroup, s.authMid).Route()
	controller.NewPaymentController(s.payUc, routerGroup, s.authMid).Route()
}

func (s *Server) Run() {
	s.initRoutes()
	err := s.router.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	config, _ := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Name)

	db, err := sql.Open(config.DbConfig.Driver, dsn)
	if err != nil {
		panic("connection error")
	}

	// userRepo := repository.NewUserRepository(db)
	custRepo := repository.NewCustomerRepository(db)
	expRepo := repository.NewExpiredRepository(db)
	transRepo := repository.NewTransactionRepository(db)
	mercRepo := repository.NewMerchantRepository(db)
	jwtService := service.NewJWTService(config.JwtConfig)
	authUc := usecase.NewAuthUsecase(custRepo,jwtService,expRepo)
	payUc := usecase.NewPaymentUsecase(transRepo,custRepo,mercRepo)


	authMid := middleware.NewAuthMiddleware(jwtService, expRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", config.ApiConfig.ApiPort)

	file, err := os.OpenFile("assets/history.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to create a log file")
	}
	logrus.SetOutput(file)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	
	return &Server{
		jwtService: jwtService,
		authUc:     authUc,
		payUc:      payUc,
		authMid: 	authMid,
		router:     engine,
		host:       host,
	}
}