package page

//import "github.com/golang/glog"

type MysqlPageHelper struct {
	/*
	 * @param pageNum  页码
     * @param pageSize 每页显示数量
	 * @param count    是否进行count查询
	 */
	pageNum uint64
	pageSize uint64
	totalNum uint64
	count bool
}

func NewMysqlPageHelper(ps uint64)*MysqlPageHelper{
	h := new(MysqlPageHelper)
	h.count = true
	h.pageSize = ps
	return h
}

func (mph *MysqlPageHelper)StartPage(data []map[string]interface{})*PageInfo{
	page := newPageInfoWithNavigatePages(data,mph.pageNum,mph.pageSize,mph.totalNum,10)
	return page
}

func (mph *MysqlPageHelper)NeedCount() bool{
	if mph.count{
		if mph.totalNum>0 {
			mph.count = false
		}
	}
	return mph.count

}


func (mph *MysqlPageHelper)CalPages(curPage,totalNum uint64)(startPos,enPos uint64){
	mph.totalNum = totalNum
	mph.pageNum = curPage

	var totalPages uint64
	if totalNum%mph.pageSize >0{
		totalPages = totalNum/mph.pageSize +1
	}else{
		totalPages = totalPages/mph.pageSize
	}
	//glog.Info("totalPages~~~~~~~~~~",totalPages,"  ; curPage~~~~~~~~~~~",curPage)
	if curPage >0 && curPage<=totalPages-1{
		//glog.Info("~~~~1")
		return (curPage-1)*mph.pageSize,curPage*mph.pageSize
	}else if curPage==totalPages{
		//glog.Info("~~~~2")
		return (curPage-1)*mph.pageSize,mph.totalNum
	}else{
		return 0,0
	}
}

func (mph *MysqlPageHelper)SimpleCalPages(curPage uint64)(startPos,enPos uint64){
	return mph.CalPages(curPage,mph.totalNum)
}