ODA Example
===================

An adaption of the Hello World example with various in-process nodes working 
towards an implementation of a component as defined by TMForum


Building the CSAR
-----------------

* [Package as CSAR file](scripts/build-csar)


Deploying
---------

    turandot service deploy oda --file=dist/oda.csar

If you want to access the deployed web server from outside the cluster you will need to have
loadbalancing supported on your Kubernetes cluster. On Minikube you can just
[start a tunnel](https://minikube.sigs.k8s.io/docs/handbook/accessing/#using-minikube-tunnel).

If supported, the "url" output of the service template will work. To open from your default web
browser:

    xdg-open $(turandot service output oda url)
