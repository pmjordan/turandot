#!/bin/sh
# installs packages to ubuntu like my local machine. Consider also changing equivalent file used for azure centos.
# install helm
cd /tmp
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

# install canvas
cd /opt
sudo git clone https://github.com/tmforum-oda/oda-canvas-charts

# install turandot on Centos. Assumes required published packages already installed by cloud-init
echo "this has been written by turandot script pj_install.sh " + $(date) >> ~/pj_turandot_install_logs.txt


#Install turandot binary
cd /tmp
wget -O turandot.deb https://github.com/tliron/turandot/releases/download/v0.5.0/turandot_0.5.0_linux_amd64.deb
sudo apt install ./turandot.deb

wget -O reposure.deb https://github.com/tliron/reposure/releases/download/v0.1.4/reposure_0.1.4_linux_amd64.deb
sudo apt install ./reposure.deb

#Install puccini binary
# pin to version 15 as higher version require glibc v 2.32 which is not available in OS images available in azure
wget -O puccini.deb https://github.com/tliron/puccini/releases/download/v0.15.0/puccini_0.15.0_linux_amd64.deb
sudo apt install ./puccini.deb

#Install kubectl and minikube here as can't get them from cloud-init packages
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
sudo chmod 755 /usr/local/bin/minikube

wget https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo cp minikube-linux-amd64 /usr/local/bin/minikube

sudo wall -n "Completed turandot tools installation. Start a new session to use new permissions and cd to /opt/turandot"


