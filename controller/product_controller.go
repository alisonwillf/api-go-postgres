package controller

import (
	"net/http"
	"strconv"

	"github.com/alisonferreira/APIGO/model"
	"github.com/alisonferreira/APIGO/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUsecase *usecase.ProductUsecase
	//Usecase
}

func NewProductController(usecase *usecase.ProductUsecase) *productController {
	return &productController{
		ProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p *productController) GetProductsById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		reponse := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Product ID is required",
		}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		reponse := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Product ID is required",
		}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	product, err := p.ProductUsecase.GetProductsById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if product == nil {
		response := model.Response{
			Message: "Produto nao encontrado!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
