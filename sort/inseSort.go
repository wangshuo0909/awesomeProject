package main

import (
	"time"
	"fmt"
	"math/rand"
)
func InseSort(arr []int) {
	for i:=1;i< len(arr);i++  {
		n:=i-1
		key := arr[i]
		for ;n>=0&&arr[n]>key ;n--  {
			arr[n+1]=arr[n]
		}
		arr[n+1]=key
	}

}
func mergeSort(arr []int) []int{
	if len(arr)==0||len(arr)==1{
		return arr
	}
	harf:=len(arr)/2
	arrL:=mergeSort(arr[:harf])
	arrR:=mergeSort(arr[harf:])
	arrM:=merge(arrL,arrR)
	return arrM
}
func merge(arrLS []int ,arrRS []int)[]int {
	 i:=0
	 j:=0 
	 const INT_MAX=int(^uint(0)>>1)
	 arrL:=make([]int, len(arrLS))
	 copy(arrL, arrLS)
	 arrR:=make([]int, len(arrRS))
	 copy(arrR, arrRS)
	 arrL=append(arrL,INT_MAX)
	 arrR=append(arrR,INT_MAX)
	 arrM:=[]int{}

	 for i<len(arrL)-1||j<len(arrR)-1{
		if arrL[i]<arrR[j]{
			arrM=append(arrM,arrL[i])
			i++
		}else{
			arrM=append(arrM,arrR[j])
			j++
		}
	 }
	 return arrM
}
 //快速排序

 func QuickSort(arr []int ,start int ,end int ){
	
	if start<end&&arr!=nil{
			p:=Pritition(arr,start,end)
			QuickSort(arr,start,p-1)
			QuickSort(arr,p+1,end)
	}

 }
func Pritition (arr []int,start int ,end int) int {
	x:=arr[end]
	i:=start-1
	for j:=start;j<end;j++{
		if arr[j]<=x{
			i++
			tmp:=arr[i]
			arr[i]=arr[j]
			arr[j]=tmp
		}
	}
	arr[end]=arr[i+1]
	arr[i+1]=x
	return i+1
}
func main() {
	arr:=[]int{}
	for i:=0;i<100000;i++{
		arr=append(arr, rand.Intn(100000))
	}
	t1:=time.Now().UnixNano()
	mergeSort(arr)
	// fmt.Println(arr)
	t2:=time.Now().UnixNano()
	QuickSort(arr, 0, len(arr)-1)
	t3:=time.Now().UnixNano()

	// fmt.Println(arr)
	fmt.Printf("%d<>%d", t2-t1, t3-t2)

}
