package main

import (
	"context"
	"flag"
	"fmt"
	k8s "github.com/ericchiang/k8s"
	k8sc "github.com/uswitch/k8sc"
	"os"
)

func main() {
	var kubeconfig = flag.String("kubeconfig", "", "Path to kube config, normally ~/.kube/config")
	flag.Parse()

	c, err := k8sc.NewClient(*kubeconfig)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}

	pods, err := c.CoreV1().ListPods(context.Background(), k8s.AllNamespaces)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}

	for _, p := range pods.Items {
		fmt.Printf("%s/%s\n", *p.Metadata.Namespace, *p.Metadata.Name)
	}
}
