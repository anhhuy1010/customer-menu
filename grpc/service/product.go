package service

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// pbStore "github.com/Diqit-A1-Branch/cpos-microservice-store/grpc/proto/store"

	pb "github.com/anhhuy1010/customer-menu/grpc/proto/product"

	// "github.com/Diqit-A1-Branch/cpos-microservice-tenant/helpers/util"
	"github.com/anhhuy1010/customer-menu/models"
)

type ProductService struct {
}

func NewProductServer() pb.ProductServer {
	return &ProductService{}
}

func (s *ProductService) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.DetailResponse, error) {
	conditions := bson.M{"uuid": req.Uuid}
	date, error := time.Parse("2006-01-02", req.Date)
	if error != nil {
		fmt.Println(error)
	}

	if req.Date != "" {
		conditions["start_date"] = bson.M{"$lte": date}
		conditions["end_date"] = bson.M{"$gte": date}
	}
	fmt.Println(conditions)
	result, err := new(models.Products).FindOne(conditions)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res := &pb.DetailResponse{
		Uuid:        result.Uuid,
		Image:       result.Image,
		Name:        result.Name,
		Sequence:    int32(result.Sequence),
		Quantity:    int32(result.Sequence),
		Description: result.Description,
		Gallery:     result.Gallery,
		IsDelete:    int32(result.IsDelete),
		IsActive:    int32(result.IsActive),
		StartDate:   result.StartDate.Format("2006-01-02"),
		EndDate:     result.EndDate.Format("2006-01-02"),
		CreatedAt:   result.CreatedAt.Format("2006-01-02"),
		UpdatedAt:   result.UpdatedAt.Format("2006-01-02"),
	}

	return res, nil
}
