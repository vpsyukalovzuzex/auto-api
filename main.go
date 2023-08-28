package main

import (
	"Demo/controllers"
	"Demo/db"
	"Demo/repositories"
	"Demo/routers"
	"Demo/usecases"
	"Demo/validators"
	"context"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func main() {
	var e error
	e = godotenv.Load()
	if e != nil {
		log.Fatalln(e)
	}
	ctx := context.Background()
	c := redis.NewClient(&redis.Options{
		Addr: os.Getenv("RDBADDR"),
	})
	if _, e := c.Ping(ctx).Result(); e != nil {
		log.Fatalln(e)
	}
	ac := db.InitAuthCache(c, ctx)
	db := db.Create()
	v := validator.New()
	ur := repositories.InitUserRepository(db)
	cr := repositories.InitCarRepository(db)
	uv := validators.InitUserValidator(v)
	uu := usecases.InitUserUsecase(ur, uv)
	cu := usecases.InitCarUsecase(cr)
	uc := controllers.InitUserController(uu, ac)
	cc := controllers.InitCarController(cu)
	echo := routers.Create(uc, cc)
	e = echo.Start(":8080")
	if e != nil {
		log.Fatalln(e)
	}
}
