k8s_yaml('kubernetes/app-golang-secrets.yaml')

k8s_yaml('kubernetes/app-todolist-db-task.yaml')
k8s_resource('app-todolist-db-task', port_forwards=5432)

k8s_yaml('kubernetes/app-todolist-db-user.yaml')
k8s_resource('app-todolist-db-user', port_forwards=5433)

docker_build('go-todo_app-api-gw', 'api-gw', dockerfile='api-gw/build/Dockerfile')
k8s_yaml('kubernetes/app-todolist-api-gw.yaml')
k8s_resource('app-todolist-api-gw', port_forwards=8080)

docker_build('go-todo_app-user', 'user', dockerfile='user/build/Dockerfile')
k8s_yaml('kubernetes/app-todolist-user.yaml')
k8s_resource('app-todolist-user', port_forwards=8081)

docker_build('go-todo_app-task', 'task', dockerfile='task/build/Dockerfile')
k8s_yaml('kubernetes/app-todolist-task.yaml')
k8s_resource('app-todolist-task', port_forwards=8082)