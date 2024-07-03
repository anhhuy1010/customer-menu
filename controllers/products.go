package controllers

import (
	"fmt"
	"math"
	"net/http"

	"github.com/anhhuy1010/customer-menu/helpers/respond"

	"github.com/anhhuy1010/customer-menu/models"
	request "github.com/anhhuy1010/customer-menu/request/product"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductController struct {
}

func (productClt ProductController) List(c *gin.Context) {
	productModel := new(models.Products)
	var req request.GetListRequest
	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	cond := bson.M{
		"is_delete": 0,
		"is_active": 1,
		"price":     bson.M{"$gt": 0},
	}
	if req.Date != nil {
		cond["start_date"] = bson.M{"$lte": req.Date}
		cond["end_date"] = bson.M{"$gte": req.Date}
	}
	optionsQuery, page, limit := models.GetPagingOption(req.Page, req.Limit, req.Sort)
	var respData []request.ListResponse
	productt, _ := productModel.Pagination(c, cond, optionsQuery)

	for _, productt := range productt {

		res := request.ListResponse{
			Name:      productt.Name,
			Price:     productt.Price,
			Image:     productt.Image,
			StartDate: productt.StartDate,
			EndDate:   productt.EndDate,
			Quantity:  productt.Quantity,
			Sequence:  productt.Sequence,
		}
		respData = append(respData, res)
	}
	total, err := productModel.Count(c, cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	pages := int(math.Ceil(float64(total) / float64(limit)))
	c.JSON(http.StatusOK, respond.SuccessPagination(respData, page, limit, pages, total))
}

func (productClt ProductController) Detail(c *gin.Context) {
	productModel := new(models.Products)
	var reqUri request.GetDetailUri
	err := c.ShouldBindUri(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}

	condition := bson.M{"uuid": reqUri.Uuid}
	productt, err := productModel.FindOne(condition)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, respond.ErrorCommon("Product no found!"))
		return
	}

	response := request.GetDetailResponse{
		Price:       productt.Price,
		Image:       productt.Image,
		Name:        productt.Name,
		Sequence:    productt.Sequence,
		Quantity:    productt.Quantity,
		Description: productt.Description,
		Gallery:     productt.Gallery,
		StartDate:   productt.StartDate,
		EndDate:     productt.EndDate,
	}
	c.JSON(http.StatusOK, respond.Success(response, "Successfully"))
}
