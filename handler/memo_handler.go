package main

import (
	"checkapp/model"
	"homepage/model"
	"homepage/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMemos(c echo.Context) error {
	memos, err := repository.GetAllMemos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, memos)
}

func CreateMemo(c echo.Context) error {
	var memo model.Memo
	if err := c.Bind(&memo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := repository.InsertMemo(memo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, memo)
}

func DeleteMemos(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteMemo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent((http.StatusNoContent))
}
