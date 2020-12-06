package jd

const JD_SERVER_URL = "https://router.jd.com/api"

const v = "1.0"
const format = "json"
const sign_method = "md5"

//京粉精选商品查询接口
const METHOD_JINGFEN = "jd.union.open.goods.jingfen.query"

// JDClient服务
type JDClient struct {
}

type JingfenRequest struct {
	//频道ID:1-好券商品,2-精选卖场,10-9.9包邮,15-京东配送,22-实时热销榜,23-为你推荐,24-数码家电,25-超市,26-母婴玩具,27-家具日用,28-美妆穿搭,30-图书文具,31-今日必推,32-京东好物,33-京东秒杀,34-拼购商品,40-高收益榜,41-自营热卖榜,108-秒杀进行中,109-新品首发,110-自营,112-京东爆品,125-首购商品,129-高佣榜单,130-视频商品,153-历史最低价商品榜，210-极速版商品
	EliteId   int
	PageIndex int
	PageSize  int
	//排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortName string
	//	asc,desc升降序,默认降序
	Sort string
	// 联盟id_应用id_推广位id，三段式
	Pid string
	//支持出参数据筛选，逗号','分隔，目前可用：videoInfo,documentInfo
	Fields string
	// 10微信京东购物小程序禁售，11微信京喜小程序禁售
	ForbidTypes string
}

type JingfenResponse struct {
}

// 京粉接口 https://union.jd.com/openplatform/api/10417
func (client *JDClient) JingfenQuery(req JingfenRequest) (*JingfenResponse, error) {
	return nil, nil
}
