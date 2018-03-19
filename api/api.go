package api

import (
	"fmt"
	"net/http"

	"k8s.io/client-go/kubernetes"
)

func GetPods(w http.ResponseWriter, r *http.Request) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// Get pods -- package metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fmt.Println("Pod List:")
	podslist, _ := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	for _, p := range podslist.Items {
		fmt.Println("\t", p.GetName())
	}
	//fmt.Fprintf(w, "Pods, %q", html.EscapeString(r.URL.Path))
}

func GetNodes(w http.ResponseWriter, r *http.Request) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// Get nodes -- package metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fmt.Println("Node List:")
	nodeslist, _ := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	for _, n := range nodeslist.Items {
		fmt.Println("\t", n.GetName())
	}

	//fmt.Fprintf(w, "Nodes, %q", html.EscapeString(r.URL.Path))
}
