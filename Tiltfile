k8s_yaml('deployments/app-golang-secrets.yaml')

k8s_yaml('deployments/app-golang-postgres-deployment.yaml')
k8s_resource('app-golang-postgres', port_forwards=5432)


k8s_yaml('deployments/app-golang-todo-deployment.yaml')
k8s_resource('app-golang-todo', port_forwards=8080)