package routes

import (
	"github.com/gabriel2443/workOutProject/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux{

r:= chi.NewRouter()
r.Get("/health", app.HealthCheck)
r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)
r.Delete ("/workouts/{id}", app.WorkoutHandler.HandleDeleteByID)

r.Post("/users", app.UserHandler.HandlerRegisterUser)
r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

return r

}