package routes

import (
	"github.com/gabriel2443/workOutProject/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux{

r:= chi.NewRouter()

r.Group(func (r chi.Router){
	r.Use(app.Middleware.Authenticate)
	// r.Use(app.Middleware.Authenticate) // that checks all the required users

	r.Get("/workouts/{id}", app.Middleware.RequireUser( app.WorkoutHandler.HandleGetWorkoutByID))

    r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.HandleCreateWorkout))
    r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleUpdateWorkoutByID))
     r.Delete ("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleDeleteByID))

})

r.Get("/health", app.HealthCheck)
r.Post("/users", app.UserHandler.HandlerRegisterUser)
r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

return r

}