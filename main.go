package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/saalcazar/ceadlbk-info/autorization"
	"github.com/saalcazar/ceadlbk-info/handler"
	"github.com/saalcazar/ceadlbk-info/storage"
)

func main() {

	err := autorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se cargaron los certificados: %v", err)
	}

	storage.NewPostgresDB()

	//Super Login
	storageSuperLogin := storage.NewPsqlSuperLogin(storage.Pool())
	serviceSuperLogin := handler.NewServiceSuperLogin(storageSuperLogin)

	//Login
	storageLogin := storage.NewPsqlLogin(storage.Pool())
	serviceLogin := handler.NewServiceLogin(storageLogin)

	//Profile
	storageProfile := storage.NewPsqlProfile(storage.Pool())
	serviceProfile := handler.NewServiceProfile(storageProfile)

	//Founder
	storageFounder := storage.NewPsqlFounder(storage.Pool())
	serviceFounder := handler.NewServiceFounder(storageFounder)

	//User
	storageUser := storage.NewPsqlUser(storage.Pool())
	serviceUser := handler.NewServiceUser(storageUser)

	//Proyect
	storageProyect := storage.NewPsqlProyect(storage.Pool())
	serviceProyect := handler.NewServiceProyect(storageProyect)

	//Especific
	storageEspecific := storage.NewPsqlEspecific(storage.Pool())
	serviceEspecific := handler.NewServiceEspecific(storageEspecific)

	//Result
	storageResult := storage.NewPsqlResult(storage.Pool())
	serviceResult := handler.NewServiceResult(storageResult)

	//Project Activity
	storageProjectActivity := storage.NewPsqlProjectActivity(storage.Pool())
	serviceProjectActivity := handler.NewServiceProjectActivity(storageProjectActivity)

	//Activity
	storageActivity := storage.NewPsqlActivity(storage.Pool())
	serviceActivity := handler.NewServiceActivity(storageActivity)

	//Report
	storageReports := storage.NewPsqlReport(storage.Pool())
	serviceReport := handler.NewServiceReport(storageReports)

	//Quantitative
	storageQuantitative := storage.NewPsqlQuantitative(storage.Pool())
	serviceQuantitative := handler.NewServiceQuantitative(storageQuantitative)

	//Application
	storageApplication := storage.NewPsqlApplication(storage.Pool())
	serviceApplication := handler.NewServiceApplication(storageApplication)

	//Budget
	storageBudget := storage.NewPsqlBudget(storage.Pool())
	serviceBudget := handler.NewServiceBudget(storageBudget)

	//Accountability
	storageAccountability := storage.NewPsqlAccountability(storage.Pool())
	serviceAccountability := handler.NewServiceAccountability(storageAccountability)

	//Surrender
	storageSurrender := storage.NewPsqlSurrender(storage.Pool())
	serviceSurrender := handler.NewServiceSurrender(storageSurrender)

	//DataBase
	storageDataBase := storage.NewPsqlDataBase(storage.Pool())
	serviceDataBase := handler.NewServiceDataBase(storageDataBase)

	mux := http.NewServeMux()

	handler.RouteSuperLogin(mux, serviceSuperLogin)
	handler.RouteLogin(mux, serviceLogin)
	handler.RouteProfile(mux, serviceProfile)
	handler.RouteFounder(mux, serviceFounder)
	handler.RouteUser(mux, serviceUser)
	handler.RouteProyect(mux, serviceProyect)
	handler.RouteEspecific(mux, serviceEspecific)
	handler.RouteProjectResult(mux, serviceResult)
	handler.RouteProjectActivity(mux, serviceProjectActivity)
	handler.RouteActivity(mux, serviceActivity)
	handler.RouteReport(mux, serviceReport)
	handler.RouteQuantitative(mux, serviceQuantitative)
	handler.RouteApplication(mux, serviceApplication)
	handler.RouteBudget(mux, serviceBudget)
	handler.RouteAccountability(mux, serviceAccountability)
	handler.RouteSurrender(mux, serviceSurrender)
	handler.RouteDataBase(mux, serviceDataBase)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://ceadl-info.vercel.app", "https://ceadl-info.vercel.app/login", "http://localhost:5173", "https://app.ceadl.org.bo"},
		AllowedMethods: []string{"DELETE", "GET", "POST", "PUT"},
	})

	handler := c.Handler(mux)

	http.ListenAndServe(":8080", handler)

}
