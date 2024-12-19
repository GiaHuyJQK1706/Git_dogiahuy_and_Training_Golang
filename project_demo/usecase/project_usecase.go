package usecase

import (
	"project_demo/dto"
	"project_demo/entities"
	"project_demo/repository"
)

type ProjectUsecase interface {
	CreateProject(dto.CreateProjectRequest) (entities.Project, error)
	UpdateProject(id uint, request dto.CreateProjectRequest) (entities.Project, error)
	DeleteProject(id uint) error
	GetProjectByID(id uint) (entities.Project, error)
	ListProjects() ([]entities.Project, error)
}

type projectUsecase struct {
	PrọectRepo repository.ProjectRepository
}

// Tao project moi
func (puc *projectUsecase) CreateProject(request dto.CreateProjectRequest) (entities.Project, error) {
	project := entities.Project{
		Name:             request.Name,
		Category:         request.Category,
		ProjectSpend:     request.ProjectSpend,
		ProjectVariance:  request.ProjectVariance,
		ProjectStartedAt: *request.ProjectStartedAt,
		ProjectEndedAt:   request.ProjectEndedAt,
	}

	err := puc.PrọectRepo.Create(&project)
	return project, err
}

// Cap nhat project
func (puc *projectUsecase) UpdateProject(id uint, request dto.CreateProjectRequest) (entities.Project, error) {
	project, err := puc.PrọectRepo.GetByID(id)
	if err != nil {
		return project, err
	}

	project.Name = request.Name
	project.Category = request.Category
	project.ProjectSpend = request.ProjectSpend
	project.ProjectVariance = request.ProjectVariance
	project.ProjectStartedAt = *request.ProjectStartedAt
	project.ProjectEndedAt = request.ProjectEndedAt

	err = puc.PrọectRepo.Update(&project)
	return project, err
}

// Xoa project
func (puc *projectUsecase) DeleteProject(id uint) error {
	return puc.PrọectRepo.Delete(id)
}

// Lay (Xem chi tiet) project theo ID
func (puc *projectUsecase) GetProjectByID(id uint) (entities.Project, error) {
	return puc.PrọectRepo.GetByID(id)
}

// Lay (Xem) danh sach tat ca cac project
func (puc *projectUsecase) ListProjects() ([]entities.Project, error) {
	return puc.PrọectRepo.List()
}

// DI
func NewProjectUsecase(repo repository.ProjectRepository) ProjectUsecase {
	return &projectUsecase{PrọectRepo: repo}
}
