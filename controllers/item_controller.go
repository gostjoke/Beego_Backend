package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type ItemController struct {
	web.Controller
}

// GET /items
func (c *ItemController) Get() {
	items := []map[string]interface{}{
		{
			"id":   1,
			"name": "81-23312-200",
			"price_history": map[string]interface{}{
				"PreviousPrice": 100,
				"CurrentPrice":  200,
			},
		},
		{
			"id":   2,
			"name": "85-23552-100",
			"price_history": map[string]interface{}{
				"PreviousPrice": 300,
				"CurrentPrice":  400,
			},
		},
	}
	c.Data["json"] = items
	c.ServeJSON()
}

// POST /items
func (c *ItemController) Post() {
	var newItem map[string]interface{}
	c.ParseForm(&newItem) // 獲取請求中的數據
	newItem["id"] = 3     // 模擬添加數據
	c.Data["json"] = newItem
	c.ServeJSON()
}

// GET /items/:id
func (c *ItemController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	item := map[string]interface{}{
		"id":    id,
		"name":  "item" + id,
		"price": 100 * len(id), // 模擬查詢
	}
	c.Data["json"] = item
	c.ServeJSON()
}

// PUT /items/:id
func (c *ItemController) Put() {
	id := c.Ctx.Input.Param(":id")
	updatedItem := map[string]interface{}{
		"id":    id,
		"name":  "updated_item",
		"price": 300,
	}
	c.Data["json"] = updatedItem
	c.ServeJSON()
}

// DELETE /items/:id
func (c *ItemController) Delete() {
	id := c.Ctx.Input.Param(":id")
	c.Data["json"] = map[string]string{"message": "Item " + id + " deleted"}
	c.ServeJSON()
}
