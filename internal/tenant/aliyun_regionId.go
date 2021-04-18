package tenant

// reference: https://help.aliyun.com/document_detail/40654.html?spm=a2c6h.13066369.0.0.54a17471zHMvv8#title-vlh-oy6-405
type AliyunRegionID int

const (
	qingdao     AliyunRegionID = iota + 1 // 青岛
	beijing                               // 北京
	zhangjiakou                           // 张家口
	huhehaote                             // 呼和浩特
	wulanchabu                            // 乌兰察布
	hangzhou                              // 杭州
	shanghai                              // 上海
	shenzhen                              // 深圳
	heyuan                                // 河源
	guangzhou                             // 广州
	chengdu                               // 成都
)
