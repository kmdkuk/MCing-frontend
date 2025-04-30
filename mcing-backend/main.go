package main

import (
	"encoding/json"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Kubernetesクライアントの初期化
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Dynamicクライアントの初期化
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic client: %v", err)
	}

	// カスタムリソースのGroupVersionResourceを定義
	minecraftGVR := schema.GroupVersionResource{
		Group:    "mcing.kmdkuk.com",
		Version:  "v1alpha1",
		Resource: "minecrafts",
	}

	// HTTPハンドラー
	http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
		// デフォルトのNamespaceのPodを取得
		pods, err := clientset.CoreV1().Pods("default").List(r.Context(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, "Failed to fetch pods", http.StatusInternalServerError)
			log.Printf("Error fetching pods: %v", err)
			return
		}

		// レスポンスをJSONで返す
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pods)
	})

	http.HandleFunc("/minecrafts", func(w http.ResponseWriter, r *http.Request) {
		// デフォルトのNamespaceのカスタムリソースを取得
		minecrafts, err := dynamicClient.Resource(minecraftGVR).Namespace("default").List(r.Context(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, "Failed to fetch minecrafts", http.StatusInternalServerError)
			log.Printf("Error fetching minecrafts: %v", err)
			return
		}
		// レスポンスをJSONで返す
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(minecrafts)
	})

	log.Println("Starting server on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
