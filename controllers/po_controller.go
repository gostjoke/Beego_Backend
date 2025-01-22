package controllers

import (
	"Beego_Backend/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type PoController struct {
	web.Controller
}

func (c *PoController) GetPo() {
	// 從路徑參數中獲取 ID
	idStr := c.Ctx.Input.Param(":id")
	poID, err := strconv.Atoi(idStr) // 將字串轉換為整數
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid po_id"}
		c.ServeJSON()
		return
	}

	// ORM 操作
	o := orm.NewOrm()
	po := models.Po_table{Id: poID}

	// 查詢主表
	if err := o.Read(&po); err != nil {
		c.Data["json"] = map[string]string{"error": "Po_table not found"}
		c.ServeJSON()
		return
	}

	// 加載相關的 Pn_table 數據
	if _, err := o.LoadRelated(&po, "Pn_tables"); err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to load related Pn_tables"}
		c.ServeJSON()
		return
	}

	// 返回 JSON
	c.Data["json"] = po
	c.ServeJSON()
}
