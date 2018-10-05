package main

import (
	"fmt"
)

type(



  tree struct {

	  root *nd
  }
   nd struct {
	   k      int
	   leftC  *nd
	   rightC *nd
	   p *nd
   }
)

func  NewTree() *tree  {
	return &tree{
		nil,
	}
}
//添加
func(t *tree) insert(k int){
	if t.root==nil{
		newnd := &nd{k, nil, nil,nil }
		t.root=newnd
	}else {

		for demoKey := t.root; demoKey!=nil; {

			if k < demoKey.k {
				if demoKey.leftC == nil {
					newnd := &nd{k, nil, nil, demoKey}
					demoKey.leftC = newnd
					return
				}
				demoKey = demoKey.leftC
			}

			if k > demoKey.k {
				if demoKey.rightC == nil {
					newnd := &nd{k, nil, nil, demoKey}
					demoKey.rightC = newnd
					return
				}
				demoKey = demoKey.rightC
			}
		}

	}

}

//最小关键字
func (tr *tree)treeMINI() (mini *nd) {
	if tr.root!=nil {

		for mini=tr.root;mini.leftC!=nil;{
			mini=mini.leftC
		}
		return mini
	}
	return nil
}
//最大关键字
func (tr *tree)treeMAXI() (maxi *nd) {
	if tr.root!=nil {

		for maxi=tr.root;maxi.rightC!=nil;{
			maxi=maxi.leftC
		}
		return maxi
	}
	return nil
}

func ressort(root *nd){
	if root !=nil {
		ressort(root.leftC)
		fmt.Print(root.k, ",")
		ressort(root.rightC)
	}
}
//遍历排序
func (t *tree)sort(){
	ressort(t.root)
}

//查询
func (t *tree)search(key int)(demo *nd,meg interface{}) {
	if t.root==nil{
		return nil,"mou"
	}
	if t!=nil {
		for demo = t.root; demo != nil ; {
			if demo.k==key{
				return  demo ,"yes"
			}
			if demo.k > key {
				demo = demo.leftC
			} else {
				demo = demo.rightC
			}

		}

	}
	return nil,"not"


}
func main() {
	t:=NewTree()
	t.insert(4)
	t.insert(2)
	t.insert(3)
	t.insert(1)
	t.insert(8)
	t.insert(9)
	 mini:=t.treeMINI()

	 fmt.Println(mini)
	 t.sort()
	//_,msg:=t.search(4)
	//fmt.Printf("%s",msg)

}