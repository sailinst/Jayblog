package controllers

func (c *UserController) API_Profile() {
	type user struct {
		Userid string
		Hobby  []string
	}
	u := user{
		" Jay",
		[]string{"chess", "code"},
	}

	c.Data["JSON"] = u
	c.ServeJSON()
}
