package main

import (
	"context"
	"fmt"
	"time"

//	"golang.org/x/tools/go/analysis/passes/nilfunc"
//	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
  //Create the in cluster config
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
  }

  //Create the clientset
  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  for {
    //Get all the pods in all namespaces
    pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
      panic(err.Error())
    }

    fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))


		time.Sleep(10 * time.Second)
  }
}
