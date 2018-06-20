package page

type PageHelper interface {
	/**
     * 开始分页
     *
     * @param pageNum  页码
     * @param pageSize 每页显示数量
	 * @param count    是否进行count查询
     */
	StartPage(list []map[string]interface{})*PageInfo

	NeedCount()bool

	CalPages(curPage,totalNum uint64)(startPos,enPos uint64)

	SimpleCalPages(curPage uint64)(startPos,enPos uint64)
}


