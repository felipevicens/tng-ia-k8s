package api

import (
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clientset = config()

func GetPods(w http.ResponseWriter, r *http.Request) {
	// Get pods -- package metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fmt.Println(r.Header)
	podslist, _ := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	w.Write([]byte("Pod List:"))
	for _, p := range podslist.Items {
		w.Write([]byte(p.GetName()))
		//fmt.Println("\t", p.GetName())
	}
	//fmt.Fprintf(w, "Pods, %q", html.EscapeString(r.URL.Path))
}

func GetNodes(w http.ResponseWriter, r *http.Request) {
	// Get nodes -- package metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fmt.Println(w.Header)
	nodeslist, _ := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	w.Write([]byte("Node List:"))
	for _, n := range nodeslist.Items {

		w.Write([]byte(n.GetName()))
		//fmt.Println("\t", n.GetName())
	}

	//fmt.Fprintf(w, "Nodes, %q", html.EscapeString(r.URL.Path))
}
