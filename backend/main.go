package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	firebase "firebase.google.com/go"
	"github.com/auth0/go-auth0/management"
	cors "github.com/gin-contrib/cors"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	_ "github.com/jacobmaizel/swiftui-golang-app-demo/ent/runtime"
	"github.com/jacobmaizel/swiftui-golang-app-demo/middleware"
	"github.com/jacobmaizel/swiftui-golang-app-demo/redis"
	"github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/tasks"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/rueidis"
	"google.golang.org/api/option"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/actions"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/competitions"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/config"
	goals "github.com/jacobmaizel/swiftui-golang-app-demo/api/goals"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/invites"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/notifications"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/profiles"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/services"
	"github.com/jacobmaizel/swiftui-golang-app-demo/api/squads"
	workouts "github.com/jacobmaizel/swiftui-golang-app-demo/api/workouts"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func createEntClient(db_url string) (*ent.Client, error) {
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(15)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("postgres", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init() {

	// Load Env Vars

	//if this is dev, we load from env file.
	env := os.Getenv("ENVIRONMENT")

	if env == "DEV" || env == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading the .env file: %v", err)
		}
	}

	// Sentry, only run in production/staging
	if env == "STAGING" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:           os.Getenv("SENTRY_DSN"),
			Environment:   os.Getenv("ENVIRONMENT"),
			EnableTracing: true,
			Debug:         os.Getenv("ENVIRONMENT") == "DEV",
			Release:       "1.0",
			BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
				// if hint.Context != nil {
				// if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
				// You have access to the original Request here

				// log.Println("req here from before send in sentry init:::", req)
				// }
				// }

				return event
			},

			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
			// TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
			// 	// As an example, this custom sampler does not send some
			// 	// transactions to Sentry based on their name.
			// 	hub := sentry.GetHubFromContext(ctx.Span.Context())
			// 	log.Println("got the hub here boy:::", hub)
			// 	// name := hub.Scope().Transaction()
			// 	// if name == "GET /favicon.ico" {
			// 	// 	return 0.0
			// 	// }

			// 	return 1.0
			// }),
		}); err != nil {
			log.Fatalf("Sentry initialization failed: %v\n", err)
		}
	}

	db_url := os.Getenv("DATABASE_URL")

	// client, cliError := createEntClient(db_url)

	/* if cliError != nil { */
	/* 	log.Fatal("Failed to create ent client") */
	/* } */

	// client, err := ent.Open("postgres", db_url)

	client, err := createEntClient(db_url)

	if err != nil {
		log.Fatalf("Could not enable client:::%s", err)
	}

	server.Db = client

	server.Db.Use(
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				start := time.Now()
				defer func() {
					log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
				}()
				return next.Mutate(ctx, m)
			})
		})

	//Auth0 management client
	opts := management.WithClientCredentialsAndAudience(context.Background(), os.Getenv("AUTH0_MANAGEMENT_CLIENT_ID"), os.Getenv("AUTH0_MANAGEMENT_SECRET"), os.Getenv("AUTH0_MANAGEMENT_AUDIENCE"))

	var deb management.Option

	if env == "DEV" {
		deb = management.WithDebug(true)
	} else {
		deb = management.WithDebug(false)
	}

	m, err := management.New(os.Getenv("AUTH0_DOMAIN"), deb, opts)

	if err != nil {
		log.Fatalf("failed to create auth0 client!! %s", err)
	}

	server.Auth0 = m
}

func runServer() {

	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		log.Fatal("ENVIRONMENT must be set")
	}

	var r *gin.Engine
	var allowedOrigins []string

	corsConfig := cors.DefaultConfig()

	switch env {
	case "DEV":
		r = gin.Default()
		allowedOrigins = append(allowedOrigins, "http://localhost:8080/")
		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			// log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
		}
	case "STAGING":
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		allowedOrigins = append(allowedOrigins, "https://swiftui-golang-app-demo-staging.herokuapp.com/")
	case "TEST":
		gin.SetMode(gin.TestMode)
		r = gin.Default()
	default:
		log.Fatal("Invalid environment type, not DEV or STAGING")
	}

	server.Http = r

	corsConfig.AllowOrigins = allowedOrigins

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("swiftui-golang-app-demo"),
		newrelic.ConfigLicense(""),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Attaching New Relic Middleware")
		r.Use(nrgin.Middleware(app))
	}

	r.Use(cors.New(corsConfig))

	// r.Use(middleware.Logger())
	r.Use(sentrygin.New(sentrygin.Options{}))
	r.Use(adapter.Wrap(middleware.EnsureValidToken()))
	r.Use(middleware.PullClaimsUserConfig())

	// str := fmt.Sprintf("hi from %s", env)

	// sentry.CaptureMessage(str)

	base := r.Group("/")

	// base.GET("/allusers", allUsersProfiles)

	// base.GET("/me", me)

	profiles.AddProfileRoutes(base)
	competitions.AddCompetitionRoutes(base)
	actions.AddActionRoutes(base)
	invites.AddInviteRoutes(base)
	squads.AddSquadRoutes(base)
	workouts.AddWorkoutRoutes(base)
	goals.AddGoalRoutes(base)
	services.AddServicesRoutes(base)
	notifications.AddNotificationRoutes(base)
	config.AddConfigRoutes(base)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// log.Printf("Starting API with::: PORT %s and ENV %s", port, env)

	r.Run(":" + port)

}

func initFirebaseClient(c context.Context) error {

	fbSk := os.Getenv("FIREBASE_CONFIG")

	opts := []option.ClientOption{option.WithCredentialsJSON([]byte(fbSk))}
	app, err := firebase.NewApp(c, nil, opts...)

	if err != nil {
		log.Println("\n\nFAILED TO CREATE FIREBASE APP")
		return err
	}

	server.Firebase = app

	server.Fcm, err = app.Messaging(c)

	if err != nil {
		log.Println("\n\nFAILED TO CREATE FIREBASE MESSAGING CLIENT")
		return err
	}

	return nil

}

func initRedisClient() error {
	log.Println("\n\nINITIALIZING REDIS CLIENT")
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{os.Getenv("REDIS_URL")}, Password: os.Getenv("REDIS_PASSWORD"), Username: os.Getenv("REDIS_USERNAME"), DisableCache: true})

	if err != nil {
		log.Println(err)
	}
	// defer client.Close()

	server.Redis = client

	// flush all redis data and remove indices
	// fa := client.B().Flushall().Build()

	// // do multi
	// res := server.Redis.Do(context.Background(), fa)

	// if res.Error() != nil {
	// 	log.Println("error flushing redis:::", res.Error())
	// }

	return nil

}

func configureLogger() {
	log.Println("STARTUP:::Configuring Logger")

	server.L = log.New(os.Stdout, "[swiftui-golang-app-demo]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

	server.L.Println("STARTUP:::Done Configuring Logger")

}

func main() {
	c := context.Background()
	Init()

	configureLogger()

	if err := initFirebaseClient(c); err != nil {
		// log.Println("hey there", reflect.TypeOf(err))
		log.Fatalln("error initializing firebase client:::", err)
	}
	log.Println("initialized firebase client")

	if err := initRedisClient(); err != nil {
		log.Fatalln("error initializing redis client:::", err)
	}

	redis.CreateRedisIndices(c)

	redis.InsertProfileDocuments(c)
	redis.InsertCompetitionDocuments(c)
	redis.InsertSquadDocuments(c)

	done := make(chan bool)

	go tasks.CompAboutToEndNotiTask(c, done)
	go tasks.CompAboutToStartNotiTask(c, done)

	// t := server.Redis.B().Publish().Channel("test").Message("boom").Build()
	// server.Redis.Do(c, t)

	runServer()

}

// type TaskFunc func(uuid.UUID)

// type Task struct {
// 	interval *time.
// 	run TaskFunc
// 	oneOff bool

// }
