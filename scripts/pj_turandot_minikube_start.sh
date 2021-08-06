# ingress addon is required by canvas operator, insecure-registry is required by cert-manager quicker to start them here
minikube delete
minikube start --addons registry ingress --insecure-registry "10.0.0.0/8","192.168.0.0/16"
#install canvas
pushd
cd /opt/oda-canvas-charts
kubectl create namespace canvas
kubectl config set-context --current --namespace=canvas
./install_canvas.sh
#install turandot
cd /opt/turandot/
kubectl config set-context --current --namespace=workspace
turandot operator install --role=view --site=central --wait -vv
reposure operator install --role=view --wait -v
reposure registry create default --provider=minikube --wait -v
kubectl apply --filename=assets/kubernetes/cert-manager.yaml
reposure simple install --authentication --wait -v
reposure registry create default --provider=simple --wait -v
minikube status
reposure registry list
popd



