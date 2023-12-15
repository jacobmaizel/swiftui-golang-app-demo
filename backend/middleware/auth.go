package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent"
	"github.com/jacobmaizel/swiftui-golang-app-demo/ent/user"
	"github.com/jacobmaizel/swiftui-golang-app-demo/server"
	"github.com/jacobmaizel/swiftui-golang-app-demo/utils"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	// Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		// validator.WithCustomClaims(
		// 	func() validator.CustomClaims {
		// 		return &CustomClaims{}
		// 	},
		// ),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	customAuthHeaderExtractor := func(r *http.Request) (string, error) {
		authHeader := r.Header.Get("Auth")
		if authHeader == "" {
			return "", nil // No error, just no JWT.
		}

		authHeaderParts := strings.Fields(authHeader)
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return "", errors.New("auth header format must be Bearer {token}")
		}

		return authHeaderParts[1], nil
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
		jwtmiddleware.WithTokenExtractor(customAuthHeaderExtractor),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

// Pulls the validated claims from context and get or creates user + profile,
// Sets the user + profile objects into context
func PullClaimsUserConfig() gin.HandlerFunc {

	return func(c *gin.Context) {

		claims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if !ok {
			log.Println("Could not Pull claims from context")
		}

		sub := claims.RegisteredClaims.Subject

		// log.Printf("\n\nSUB::: %s \n\n", sub)
		userDB, err := server.Db.User.Query().Where(user.Subject(sub)).Only(c)

		if err != nil {
			switch e := err.(type) {
			case *ent.NotFoundError:
				log.Printf("Auth Middleware::: user was not found in DB::: %s %s", err, e)

				// User was not found in DB, so we get their info from auth0 and create a profile
				if userDB == nil {
					log.Println("user db was NIL!!")
					userDB, err = server.Db.User.Create().SetSubject(sub).Save(c)

					if err != nil {
						log.Println("Failed to create new user")
					}

					log.Println("User successfully created")

					// Get the users profile info from auth0 service
					userInfo, err := server.Auth0.User.Read(c, sub)
					if err != nil {
						log.Println(" middleware:::Failed to get userinfo from auth0", err)
					}

					first := utils.DerefString(userInfo.GivenName)
					last := utils.DerefString(userInfo.FamilyName)
					picture := utils.DerefString(userInfo.Picture)

					// Using the newly created user from above, we create a profile
					profile, err := server.Db.Profile.Create().SetUser(userDB).SetFirstName(first).SetLastName(last).SetPicture(picture).Save(c)
					if err != nil {
						log.Println("Failed to create profile for new user in middleware", err)
					}

					// Set both the profile and the user into context
					c.Set("profile", profile)
					c.Set("user", userDB)

				}

			case *ent.NotSingularError:
				log.Printf("Auth Middleware::: There were multiple users found with sub %s:::%s", sub, err)
			default:
				log.Println("Error querying for user existence in middleware", err)
			}
		}

		// log.Println("we found a user and we are just setting the contexts :)")

		// user with a profile WAS found! we just need to set the context and get its profile
		c.Set("user", userDB)

		// log.Println(userDB)

		profile, err := server.Db.User.Query().Where(user.ID(userDB.ID)).QueryProfile().Only(c)

		if err != nil {
			switch e := err.(type) {
			case *ent.NotFoundError:
				log.Printf("Auth Middleware::: profile was not found in DB::: %s %s", err, e)

				// profile was not found in DB, so we get their info from auth0 and create a profile
				// log.Println("profile db was NIL!!")

				// Get the users profile info from auth0 service
				uinfo2, err2 := server.Auth0.User.Read(c, sub)

				if err2 != nil {
					log.Fatal(" middleware:::Failed to get userinfo from auth0 when creating profile", err2)
					// log.Println("got error from profile auth0 search:::", err2)
				}

				first := utils.DerefString(uinfo2.GivenName)
				last := utils.DerefString(uinfo2.FamilyName)
				picture := utils.DerefString(uinfo2.Picture)

				// Using the newly created user from above, we create a profile
				createdProfile, err := server.Db.Profile.Create().SetUser(userDB).SetFirstName(first).SetLastName(last).SetPicture(picture).Save(c)

				if err != nil {
					log.Println("Failed to create profile for new user in middleware", err)
				}

				// Set both the profile and the user into context
				c.Set("profile", createdProfile)
				c.Set("user", userDB)

			case *ent.NotSingularError:
				log.Printf("Auth Middleware::: There were multiple profiles found with sub %s:::%s", sub, err)
			default:
				log.Println("Error querying for profile existence in middleware", err)
			}
		}

		c.Set("profile", profile)

		if err != nil {
			log.Println("Failed to find profile in middleware, but we SHOULD have one here", err)
		}

		c.Next()

	}
}
