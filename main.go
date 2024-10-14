package main

import (
	"fmt"
	"golang-test/controller/topupcont"
	"golang-test/controller/transfercont"
	"golang-test/controller/userscont"
	"golang-test/db"
	"golang-test/middleware"
	"golang-test/repository/tbltopups"
	"golang-test/repository/tbltransfer"
	"golang-test/repository/tblusers"
	"golang-test/usecase/accounts"
	"golang-test/usecase/transfers"
	"golang-test/usecase/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := tblusers.NewUserRepository(db)
	userUsecase := users.NewUsersUsecase(userRepo)
	userCont := userscont.NewUsersCont(userUsecase)

	topuprepo := tbltopups.NewUserRepository(db)
	topupusecase := accounts.NewAccountUsecase(topuprepo)
	topupCont := topupcont.NewTopUpController(topupusecase)

	transferRepo := tbltransfer.NewTransferRepository(db)
	transferUsecase := transfers.NewTransferUsecase(transferRepo, userRepo)
	transferCont := transfercont.NewTransferController(transferUsecase)

	router := gin.Default()

	router.POST("/register", userCont.Register)
	router.POST("/login", userCont.Login)

	router.Use(middleware.AuthMiddleware())

	router.POST("/topup", topupCont.TopUp)
	router.POST("/transfer", transferCont.Transfer)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Server failed to start", err)
	}
}
