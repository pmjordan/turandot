#!/bin/sh
# installs packages to ubuntu like my local machine. Consider also changing equivalent file used for azure centos.



# install helm
cd /tmp
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

# install canvas
cd /opt
sudo git clone https://github.com/pmjordan/oda-canvas-charts

# install turandot on ubuntu. Assumes required published packages already installed by cloud-ini


#Install turandot binary
cd /tmp
wget -O turandot.deb https://github.com/tliron/turandot/releases/download/v0.5.0/turandot_0.5.0_linux_amd64.deb
sudo apt install ./turandot.deb

wget -O reposure.deb https://github.com/tliron/reposure/releases/download/v0.1.4/reposure_0.1.4_linux_amd64.deb
sudo apt install ./reposure.deb

#Install puccini binary
wget -O puccini.deb https://github.com/tliron/puccini/releases/download/v0.18.0/puccini_0.18.0_linux_amd64.deb
sudo apt install ./puccini.deb

#Install kubectl and minikube here as can't get them from cloud-init packages
# install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# install minikube 
cd /tmp
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
dpkg -i minikube_latest_amd64.deb

sudo wall -n "Completed turandot tools installation. Start a new session to use new permissions and cd to /opt/turandot"


