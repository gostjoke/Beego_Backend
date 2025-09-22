package controllers

import (
	"Beego_Backend/models"
	"encoding/json"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

// GET /users
func (c *UserController) GetAll() {
	o := orm.NewOrm()
	var users []models.User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

// GET /users/:id
func (c *UserController) GetOne() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	user := models.User{Id: id}
	err := o.Read(&user)
	if err == nil {
		c.Data["json"] = user
	} else {
		c.Data["json"] = map[string]string{"error": "not found"}
	}
	c.ServeJSON()
}

// POST /users
func (c *UserController) Create() {
	var user models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		o := orm.NewOrm()
		_, err := o.Insert(&user)
		if err == nil {
			c.Data["json"] = user
		} else {
			c.Data["json"] = map[string]string{"error": err.Error()}
		}
	} else {
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}

// PUT /users/:id
func (c *UserController) Update() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	user := models.User{Id: id}
	if o.Read(&user) == nil {
		var newUser models.User
		json.Unmarshal(c.Ctx.Input.RequestBody, &newUser)
		user.Name = newUser.Name
		user.Email = newUser.Email
		o.Update(&user)
		c.Data["json"] = user
	} else {
		c.Data["json"] = map[string]string{"error": "not found"}
	}
	c.ServeJSON()
}

// DELETE /users/:id
func (c *UserController) Delete() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	if num, err := o.Delete(&models.User{Id: id}); err == nil {
		c.Data["json"] = map[string]interface{}{"deleted": num}
	} else {
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJSON()
}
