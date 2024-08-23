# go-listtotree

`golang`中实现快速实现切片list转树tree结构。只有一个接口GenerateTree，传入对应的参数即可。

ps: 想要一个类似java的Hutool那样简单便捷的构建树，找了一圈没有找到合适。如果你用过Hutool生成树，那很容易就能上手，因为就是参照Hutool做的。如果帮助到你，请点个star以帮助到更多的人。

## 1 使用

```shell
go get github.com/U8tou/go-listtotree
```

## 2 用例

```go
// 定义我们自己的菜单对象
type SystemMenu struct {
	Id       int    `json:"id"`        //id
	FatherId int    `json:"father_id"` //上级菜单id
	Name     string `json:"name"`      //菜单名
	Route    string `json:"route"`     //页面路径
	Icon     string `json:"icon"`      //图标路径
}

func TestGenerateTree(t *testing.T) {
	// 模拟获取数据库中所有菜单,待操作数据源
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
}
```
执行结果：
```json
[
  {
    "children": [
      {
        "children": [
          {
            "father_id": 3,
            "icon": "icon-device",
            "id": 6,
            "name": "设备",
            "route": "/device"
          },
          {
            "father_id": 3,
            "icon": "icon-device",
            "id": 7,
            "name": "机柜",
            "route": "/device"
          }
        ],
        "father_id": 1,
        "icon": "icon-asset",
        "id": 3,
        "name": "资产",
        "route": "/asset"
      },
      {
        "father_id": 1,
        "icon": "icon-pe",
        "id": 4,
        "name": "动环",
        "route": "/pe"
      }
    ],
    "father_id": 0,
    "icon": "icon-system",
    "id": 1,
    "name": "系统总览",
    "route": "/systemOverview"
  },
  {
    "children": [
      {
        "father_id": 2,
        "icon": "icon-menu-config",
        "id": 5,
        "name": "菜单配置",
        "route": "/menuConfig"
      }
    ],
    "father_id": 0,
    "icon": "icon-config",
    "id": 2,
    "name": "系统配置",
    "route": "/systemConfig"
  }
]
```

> 源码地址：[https://github.com/U8tou/go-listtotree](https://github.com/azhengyongqin/golang-tree-menu)