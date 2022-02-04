package repo

func (c *Confirm) Get(code string) (string, bool) {
	c.Mx.Lock()
	defer c.Mx.Unlock()
	val , ok := c.Users[code]
	return val, ok
}

func (c *Confirm)Insert(email , code string){
	c.Mx.Lock()
	defer c.Mx.Unlock()
	c.Users[code] = email
}