package handler

import "github.com/saalcazar/ceadlbk-info/model"

type StorageSuperlogin interface {
	SuperLogin(*model.SuperLogin) bool
}

type Storagelogin interface {
	Login(*model.Login) (bool, []*model.DataUser, error)
}

type Storageprofile interface {
	Create(*model.Profile) error
	Update(*model.Profile) error
	Delete(uint) error
	GetAll() (model.Profiles, error)
}

type Storageuser interface {
	Create(*model.User) error
	Update(*model.User) error
	Delete(uint) error
	GetByID(uint) (*model.User, error)
	GetAll() (model.Users, error)
}

type Storagefounder interface {
	Create(*model.Founder) error
	Update(*model.Founder) error
	Delete(uint) error
	GetByID(uint) (*model.Founder, error)
	GetAll() (model.Founders, error)
}

type Storageproyect interface {
	Create(*model.Proyect) error
	Update(*model.Proyect) error
	Delete(uint) error
	GetByID(uint) (*model.Proyect, error)
	GetAll() (model.Proyects, error)
}

type Storageespecific interface {
	Create(model.Especifics) error
	Update(*model.Especific) error
	Delete(uint) error
	GetByNameProyect(string) (model.Especifics, error)
	GetAll() (model.Especifics, error)
}

type Storageprojectresult interface {
	Create(model.ProjectResults) error
	Update(*model.ProjectResult) error
	Delete(uint) error
	GetByNameProyect(string) (model.ProjectResults, error)
	GetAll() (model.ProjectResults, error)
}

type Storageprojectactivity interface {
	Create(model.ProjectActivities) error
	Update(*model.ProjectActivity) error
	Delete(uint) error
	GetByNameProyect(string) (model.ProjectActivities, error)
	GetAll() (model.ProjectActivities, error)
}

type Storageactivity interface {
	Create(*model.Activity) error
	Update(*model.Activity) error
	Delete(uint) error
	GetByID(uint) (*model.Activity, error)
	GetAll() (model.Activities, error)
}

type Storagereport interface {
	Create(*model.Report) error
	Update(*model.Report) error
	Delete(uint) error
	GetByID(uint) (*model.Report, error)
	// GetByTitle(string) (*model.Report, error)
	GetAll() (model.Reports, error)
}

type Storagequantitative interface {
	Create(model.Quantitatives) error
	Update(*model.Quantitative) error
	Delete(uint) error
	DeleteAll(uint) error
	GetAll() (model.Quantitatives, error)
}

type Storageapplication interface {
	Create(*model.Application) error
	Update(*model.Application) error
	Approved(*model.Application) error
	Delete(uint) error
	GetByID(uint) (*model.Application, error)
	GetAll() (model.Applications, error)
}

type Storagebudget interface {
	Create(model.Budgets) error
	Update(*model.Budget) error
	Delete(uint) error
	DeleteAll(uint) error
	GetAll() (model.Budgets, error)
}

type Storageaccountability interface {
	Create(*model.Accountability) error
	Update(*model.Accountability) error
	Approved(*model.Accountability) error
	Delete(uint) error
	GetByID(uint) (*model.Accountability, error)
	GetAll() (model.Accountabilities, error)
}

type Storagesurrender interface {
	Create(model.Surrenders) error
	Update(*model.Surrender) error
	Delete(uint) error
	DeleteAll(uint) error
	GetAll() (model.Surrenders, error)
}

type Storagedatabase interface {
	Create(model.DataBases) error
	Update(*model.DataBase) error
	Delete(uint) error
	GetAll() (model.DataBases, error)
}
