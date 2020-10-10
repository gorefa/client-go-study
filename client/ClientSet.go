package client

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/lisai/.kube/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	list, err := podClient.List(context.Background(), metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf("NAMESPACE: %v \t NAME: %v \t  STATUS: %+v \n", d.Namespace, d.Name, d.Status.Phase)
	}

}
