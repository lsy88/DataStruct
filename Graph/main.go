package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"sync"
)

/*
	G(V,E)是有一个非空的有限顶点集合V和一个有限边集和E组成
	邻接矩阵G[N][N]--N个顶点从0到N-1编号
	G[i][j] = 1 则是图中的边
			= 0 不是边

	邻接表G[N]为指针数组，对应矩阵每行一个链表，只存非0元素，节约稀疏图的空间
*/

/*
	图的连通
	强连通：有向图中顶点v和w之间存在双向路径，则称v和w时强连通的
	强连通图：有向图中任意两顶点均强连通
	强连通分量：有向图的极大强连通子图
*/

type GraphInterface interface {
	Create() Graph //建立并返回空图
	//InsertVertex(v Vertex) //将v插入图
	//InsertEdge(e Edge)     //将e插入图
	//DFS(v Vertex)          //从顶点v出发深度优先遍历
	//BFS(v Vertex)          //广度优先遍历
	MST() //图G的最小生成树
}

//DFS原理
//func DFS(v Vertex) {
//	visited[v] = true
//	for v的邻接点W {
//		if !visited[w] {
//			DFS(w)
//		}
//	}
//}

//BFS原理和树的层序遍历相似
//func BFS(v Vertex) {
//	visited[v] = true
//	Enqueue(v, Q)   //把v入队列
//	for !IsEmpty(Q) {
//		v = Dequeue(Q)   //出队
//		for v的每个邻接点w {
//			if !visited[w] {
//				visited[w] = true
//				Enqueue(w, Q)  //将邻接点入队
//			}
//		}
//	}
//}

type Graph struct {
	Nodes []*Node          //节点集
	Edge  map[Node][]*Node //邻接表表示的无向图
	lock  sync.RWMutex     //保证线程安全
}

type Node struct {
	value generic.Type
}

//增加节点
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.Nodes = append(g.Nodes, n)
}

//增加边
func (g *Graph) AddEdge(u, v *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	//首次建立图
	if g.Edge == nil {
		g.Edge = make(map[Node][]*Node)
	}
	g.Edge[*u] = append(g.Edge[*u], v) //建立u->v的边
	g.Edge[*v] = append(g.Edge[*v], u) //无向图同时存在v->u的边
}

//输出图
func (g *Graph) String() {
	g.lock.Lock()
	defer g.lock.Unlock()
	
	str := ""
	for _, v := range g.Nodes {
		str += v.String() + "->"
		nexts := g.Edge[*v]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

//BFS 广度优先遍历,是从一个节点发散性的遍历周围节点，从某个节点出发，访问她所有的邻接节点，再从这些节点出发，访问未被访问过的节点
//类似树的层序遍历，但图存在成环的情形，访问过的节点可能会再次访问，所以需要用一个辅助队列来存放待访问的邻接节点

//队列
type NodeQueue struct {
	Nodes []Node
	lock  sync.RWMutex
}

//初始化队列
func NewNodeQueue() *NodeQueue {
	q := NodeQueue{}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Nodes = []Node{}
	return &q
}

//入队列
func (q *NodeQueue) Enqueue(n Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Nodes = append(q.Nodes, n)
}

//出队列
func (q *NodeQueue) Dequeue() *Node {
	q.lock.Lock()
	defer q.lock.Unlock()
	
	node := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return &node
}

func (q *NodeQueue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.Nodes) == 0
}

func (g *Graph) BFS(f func(node *Node)) {
	g.lock.Lock()
	defer g.lock.Unlock()
	
	//初始化队列
	q := NewNodeQueue()
	//取图的第一个节点入队列
	head := g.Nodes[0]
	q.Enqueue(*head)
	//标识节点是否已经被访问过
	visited := make(map[*Node]bool)
	visited[head] = true
	
	//遍历所有节点直到队列为空
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		//visited[node] = true
		nexts := g.Edge[*node]
		
		//将所有未被访问过的邻接节点入队列
		for _, next := range nexts {
			//如果节点已经被访问过
			if visited[next] {
				continue
			}
			q.Enqueue(*next)
			visited[next] = true
		}
		
		//对每个正在遍历的节点执行回调
		if f != nil {
			f(node)
		}
	}
}
