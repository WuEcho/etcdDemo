package main

import (
	"0528/etcdDemo/etcdProject"
	"fmt"
    "time"
)


func main()  {

  etcd,err := etcdProject.CreatCli([]string{"127.0.0.1:2379"},2)

  defer etcd.Close()

  if err != nil {
     fmt.Println(err)
  }

  ptr := etcdProject.PutValue(etcd,"abc","123")

  fmt.Println("存放是否成功",ptr)

  gtr := etcdProject.GetValue(etcd,"abc")

  fmt.Println("查询结果",gtr)

  etcdProject.DelValue(etcd,"abc")

  gtr = etcdProject.GetValue(etcd,"abc")

  fmt.Println("查询结果",gtr)

  ptr = etcdProject.PutValue(etcd,"abcdef","1234")

  gtbP := etcdProject.GetValueByPrefix(etcd,"abcd")

  fmt.Println("查询结果",gtbP)

  dt := etcdProject.DelValueByPrefix(etcd,"abc")

  fmt.Println("删除是否成功",dt)

  gtr = etcdProject.GetValue(etcd,"abcdef")

  fmt.Println("查询结果",gtr)

  etcdProject.PutValue(etcd,"abc","123")

  ofDeal := etcdProject.OfficeDeal(etcd,"abc","123","456")

  fmt.Println("事务是否成功",ofDeal)

  gtr = etcdProject.GetValue(etcd,"abc")

  fmt.Println("查询结果",gtr)

  ofDeal = etcdProject.OfficeDeal(etcd,"abc","456","345")

  fmt.Println("事务是否成功",ofDeal)

  gtr = etcdProject.GetValue(etcd,"abc")

  fmt.Println("查询结果",gtr)

  var cha  chan bool
  var chan1 chan bool


  go func() {
    m :=  etcdProject.WatchKey(etcd,"abc")
    fmt.Println("监看键值",string(m))
    cha<-true
  }()

  go func() {
    i := etcdProject.WatchKeyWithPrefix(etcd,"a")
    fmt.Println("监看带有前缀的键",string(i))
    chan1<-true
  }()


  etcdProject.PutValue(etcd,"abc","1234")

  etcdProject.PutValue(etcd,"abc","121")

  cySetkv := etcdProject.CycleSetKeyValueWithTime(etcd,10,"haha","balbal")

  fmt.Println("周期是否成功",cySetkv)

  gv := etcdProject.GetValue(etcd,"haha")

  fmt.Println("周期结束之前",gv)

  time.Sleep(11 * time.Second)

  gv = etcdProject.GetValue(etcd,"haha")

  fmt.Println("周期结束之后",gv)

  <-cha
  <-chan1
}
