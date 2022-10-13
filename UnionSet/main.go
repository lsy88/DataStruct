package main

import "fmt"

//并查集
type UnionSet struct {
	Father []int
}

//初始化
func (u UnionSet) Init(n int) []int {
	//var father [math.MaxInt]int
	for i := 1; i <= n; i++ {
		u.Father[i] = i
	}
	return u.Father
}

//合并
func (u UnionSet) Union(i int, j int) {
	iFather := u.Father[i]      //找到i的祖先节点
	jFather := u.Father[j]      //找到j的祖先节点
	u.Father[iFather] = jFather //i的祖先指向j的祖先
}

//查找
func (u UnionSet) Find(i int) int {
	if i == u.Father[i] {
		return i
	} else {
		u.Father[i] = u.Find(u.Father[i]) //进行路径压缩
		return u.Father[i]                //返回父节点
	}
}

/*
	现在有若干家族图谱关系，给出了一些亲戚关系，如marry和tom是亲戚，tom和ben是亲戚，
	第一部分是以M,N开始，N为人数，这些人编号为1,2,3...N,下面有M行，每行有两个数a,b。表示a和b是亲戚
	第二部分以Q开始，以下Q行有Q个询问，每行为c,d，表示询问c,d是否是亲戚
*/
func main() {
	tmp := [][]int{{10, 7}, {2, 4}, {5, 7}, {1, 3}, {8, 9}, {1, 2}, {5, 6}, {2, 3}, {3}, {3, 4}, {7, 10}, {8, 9}}
	unionset := UnionSet{
		Father: make([]int, tmp[0][0]+1),
	}
	unionset.Init(tmp[0][0])
	for i := 1; i <= tmp[0][1]; i++ {
		unionset.Union(tmp[i][0], tmp[i][1]) //将关系压入并查集
	}
	if len(tmp[tmp[0][1]+1]) != 1 {
		return
	} //{3}问询的数量
	for i := tmp[0][1] + 2; i < len(tmp); i++ {
		if unionset.Find(tmp[i][0]) == unionset.Find(tmp[i][1]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
