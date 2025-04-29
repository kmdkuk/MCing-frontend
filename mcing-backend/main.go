package main

import (
	"encoding/json"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func main() {
	http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
		pods := corev1.PodList{
			Items: []corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "example-pod",
						Namespace: "default",
					},
				},
			},
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pods)
	})

	http.ListenAndServe(":80", nil)
}
