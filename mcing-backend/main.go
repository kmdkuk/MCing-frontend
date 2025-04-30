package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	log.Println("Starting server on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

// ホームディレクトリを取得
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windowsの場合
}
