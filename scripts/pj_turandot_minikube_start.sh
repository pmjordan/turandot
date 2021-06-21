# ingress addon is required by canvas operator, quicker to start it here
minikube start --addons registry ingress --insecure-registry "192.168.0.0/16"
#minikube start --addons=registry ...
#install canvas
cd /opt/oda-canvas-charts
kubectl create namespace canvas
./install_canvas.sh
cd ~
#install turandot
kubectl config set-context --current --namespace=workspace
turandot operator install --role=view --site=central --wait -vv
reposure operator install --role=view --wait -v
reposure registry create default --provider=minikube --wait -v
kubectl apply --filename=assets/kubernetes/cert-manager.yaml
reposure simple install --authentication --wait -v
reposure registry create default --provider=simple --wait -v
minikube status
reposure registry list



