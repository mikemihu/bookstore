package delivery

import (
	"gotu-bookstore/internal"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookDelivery struct {
	cfg    *config.Cfg
	bookUC internal.BookUC
}

func NewBookDelivery(
	cfg *config.Cfg,
	bookUC internal.BookUC,
) internal.BookDelivery {
	return &BookDelivery{
		cfg:    cfg,
		bookUC: bookUC,
	}
}

func (b *BookDelivery) GetList(c *gin.Context) {
	var req entity.BookGetListRequest
	err := c.BindQuery(&req)
	if err != nil {
		return
	}

	filter := entity.BookFilter{
		Search: req.Search,
	}
	books, err := b.bookUC.GetList(c, filter)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (b *BookDelivery) Get(c *gin.Context) {
	idRaw := c.Param("id")

	id, err := uuid.Parse(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid uuid"})
		return
	}

	filter := entity.BookFilter{
		Book: entity.Book{ID: id},
	}
	book, err := b.bookUC.Get(c, filter)
	if err != nil {
		handlerErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, book)
}
