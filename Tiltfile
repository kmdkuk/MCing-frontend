# Dockerイメージをビルド
docker_build('mcing-frontend', './mcing-frontend')

# Kubernetesマニフェストを適用
k8s_yaml('k8s/deployment.yaml')

# ポートフォワーディング
k8s_resource('mcing-frontend', port_forwards='8080:80')
