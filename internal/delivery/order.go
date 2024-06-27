package delivery

import (
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderDelivery struct {
	cfg     *config.Cfg
	orderUC internal.OrderUC
}

func NewOrderDelivery(
	cfg *config.Cfg,
	orderUC internal.OrderUC,
) internal.OrderDelivery {
	return &OrderDelivery{
		cfg:     cfg,
		orderUC: orderUC,
	}
}

func (b *OrderDelivery) GetList(c *gin.Context) {
	orders, err := b.orderUC.GetList(c)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (b *OrderDelivery) Get(c *gin.Context) {
	idRaw := c.Param("id")

	id, err := uuid.Parse(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid uuid"})
		return
	}

	order, err := b.orderUC.Get(c, id)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, order)
}

func (b *OrderDelivery) Create(c *gin.Context) {
	var req entity.OrderCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// validation
	if len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "items is required"})
		return
	}
	for i := range req.Items {
		item := &req.Items[i]
		if item.BookID == uuid.Nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid book id"})
			return
		}
		if item.Qty <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid qty"})
			return
		}
	}

	id, err := b.orderUC.Create(c, req)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id.String()})
}
