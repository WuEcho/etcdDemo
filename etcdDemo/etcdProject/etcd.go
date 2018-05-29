package etcdProject

import (
	"github.com/coreos/etcd/clientv3"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"context"
	"time"
)

func CreatCli(endPoint []string,dialTimeout int) (* clientv3.Client,error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endPoint,
		DialTimeout: time.Duration(dialTimeout) * time.Second,
	})


	if err != nil{
		return nil,err
	}
	return cli,err
}

func Close(cli * clientv3.Client)  {
	//关闭
	 cli.Close()
}


//插入数据
func PutValue(cli *clientv3.Client,key ,value string) bool {
	_,err := cli.Put(context.TODO(),key,value)
	if err != nil{
		return false
	}
	return true
}


//查
func GetValue(cli *clientv3.Client, key string)[] *mvccpb.KeyValue {
	resp,err := cli.Get(context.TODO(),key)
	if err != nil {
		fmt.Println(err)
	}else {
		//返回查询结果
		return resp.Kvs
	}
	return nil
}

//删除
func DelValue(cli *clientv3.Client,key string) int64 {
	delRes,err := cli.Delete(context.TODO(),key)
	if err != nil {
		fmt.Println(err)
	}else {
		return delRes.Deleted
	}
	return 0
}

//按前缀删除某些数据
func DelValueByPrefix(cli *clientv3.Client,key string) int64 {

	resp,err := cli.Delete(context.TODO(),key,clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	}else {
		return resp.Deleted
	}
	return 0
}

//按照前缀获取数据
func GetValueByPrefix(cli *clientv3.Client,key string) []*mvccpb.KeyValue {

	resp,err := cli.Get(context.TODO(),key,clientv3.WithPrefix())

	if err != nil {
		fmt.Println(err)
	}else {
		return resp.Kvs
	}
	return nil
}

//事务
func OfficeDeal(cli *clientv3.Client,key, value, otherValue string) bool {
	_,err1 := cli.Txn(context.TODO()).
		If(clientv3.Compare(clientv3.Value(key),"=",value)).
		Then(clientv3.OpPut(key,value)).
		Else(clientv3.OpPut(key,otherValue)).Commit()

	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	return true
}

//监看
func WatchKey(cli *clientv3.Client,key string) []byte{
	rch := cli.Watch(context.Background(),key)
	for wresp := range rch{
		for _,ev := range wresp.Events{
			fmt.Printf("watch foo ----%s  %q : %q \n",ev.Type,ev.Kv.Key,ev.Kv.Value)
			return ev.Kv.Value
		}
	}
	return nil
}

//监看带有某种前缀的键
func WatchKeyWithPrefix(cli *clientv3.Client,prefixKey string) []byte {
	rch := cli.Watch(context.Background(),prefixKey,clientv3.WithPrefix())
	for wresp := range rch{
		for _,ev := range wresp.Events{
			fmt.Printf("watch foo ----%s  %q : %q \n",ev.Type,ev.Kv.Key,ev.Kv.Value)
			return ev.Kv.Value
		}
	}
	return nil
}


//监看在某一范围的键
func WatchKeyWithRang(cli *clientv3.Client,beginKey,endKey string) []byte {

	rch := cli.Watch(context.Background(),beginKey,clientv3.WithRange(endKey))

	for wresp := range rch{
		for _,ev := range wresp.Events{
			fmt.Printf("watch foo ----%s  %q : %q \n",ev.Type,ev.Kv.Key,ev.Kv.Value)
			return ev.Kv.Value
		}
	}
	return nil
}


//周期
func CycleSetKeyValueWithTime(cli *clientv3.Client,timeOut int64,tempKey,tempValue string) bool {

	resp,err := cli.Grant(context.TODO(),timeOut)
	if err != nil{
		fmt.Println(err)
		return false
	}

	_,err = cli.Put(context.TODO(),tempKey,tempValue,clientv3.WithLease(resp.ID))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}