package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var ns, label, field string
	flag.StringVar(&ns, "namespace", "", "Namespace")
	flag.StringVar(&label, "l", "", "Label Selector")
	flag.StringVar(&field, "f", "", "Field Selector")
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	api := clientset.CoreV1()
	//Setup List options
	listOptions := metav1.ListOptions{
		LabelSelector: label,
		FieldSelector: field,
	}
	pvcs, err := api.Pods(ns).List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	printPods(pvcs)

}
func printPods(pvcs *v1.PodList) {
	template := "%-32s%-8s%-12s%-10s\n"
	fmt.Printf(template, "NAME", "STATUS", "Namespace", "NODE")
	for _, pvc := range pvcs.Items {
		quant := pvc.Spec.NodeName
		fmt.Printf(template, pvc.Name, string(pvc.Status.Phase), pvc.GetNamespace(), quant)
	}

}
