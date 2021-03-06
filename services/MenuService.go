package services

import (
	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
)

type MenuService struct {
	interfaces.IMenuRepository
}

func (service *MenuService) GetMenu(idMenu int) (models.Menu, error) {

	menu, err := service.FetchMenu(idMenu)
	if err != nil {
		return models.Menu{}, err
	}

	return menu, nil
}

func (service *MenuService) GetMenus() ([]models.Menu, error) {

	menus, err := service.FetchAllMenu()
	if err != nil {
		return []models.Menu{}, err
	}

	return menus, nil
}

func (service *MenuService) GetMenuReceipt(idMenu int) ([]models.MenuIngredient, error) {

	menuIngredients, err := service.FetchMenuIngredients(idMenu)
	if err != nil {
		return []models.MenuIngredient{}, err
	}

	return menuIngredients, nil
}
