package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 定义我们自己的菜单对象
type SystemMenu struct {
	Id       int    `json:"id"`        //id
	FatherId int    `json:"father_id"` //上级菜单id
	Name     string `json:"name"`      //菜单名
	Route    string `json:"route"`     //页面路径
	Icon     string `json:"icon"`      //图标路径
}

func TestGenerateTree(t *testing.T) {
	// 模拟获取数据库中所有菜单，在其它所有的查询中，也是首先将数据库中所有数据查询出来放到数组中，
	// 后面的遍历递归，都在这个 allMenu中进行，而不是在数据库中进行递归查询，减小数据库压力。
	allMenu := []SystemMenu{
		{Id: 1, FatherId: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, FatherId: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},

		{Id: 3, FatherId: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, FatherId: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},

		{Id: 5, FatherId: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, FatherId: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		{Id: 7, FatherId: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	// 构建树节点配置
	var conf TreeConfig
	conf.IdName = "id"         // 与结构体json标签对应
	conf.PidName = "father_id" // 与结构体json标签对应
	// 扩展字段,将原有的key映射新增成另一个key
	// mp := make(map[string]string, 2)
	// mp["tkName"] = "name"
	// mp["tkRoute"] = "route"
	// conf.Mapper = mp
	// 扩展字段,在原有的结构体上添加新的key
	// ep := make(map[string]any, 2)
	// ep["mystring"] = "mengbao"
	// ep["mybool"] = true
	// ep["mynum"] = 123
	// conf.Extend = ep

	// 生成树
	resp := GenerateTree(allMenu, &conf)
	if resp == nil {
		return
	}
	bytes, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(bytes))
	// fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))
}
