package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"tng-ia-k8s/api"

	"github.com/gorilla/mux"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/nodes", api.GetNodes).Methods("GET")
	router.HandleFunc("/pods", api.GetPods).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))

		nodename, err := clientset.CoreV1().Nodes().Get("minikube", metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("The node name is: %v. \n", (nodename.Name))

		// Get nodes -- package metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		fmt.Println("Node List:")
		nodeslist, _ := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		for _, n := range nodeslist.Items {
			fmt.Println("\t", n.GetName())
		}

		// Get pods -- package metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		fmt.Println("Pod List:")
		podslist, _ := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		for _, p := range podslist.Items {
			fmt.Println("\t", p.GetName())
		}

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		pod := "redis-bd48f75f-d5bmv"
		_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(4 * time.Second)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
