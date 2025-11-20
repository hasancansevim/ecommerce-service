package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StoreController struct {
	storeService service.IStoreService
	BaseController
}

func NewStoreController(storeService service.IStoreService) *StoreController {
	return &StoreController{storeService: storeService}
}

func (storeController *StoreController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/stores", storeController.GetAllStores)
	e.GET("/api/v1/stores/:id", storeController.GetStoreById)
	e.POST("/api/v1/stores", storeController.AddStore)
	e.DELETE("/api/v1/stores/:id", storeController.DeleteStore)
	e.PUT("/api/v1/stores/:id", storeController.UpdateStore)
}

func (storeController *StoreController) GetAllStores(c echo.Context) error {
	stores := storeController.storeService.GetAllStores()
	return storeController.Success(c, stores, "Tüm Mağazalar Getirildi")
}

func (storeController *StoreController) GetStoreById(c echo.Context) error {
	id, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		return storeController.BadRequest(c, convertErr)
	}
	store, err := storeController.storeService.GetStoreById(uint(id))
	if err != nil {
		return storeController.BadRequest(c, err)
	}
	return storeController.Success(c, store, "Mağaza Getirildi")
}

func (storeController *StoreController) AddStore(c echo.Context) error {
	var addStoreRequest request.AddStoreRequest
	if bindErr := c.Bind(&addStoreRequest); bindErr != nil {
		return storeController.BadRequest(c, bindErr)
	}
	err := storeController.storeService.AddStore(addStoreRequest.ToModel())
	if err != nil {
		return storeController.BadRequest(c, err)
	}
	return storeController.Success(c, addStoreRequest, "Mağaza Eklendi")
}

func (storeController *StoreController) DeleteStore(c echo.Context) error {
	id, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		return storeController.BadRequest(c, convertErr)
	}
	err := storeController.storeService.DeleteStoreById(uint(id))
	if err != nil {
		return storeController.BadRequest(c, err)
	}
	return storeController.Success(c, nil, "Mağaza Silindi")
}

func (storeController *StoreController) UpdateStore(c echo.Context) error {
	id, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		return storeController.BadRequest(c, convertErr)
	}
	var updatedStoreRequest request.AddStoreRequest
	if bindErr := c.Bind(&updatedStoreRequest); bindErr != nil {
		return storeController.BadRequest(c, bindErr)
	}
	err := storeController.storeService.UpdateStoreById(uint(id), updatedStoreRequest.ToModel())
	if err != nil {
		return storeController.BadRequest(c, err)
	}
	return storeController.Success(c, nil, "Mağaza Güncellendi")
}
