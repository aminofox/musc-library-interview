package main

import (
	"context"
	"fmt"
	"music-libray-management/config"
	"music-libray-management/infra/mongo"
	"music-libray-management/infra/mongo/repository"
	"music-libray-management/infra/s3"
	repository2 "music-libray-management/infra/s3/repository"
	jwtPkg "music-libray-management/internal/jwt"
	passwordPkg "music-libray-management/internal/password"
	responsePkg "music-libray-management/internal/response"
	"music-libray-management/middlewares"
	"music-libray-management/routers"
	"music-libray-management/usecase/album"
	"music-libray-management/usecase/artist"
	"music-libray-management/usecase/auth"
	"music-libray-management/usecase/document"
	"music-libray-management/usecase/playlist"
	"music-libray-management/usecase/track"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
)

type App struct {
	config   *config.Environment
	mongoDb  *mongo.MongoDB
	s3Client *s3.S3Client
}

func loadEnvironment() *config.Environment {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Msgf("fail loading environment variables: %v", err)
	}
	return cfg
}

func loadMongoDB(cfg *config.MongoEnv) *mongo.MongoDB {
	db, err := mongo.Connect(cfg)
	if err != nil {
		log.Fatal().Msgf("fail load mongo database, %v", err)
	}
	return db
}

func loadS3Client(cfg *config.Environment) *s3.S3Client {
	client, err := s3.Connect(cfg)
	if err != nil {
		log.Fatal().Msgf("fail load s3 client, %v", err)
	}
	return client
}

func main() {
	cfg := loadEnvironment()
	log.Printf("AppName: %s", cfg.AppName)

	mongoDb := loadMongoDB(&cfg.MongoEnv)
	s3Client := loadS3Client(cfg)
	app := &App{
		config:   cfg,
		mongoDb:  mongoDb,
		s3Client: s3Client,
	}

	// Service
	jwtService := jwtPkg.NewJwtService(app.config)
	passwordService := passwordPkg.NewPasswordService()
	responseService := responsePkg.NewServiceResponse()

	// Repository
	userRepository := repository.NewUserRepository(app.mongoDb)
	trackRepository := repository.NewTrackRepository(app.mongoDb)
	documentRepository := repository.NewDocumentRepository(app.mongoDb)
	playlistRepository := repository.NewPlaylistRepository(app.mongoDb)
	albumRepository := repository.NewAlbumRepository(app.mongoDb)
	artistRepository := repository.NewArtistRepository(app.mongoDb)
	storageRepository := repository2.NewStorageRepository(cfg, s3Client)

	// UseCase
	authUseCase := auth.NewAuthUseCase(app.config, jwtService, passwordService, userRepository, app.mongoDb)
	trackUseCase := track.NewTrackUseCase(trackRepository, artistRepository, albumRepository, app.mongoDb)
	documentUseCase := document.NewDocumentUseCase(documentRepository, storageRepository, app.mongoDb)
	playlistUseCase := playlist.NewPlaylistUseCase(playlistRepository, app.mongoDb)
	albumUseCase := album.NewAlbumUseCase(albumRepository, app.mongoDb)
	artistUseCase := artist.NewArtistUseCase(artistRepository, app.mongoDb)

	// Middleware
	middleware := middlewares.NewMiddleware(
		jwtService,
		userRepository,
	)

	router := routers.InitRouter(
		app.config,
		middleware,
		authUseCase,
		trackUseCase,
		documentUseCase,
		playlistUseCase,
		artistUseCase,
		albumUseCase,
		responseService,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.AppPort),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	go func() {
		log.Printf("Start HTTP Server, Listening: %d", app.config.AppPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Start HTTP Server failed... Error: %s", err.Error())
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Stop server shutdown error: %v", err.Error())
	} else {
		logrus.Info("Stopped serving on Services")
	}
}
