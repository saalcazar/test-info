package handler

import (
	"net/http"
)

func RouteProfile(mux *http.ServeMux, storageProfile Storageprofile) {
	h := NewServiceProfile(storageProfile)
	mux.HandleFunc("/v1/profile/create", h.create)
	mux.HandleFunc("/v1/profile/update", h.update)
	mux.HandleFunc("/v1/profile/delete", h.delete)
	mux.HandleFunc("/v1/profile/getall", h.getall)
}

func RouteFounder(mux *http.ServeMux, storageFounder Storagefounder) {
	h := NewServiceFounder(storageFounder)
	mux.HandleFunc("/v1/founder/create", h.create)
	mux.HandleFunc("/v1/founder/update", h.update)
	mux.HandleFunc("/v1/founder/delete", h.delete)
	mux.HandleFunc("/v1/founder/getbyid", h.getById)
	mux.HandleFunc("/v1/founder/getall", h.getall)
}

func RouteUser(mux *http.ServeMux, storageUser Storageuser) {
	h := NewServiceUser(storageUser)
	mux.HandleFunc("/v1/user/create", h.create)
	mux.HandleFunc("/v1/user/update", h.update)
	mux.HandleFunc("/v1/user/delete", h.delete)
	mux.HandleFunc("/v1/user/getall", h.getall)
	mux.HandleFunc("/v1/user/getbyid", h.getById)
}

func RouteProyect(mux *http.ServeMux, storageProyect Storageproyect) {
	h := NewServiceProyect(storageProyect)
	mux.HandleFunc("/v1/proyect/create", h.create)
	mux.HandleFunc("/v1/proyect/update", h.update)
	mux.HandleFunc("/v1/proyect/delete", h.delete)
	mux.HandleFunc("/v1/proyect/getall", h.getall)
	mux.HandleFunc("/v1/proyect/getbyid", h.getById)
}

func RouteEspecific(mux *http.ServeMux, storageEspecific Storageespecific) {
	h := NewServiceEspecific(storageEspecific)
	mux.HandleFunc("/v1/especific/create", h.create)
	mux.HandleFunc("/v1/especific/update", h.update)
	mux.HandleFunc("/v1/especific/delete", h.delete)
	mux.HandleFunc("/v1/especific/getbynameproyect", h.getByNameProyect)
	mux.HandleFunc("/v1/especific/getall", h.getall)
}

func RouteProjectResult(mux *http.ServeMux, storageProjectResult Storageprojectresult) {
	h := NewServiceResult(storageProjectResult)
	mux.HandleFunc("/v1/result/create", h.create)
	mux.HandleFunc("/v1/result/update", h.update)
	mux.HandleFunc("/v1/result/delete", h.delete)
	mux.HandleFunc("/v1/result/getbynameproyect", h.getByNameProyect)
	mux.HandleFunc("/v1/result/getall", h.getall)
}

func RouteProjectActivity(mux *http.ServeMux, storageProjectActivity Storageprojectactivity) {
	h := NewServiceProjectActivity(storageProjectActivity)
	mux.HandleFunc("/v1/projectactivity/create", h.create)
	mux.HandleFunc("/v1/projectactivity/update", h.update)
	mux.HandleFunc("/v1/projectactivity/delete", h.delete)
	mux.HandleFunc("/v1/projectactivity/getbynameproyect", h.getByNameProyect)
	mux.HandleFunc("/v1/projectactivity/getall", h.getall)
}

func RouteActivity(mux *http.ServeMux, storageActivity Storageactivity) {
	h := NewServiceActivity(storageActivity)
	mux.HandleFunc("/v1/activity/create", h.create)
	mux.HandleFunc("/v1/activity/update", h.update)
	mux.HandleFunc("/v1/activity/delete", h.delete)
	mux.HandleFunc("/v1/activity/getall", h.getall)
	mux.HandleFunc("/v1/activity/getbyid", h.getById)
}

func RouteReport(mux *http.ServeMux, storage Storagereport) {
	h := NewServiceReport(storage)
	mux.HandleFunc("/v1/report/create", h.create)
	mux.HandleFunc("/v1/report/update", h.update)
	mux.HandleFunc("/v1/report/delete", h.delete)
	mux.HandleFunc("/v1/report/getall", h.getall)
	mux.HandleFunc("/v1/report/getbyid", h.getById)
}

func RouteQuantitative(mux *http.ServeMux, storageQuantitative Storagequantitative) {
	h := NewServiceQuantitative(storageQuantitative)
	mux.HandleFunc("/v1/quantitative/create", h.create)
	mux.HandleFunc("/v1/quantitative/update", h.update)
	mux.HandleFunc("/v1/quantitative/delete", h.delete)
	mux.HandleFunc("/v1/quantitative/deleteall", h.deleteall)
	mux.HandleFunc("/v1/quantitative/getall", h.getall)
}

func RouteApplication(mux *http.ServeMux, storageApplication Storageapplication) {
	h := NewServiceApplication(storageApplication)
	mux.HandleFunc("/v1/application/create", h.create)
	mux.HandleFunc("/v1/application/update", h.update)
	mux.HandleFunc("/v1/application/updateapproved", h.approved)
	mux.HandleFunc("/v1/application/delete", h.delete)
	mux.HandleFunc("/v1/application/getall", h.getall)
	mux.HandleFunc("/v1/application/getbyid", h.getById)
}

func RouteBudget(mux *http.ServeMux, storageBudget Storagebudget) {
	h := NewServiceBudget(storageBudget)
	mux.HandleFunc("/v1/budget/create", h.create)
	mux.HandleFunc("/v1/budget/update", h.update)
	mux.HandleFunc("/v1/budget/delete", h.delete)
	mux.HandleFunc("/v1/budget/deleteall", h.deleteall)
	mux.HandleFunc("/v1/budget/getall", h.getall)
}

func RouteAccountability(mux *http.ServeMux, storageAccountability Storageaccountability) {
	h := NewServiceAccountability(storageAccountability)
	mux.HandleFunc("/v1/accountability/create", h.create)
	mux.HandleFunc("/v1/accountability/update", h.update)
	mux.HandleFunc("/v1/accountability/updateapproved", h.approved)
	mux.HandleFunc("/v1/accountability/delete", h.delete)
	mux.HandleFunc("/v1/accountability/getall", h.getall)
	mux.HandleFunc("/v1/accountability/getbyid", h.getById)
}

func RouteSurrender(mux *http.ServeMux, storageSurrender Storagesurrender) {
	h := NewServiceSurrender(storageSurrender)
	mux.HandleFunc("/v1/surrender/create", h.create)
	mux.HandleFunc("/v1/surrender/update", h.update)
	mux.HandleFunc("/v1/surrender/delete", h.delete)
	mux.HandleFunc("/v1/surrender/deleteall", h.deleteall)
	mux.HandleFunc("/v1/surrender/getall", h.getall)
}

func RouteDataBase(mux *http.ServeMux, storageDataBase Storagedatabase) {
	h := NewServiceDataBase(storageDataBase)
	mux.HandleFunc("/v1/database/create", h.create)
	mux.HandleFunc("/v1/database/update", h.update)
	mux.HandleFunc("/v1/database/delete", h.delete)
	mux.HandleFunc("/v1/database/getall", h.getall)
}

func RouteLogin(mux *http.ServeMux, storageLogin Storagelogin) {
	h := NewServiceLogin(storageLogin)
	mux.HandleFunc("/v1/login", h.login)
}

func RouteSuperLogin(mux *http.ServeMux, storageSuperLogin StorageSuperlogin) {
	h := NewServiceSuperLogin(storageSuperLogin)
	mux.HandleFunc("/v1/superlogin", h.superlogin)
}
