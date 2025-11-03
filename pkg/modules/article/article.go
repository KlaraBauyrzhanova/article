package article

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"article/pkg/store/article"
)

type Service struct {
	store *article.PostgresRepository
	db    *sqlx.DB
}

// NewService creates endpoints
func NewService(store *article.PostgresRepository, db *sqlx.DB, e *echo.Echo) Service {
	s := Service{
		store: store,
		db:    db,
	}

	return s
}

func (s *Service) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	a := &article.Article{
		Id: id,
	}
	ctx := c.Request().Context()
	err = s.store.Get(ctx, a)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, a)
}

func (s *Service) Create(c echo.Context) error {
	a := &article.Article{}
	if err := c.Bind(&a); err != nil {
		return c.String(http.StatusBadRequest, "failed to bind: "+err.Error())
	}

	if a.Title == "" {
		return c.String(http.StatusBadRequest, "title is required")
	}

	a.CreatedDate = time.Now()

	ctx := c.Request().Context()
	err := s.store.Create(ctx, a)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to save article: "+err.Error())
	}

	return c.JSON(http.StatusCreated, a)
}
