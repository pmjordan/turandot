minikube start --addons=registry ...
kubectl config set-context --current --namespace=workspace
turandot operator install --site=central --role=view --wait -v
reposure registry create default --provider=minikube --wait -v
minikube status
reposure registry list


