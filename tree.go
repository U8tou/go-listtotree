package tree

import (
	"encoding/json"
	"log"
	"sort"
)

type TreeConfig struct {
	Rid        float64                // 根节点ID
	IdName     string                 // 节点名
	PidName    string                 // 父节点名
	ChiName    string                 // 子节点名
	WeightName string                 // 权重名
	Mapper     map[string]string      // 扩展字段,将原有的key映射新增成另一个key
	Extend     map[string]interface{} // 扩展字段,在原有的结构体上添加新的key
}

func GenerateTree(rows interface{}, conf *TreeConfig) (tree []map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	if conf.IdName == "" || conf.PidName == "" {
		log.Println("IdName or PidName cannot be empty")
		return tree
	}
	// 默认子节点的名称
	if conf.ChiName == "" {
		conf.ChiName = "children"
	}
	// 整理数据变成合适的List结构
	nodeData, err := buildNode(rows, conf)
	if err != nil {
		log.Printf("run time error: %v", err)
		return nil
	}
	// 构建树
	return buildRoot(nodeData, conf)
}

// buildNode List To Node
func buildNode(rows interface{}, conf *TreeConfig) ([]map[string]interface{}, error) {
	// 将切片转换为JSON字符串
	jsonSlice, err := json.Marshal(rows)
	if err != nil {
		return nil, err
	}
	// 转换
	var result []map[string]interface{}
	err = json.Unmarshal(jsonSlice, &result)
	if err != nil {
		return nil, err
	}
	// 扩展
	if conf.Mapper != nil || conf.Extend != nil {
		for _, mp := range result {
			if conf.Mapper != nil {
				for k, v := range conf.Mapper {
					mp[k] = mp[v]
				}
			}
			if conf.Extend != nil {
				for k, v := range conf.Extend {
					mp[k] = v
				}
			}
		}
	}
	return result, nil
}

// buildRoot 构建根节点
func buildRoot(list []map[string]interface{}, conf *TreeConfig) []map[string]interface{} {
	var c []map[string]interface{}
	for _, v := range list {
		if v[conf.PidName].(float64) == conf.Rid {
			children := buildChildren(list, v[conf.IdName].(float64), conf)
			if len(children) != 0 {
				v[conf.ChiName] = children
			}
			c = append(c, v)
		}
	}
	// 排序
	if conf.WeightName != "" {
		sort.Slice(c, func(i, j int) bool {
			a, aOk := c[i][conf.WeightName].(float64)
			b, bOk := c[j][conf.WeightName].(float64)
			if aOk && bOk {
				return a > b
			}
			return false
		})
	}
	return c
}

// buildRoot 构建子节点
func buildChildren(list []map[string]interface{}, id float64, conf *TreeConfig) []map[string]interface{} {
	var c []map[string]interface{}
	for _, v := range list {
		if v[conf.PidName].(float64) == id {
			children := buildChildren(list, v[conf.IdName].(float64), conf)
			if len(children) != 0 {
				v[conf.ChiName] = children
			}
			c = append(c, v)
		}
	}
	// 排序
	if conf.WeightName != "" {
		sort.Slice(c, func(i, j int) bool {
			a, aOk := c[i][conf.WeightName].(float64)
			b, bOk := c[j][conf.WeightName].(float64)
			if aOk && bOk {
				return a > b
			}
			return false
		})
	}
	return c
}
