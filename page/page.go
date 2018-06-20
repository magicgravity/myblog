package page

import (
	"bytes"
	"fmt"
	//"github.com/golang/glog"
)

/*
参考 com.github.pagehelper.PageInfo
 */
type PageInfo struct {
	pageNum uint64			//当前页
	pageSize uint64 		//每页的数量
	size uint64 			//当前页的数量

	//由于startRow和endRow不常用，这里说个具体的用法
	//可以在页面中"显示startRow到endRow 共size条数据"

	startRow uint64 		//当前页面第一个元素在数据库中的行号
	endRow uint64 			//当前页面最后一个元素在数据库中的行号
	total uint64			//总记录数

	pages uint64			//总页数

	list []map[string]interface{}		//结果集

	prePage uint64			//前一页
	nextPage uint64			//下一页

	isFirstPage bool		//是否为第一页 默认false
	isLastPage bool 		//是否为最后一页 默认false
	hasPreviousPage bool 	//是否有前一页 默认false
	hasNextPage bool		//是否有下一页 默认false
	navigatePages uint64	//导航页码数
	navigatepageNums []uint64	//所有导航页号
	navigateFirstPage uint64 	//导航条上的第一页
	navigateLastPage uint64		//导航条上的最后一页
}


func (p *PageInfo)ToString()string{
	buf := bytes.Buffer{}
	buf.WriteString("PageInfo=>{\r\n")
	buf.WriteString(fmt.Sprintf("总记录数total:%d \r\n",p.GetTotal()))
	buf.WriteString(fmt.Sprintf("总页数  pages:%d \r\n",p.GetPages()))
	buf.WriteString(fmt.Sprintf("当前页码pageNum:%d \r\n",p.GetPageNum()))
	buf.WriteString(fmt.Sprintf("每页数量pageSize:%d \r\n",p.GetPageSize()))
	buf.WriteString(fmt.Sprintf("当前页数量size:%d\r\n",p.GetSize()))
	buf.WriteString(fmt.Sprintf("当前第一个元素行号startRow:%d \r\n",p.GetStartRow()))
	buf.WriteString(fmt.Sprintf("当前最后一个元素行号endRow:%d \r\n",p.GetEndRow()))
	if p.IsHasPreviousPage(){
		buf.WriteString(fmt.Sprintf("前一页页码prePage:%d\r\n",p.GetPrePage()))
	}
	if p.IsHasNextPage(){
		buf.WriteString(fmt.Sprintf("下一页页码nextPage:%d\r\n",p.GetNextPage()))
	}
	if p.IsIsFirstPage(){
		buf.WriteString("当前页为第一页\r\n")
	}
	if p.IsIsLastPage(){
		buf.WriteString("当前页为最后一页\r\n")
	}
	if len(p.GetNavigatepageNums())>0 {
		buf.WriteString(fmt.Sprintf("所有导航页页号:[%v]\r\n",p.GetNavigatepageNums()))
	}
	buf.WriteString("}\r\n")

	return buf.String()
}


func newPageInfoWithNavigatePages(data []map[string]interface{},pageNum,pageSize,total,navigatePages uint64)*PageInfo{
	p := new(PageInfo)
	p.list = data
	p.pageNum = pageNum
	p.pageSize = pageSize
	if total%pageSize>0 {
		p.pages = total/pageSize+1
	}else{
		p.pages = total/pageSize
	}
	p.size =  uint64(len(data))
	p.total = total
	p.startRow = p.pageSize*(p.pageNum-1)
	if p.size >0 {
		p.endRow = p.pageSize*(p.pageNum-1) + p.size-1
	}else{
		p.endRow = p.startRow
	}
	p.navigatePages = navigatePages
	p.isFirstPage = true
	p.isLastPage = false
	p.hasNextPage = false
	p.hasPreviousPage = false

	if len(p.list)>0 {
		p.calcNavigatepageNums()
		p.calcPage()
		p.judgePageBoudary()
	}
	return p
}

/**
* 计算导航页
*/
func (p *PageInfo)calcNavigatepageNums(){
	/*
	 //当总页数小于或等于导航页码数时
        if (pages <= navigatePages) {
            navigatepageNums = new int[pages];
            for (int i = 0; i < pages; i++) {
                navigatepageNums[i] = i + 1;
            }
        } else { //当总页数大于导航页码数时
            navigatepageNums = new int[navigatePages];
            int startNum = pageNum - navigatePages / 2;
            int endNum = pageNum + navigatePages / 2;

            if (startNum < 1) {
                startNum = 1;
                //(最前navigatePages页
                for (int i = 0; i < navigatePages; i++) {
                    navigatepageNums[i] = startNum++;
                }
            } else if (endNum > pages) {
                endNum = pages;
                //最后navigatePages页
                for (int i = navigatePages - 1; i >= 0; i--) {
                    navigatepageNums[i] = endNum--;
                }
            } else {
                //所有中间页
                for (int i = 0; i < navigatePages; i++) {
                    navigatepageNums[i] = startNum++;
                }
            }
        }
	 */

	if p.pages <= p.navigatePages {
		//当总页数小于或等于导航页码数时
		p.navigatepageNums = make([]uint64,p.pages)
		for i :=0;uint64(i)<p.pages;i++ {
			p.navigatepageNums[i] = uint64(i+1)
		}
	}else {
		//当总页数大于导航页码数时
		p.navigatepageNums = make([]uint64,p.navigatePages)
		startNum := p.pageNum-p.navigatePages/2
		endNum := p.pageNum+p.navigatePages/2
		if startNum <1 {
			startNum = 1
			//(最前navigatePages页
			for i,j:=0,startNum;uint64(i)<p.navigatePages;i++ {
				p.navigatepageNums[i] = j
				j++
			}
		}else if endNum> p.pages{
			endNum = p.pages
			//最后navigatePages页
			//glog.Info("navigatePages==",p.navigatePages,">>>>len(navigatepageNums)==",len(p.navigatepageNums))
			for i,j:= p.navigatePages-1,endNum;i>=0; {
			//	glog.Info("i~~~",i)
				p.navigatepageNums[i] = j
				j--
				if i>0{
					i--
				}else{
					break
				}
			}
		}else{
			//所有中间页
			for i,j:=0,startNum;uint64(i)<p.navigatePages;i++ {
				p.navigatepageNums[i] = j
				j++
			}
		}
	}
}

/*
计算前后页，第一页，最后一页
 */
func (p *PageInfo)calcPage(){
	/*
	 if (navigatepageNums != null && navigatepageNums.length > 0) {
            navigateFirstPage = navigatepageNums[0];
            navigateLastPage = navigatepageNums[navigatepageNums.length - 1];
            if (pageNum > 1) {
                prePage = pageNum - 1;
            }
            if (pageNum < pages) {
                nextPage = pageNum + 1;
            }
        }
	 */
	if p.navigatepageNums != nil && len(p.navigatepageNums)>0 {
		p.navigateFirstPage = p.navigatepageNums[0]
		p.navigateLastPage = p.navigatepageNums[len(p.navigatepageNums)-1]
		if p.pageNum >1 {
			p.prePage = p.pageNum -1
		}
		if p.pageNum < p.pages{
			p.nextPage = p.pageNum +1
		}
	}
}

/**
* 判定页面边界
*/
func (p *PageInfo)judgePageBoudary()  {
	/*
	isFirstPage = pageNum == 1;
        isLastPage = pageNum == pages;
        hasPreviousPage = pageNum > 1;
        hasNextPage = pageNum < pages;
	 */
	 if p.pageNum ==1 {
	 	p.isFirstPage = true
	 }else{
	 	p.isFirstPage = false
	 }
	 if p.pageNum == p.pages {
	 	p.isLastPage = true
	 }else{
	 	p.isLastPage = false
	 }

	 if p.pageNum > 1 {
	 	p.hasPreviousPage = true
	 }else{
	 	p.hasPreviousPage = false
	 }
	 if p.pageNum < p.pages {
	 	p.hasNextPage = true
	 }else{
	 	p.hasNextPage = false
	 }
}

func (p *PageInfo)GetPageNum()uint64{
	return p.pageNum
}

func (p *PageInfo)SetPageNum(pn uint64){
	p.pageNum = pn
}

func (p *PageInfo)GetPageSize()uint64{
	return p.pageSize
}

func (p *PageInfo)SetPageSize(ps uint64){
	p.pageSize = ps
}

func (p *PageInfo)GetSize()uint64{
	return p.size
}

func (p *PageInfo)SetSize(s uint64){
	p.size = s
}

func (p *PageInfo)GetStartRow()uint64{
	return p.startRow
}

func (p *PageInfo)SetStartRow(sr uint64){
	p.startRow = sr
}

func (p *PageInfo)GetEndRow()uint64{
	return p.endRow
}

func (p *PageInfo)SetEndRow(er uint64){
	p.endRow = er
}

func (p *PageInfo)GetTotal()uint64{
	return p.total
}

func (p *PageInfo)SetTotal(t uint64){
	p.total = t
}

func (p *PageInfo)GetPages()uint64{
	return p.pages
}

func (p *PageInfo)SetPages(ps uint64){
	p.pages = ps
}

func (p *PageInfo)GetList()[]map[string]interface{}{
	return p.list
}

func (p *PageInfo)SetList(list []map[string]interface{}){
	copy(p.list,list)
}

func (p *PageInfo)GetPrePage()uint64{
	return p.prePage
}

func (p *PageInfo)SetPrePage(pp uint64){
	p.prePage = pp
}

func (p *PageInfo)GetNextPage()uint64{
	return p.nextPage
}

func (p *PageInfo)SetNextPage(np uint64){
	p.nextPage = np
}


func (p *PageInfo)IsIsFirstPage() bool{
	return p.isFirstPage
}

func (p *PageInfo)IsIsLastPage() bool{
	return p.isLastPage
}

func (p *PageInfo)IsHasPreviousPage() bool{
	return p.hasPreviousPage
}

func (p *PageInfo)IsHasNextPage() bool{
	return p.hasNextPage
}

func (p *PageInfo)GetNavigatePages() uint64{
	return p.navigatePages
}

func (p *PageInfo)GetNavigatepageNums() []uint64{
	return p.navigatepageNums
}

func (p *PageInfo)GetNavigateFirstPage() uint64{
	return p.navigateFirstPage
}

func (p *PageInfo)GetNavigateLastPage() uint64{
	return p.navigateLastPage
}


