# ingress addon is required by canvas operatpr, quicker to start it here
minikube start --addons registry ingress 
#minikube start --addons=registry ...
kubectl config set-context --current --namespace=workspace
turandot operator install --role=view --site=central --wait -vv
reposure operator install --role=view --wait -v
reposure registry create default --provider=minikube --wait -v
kubectl apply --filename=https://github.com/jetstack/cert-manager/releases/download/v1.3.1/cert-manager.yaml
reposure simple install --authentication --wait -v
reposure registry create default --provider=simple --wait -v
minikube status
reposure registry list



