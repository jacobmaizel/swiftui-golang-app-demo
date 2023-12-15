package server

import (
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/auth0/go-auth0/management"
	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/redis/rueidis"
)

var (
	Db       *ent.Client
	Http     *gin.Engine
	Auth0    *management.Management
	Firebase *firebase.App
	Fcm      *messaging.Client
	Redis    rueidis.Client
	L        *log.Logger
)
