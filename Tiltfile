load('ext://restart_process', 'docker_build_with_restart')

BACKEND_DOCKERFILE = '''FROM golang:alpine
WORKDIR /
COPY ./mcing-backend/bin/mcing-backend /
CMD ["/mcing-backend"]
'''

def backend_binary():
    return 'cd mcing-backend;mkdir -p bin;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/mcing-backend main.go'

local_resource('Watch & Compile (mcing backend)', backend_binary(), deps=['mcing-backend/main.go'])
docker_build_with_restart(
    'mcing-backend:latest', '.',
    dockerfile_contents=BACKEND_DOCKERFILE,
    entrypoint=['/mcing-backend'],
    only=['./mcing-backend/bin/mcing-backend'],
    live_update=[
        sync('./mcing-backend/bin/mcing-backend', '/mcing-backend'),
    ]
)
# Dockerイメージをビルド
docker_build('mcing-frontend', './mcing-frontend')

# Kubernetesマニフェストを適用
k8s_yaml('k8s/mcing-frontend.yaml')
k8s_yaml('k8s/mcing-backend.yaml')

# ポートフォワーディング
k8s_resource('mcing-frontend', port_forwards='8080:80')
k8s_resource('mcing-backend', port_forwards='8081:80')
