package list

import (
	"fmt"
)

//not safe list
type List struct {
	data []interface{}
}

func New(data ...interface{}) *List {
	l := &List{}
	l.Append(data...)
	return l
}

//添加一个
func (l *List) Add(v interface{}) {
	l.data = append(l.data, v)
}

//附加一组
func (l *List) Append(v ...interface{}) {
	l.data = append(l.data, v...)
}

//数量
func (l *List) Count() int {
	return len(l.data)
}

//插入
func (l *List) Insert(index int, o interface{}) int {

	if index > len(l.data) {
		index = len(l.data)
	}
	if index < 0 {
		index = 0
	}
	var R []interface{}
	R = append(R, l.data[index:]...)
	l.data = append(l.data[:index], o)
	l.data = append(l.data, R...)
	return index
}

//合并
func (l *List) Concat(k *List) {
	l.data = append(l.data, k.RawList()...)
}

//深拷贝
func (l *List) DeepCopy(k *List) {
	l.data = append(l.data[0:0], k.RawList()...)
}

//按序号移除一个节点
func (l *List) Remove(index int) interface{} {
	if index < 0 || index >= len(l.data) {
		return nil
	}
	v := l.data[index]
	l.data = append(l.data[:index], l.data[index+1:]...)
	return v
}

type RuleFunc func(interface{}) bool

//移除一个相等的单位，返回索引，失败返回false
func (l *List) RemoveEquel(o interface{}) bool {
	for index := 0; index < len(l.data); index++ {
		v := l.data[index]
		if v == o {
			l.data = append(l.data[:index], l.data[index+1:]...)
			return true
		}
	}
	return false
}

//移除一个节点
func (l *List) RemoveRule(rule RuleFunc) interface{} {
	for index := 0; index < len(l.data); index++ {
		v := l.data[index]
		if rule(v) {
			l.data = append(l.data[:index], l.data[index+1:]...)
			return v
		}
	}
	return nil
}

//移除所有符合条件节点
func (l *List) RemoveAllRule(rule RuleFunc) int {
	var i, c int
	le := len(l.data)
	for {
		if i+c >= le {
			break
		}
		v := l.data[i]
		if rule(v) {
			l.data = append(l.data[:i], l.data[i+1:]...)
			c++
		} else {
			i++
		}
	}

	return c
}

//所有节点执行f函数
func (l *List) Each(f func(o interface{})) {
	for _, v := range l.data {
		f(v)
	}
}

func (l *List) EachElse(f func(o interface{}) bool) bool {
	var isbreak = false
	for _, v := range l.data {
		if f(v) {
			isbreak = true
			break
		}
	}
	return isbreak
}

//匹配函数
func (l *List) MatchPair(f func(o1, o2 interface{}) bool) (bool, interface{}, interface{}) {
	if l.Count() < 2 {
		return false, nil, nil
	}
	for _, v := range l.data {
		for _, v2 := range l.data {
			if v != v2 && f(v, v2) {
				return true, v, v2
			}
		}
	}
	return false, nil, nil
}

type PairInterface struct {
	I1 interface{}
	I2 interface{}
}

func findPairInterface(plist []PairInterface, o interface{}) bool {
	for _, p := range plist {
		if p.I1 == o || p.I2 == o {
			return true
		}
	}
	return false
}

//匹配函数 一次遍历,会将移除的
func (list *List) MatchPairList(f func(o1, o2 interface{}) bool) []PairInterface {
	if list.Count() < 2 {
		return nil
	}

	var plist []PairInterface
	//成功的设置nil
	for _, node := range list.data {
		for _, node2 := range list.data {
			if node != nil && node2 != nil && node != node2 {
				f1 := findPairInterface(plist, node)
				f2 := findPairInterface(plist, node2)
				if !f1 && !f2 && f(node, node2) {
					p := PairInterface{node, node2}
					plist = append(plist, p)
					fmt.Println("mathc:", p)
				}
			}
		}
	}
	//清空空位
	// if len(plist) > 0 {
	// 	for _, v := range plist {
	// 		list.RemoveEquel(v.I1)
	// 		list.RemoveEquel(v.I2)
	// 		fmt.Println("rem:", v)
	// 	}
	// }

	return plist
}

//按规则查找一个
func (l *List) Find(rule RuleFunc) interface{} {
	for _, v := range l.data {
		if rule(v) {
			return v
		}
	}
	return nil
}

//按规则查找所有
func (l *List) FindAll(rule RuleFunc) []interface{} {
	var tempL []interface{}
	for _, v := range l.data {
		if rule(v) {
			tempL = append(tempL, v)
		}
	}
	return tempL
}

//原始列表
func (l *List) RawList() []interface{} {
	return l.data
}

//清理
func (l *List) Clear() {
	l.data = nil
}

func (l *List) String() string {
	return fmt.Sprintf("%v", l.data)
}
