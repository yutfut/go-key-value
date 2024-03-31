package riak

import (
	"fmt"
	"os"
	riak "github.com/basho/riak-go-client"
)

func RiakConn() {
	nodeOpts := &riak.NodeOptions{
		RemoteAddress: "127.0.0.1:8087",
	}

	var node *riak.Node
	var err error
	if node, err = riak.NewNode(nodeOpts); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nodes := []*riak.Node{node}
	opts := &riak.ClusterOptions{
		Nodes: nodes,
	}

	cluster, err := riak.NewCluster(opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func() {
		if err := cluster.Stop(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	if err := cluster.Start(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	obj := &riak.Object{
		ContentType:     "text/plain",
		Charset:         "utf-8",
		ContentEncoding: "utf-8",
		Value:           []byte("this is a value in Riak"),
	}

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithBucket("test").
		WithContent(obj).
		Build()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := cluster.Execute(cmd); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	svc := cmd.(*riak.StoreValueCommand)
	rsp := svc.Response
	fmt.Println(rsp.GeneratedKey)
}