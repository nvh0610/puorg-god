package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	customStatus "god/internal/common/error"
	"god/internal/controller"
	"god/internal/repository"
	"god/job/schedule"
	"god/pkg/logger"
	mdw "god/pkg/middleware"
	"god/pkg/resp"
	"god/platform/mysqldb"
	"god/platform/redisdb"
	"net/http"
)

func InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		resp.Return(writer, http.StatusOK, customStatus.SUCCESS, "success")
	})

	schedule.Start()
	mysqlConn, err := mysqldb.NewMysqlConnection()
	if err != nil {
		logger.Error(err.Error())
	}

	redisConn, err := redisdb.NewRedisConnection()
	if err != nil {
		logger.Error(err.Error())
	}

	baseRepo := repository.NewRegistryRepo(mysqlConn)
	baseController := controller.NewRegistryController(baseRepo, redisConn)

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", baseController.AuthCtrl.Login)
		r.Post("/forget-password", baseController.AuthCtrl.ForgetPassword)
		r.Post("/verify-otp", baseController.AuthCtrl.VerifyOtp)
		r.Post("/reset-password", baseController.AuthCtrl.ResetPassword)
		r.Post("/", baseController.UserCtrl.CreateUser)
		r.With(mdw.JwtMiddleware).Post("/change-password", baseController.AuthCtrl.ChangePassword)
	})

	r.Route("/api/users", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Get("/{id}", baseController.UserCtrl.GetUserById)
		r.Put("/{id}", baseController.UserCtrl.UpdateUser)
		r.Delete("/{id}", baseController.UserCtrl.DeleteUser)
		r.Get("/", baseController.UserCtrl.ListUser)
		r.Post("/update-role", baseController.UserCtrl.UpdateRole)
		r.Get("/me", baseController.UserCtrl.GetMe)
	})

	r.Route("/api/ingredients", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Get("/", baseController.IngredientCtrl.ListIngredient)
	})

	r.Route("/api/recipes", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Post("/", baseController.RecipeCtrl.CreateRecipe)
		r.Get("/distinct-cuisines", baseController.RecipeCtrl.GetDistinctCuisines)
		r.Get("/", baseController.RecipeCtrl.GetListRecipe)
		r.Get("/{id}", baseController.RecipeCtrl.GetRecipeById)
		r.Delete("/{id}", baseController.RecipeCtrl.DeleteRecipeById)
		r.Put("/", baseController.RecipeCtrl.UpdateRecipe)
	})

	return r
}
