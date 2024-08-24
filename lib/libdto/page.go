package libdto

type PageForm struct {
	OrderBy     []string `form:"orderBy" json:"orderBy" `        // 字段
	OrderDirect string   `form:"orderDirect" json:"orderDirect"` // 方式
	Page        int      `form:"page,default_router=1" json:"page,default_router=1"`
	Psize       int      `form:"psize,default_router=20" json:"psize,default_router=20"`
}

// Limit 获取每页记录数
func (p PageForm) Limit() int {
	if p.Psize == 0 {
		return 20
	}
	return p.Psize
}

// Offset 获取当前页开始ID
func (p PageForm) Offset() int {
	psize := p.Psize
	if psize == 0 {
		psize = 20
	}

	if p.Page == 0 {
		return 0
	}

	return (p.Page - 1) * psize
}
