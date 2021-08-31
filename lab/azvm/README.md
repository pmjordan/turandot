1. Create a VM on Azure using

`./lab/azvm/pj_azure_vm.sh create`

this initialises the VM using pj_cloud-init.txt

        (Hint cloud-init debug commands below )

            `cloud-init status --long`
            `sudo cat /var/log/cloud-init-output.log`
            

which installs a named commit of pmjordan/turandot from github

and connects to the VM. Wait for the the message on the VM that indicates the install process has finished.

2. Log on a second time using

`./lab/azvm/pj_azure_vm.sh connect`

3. Start minikube using the steps in pj_turandot_minikube_start.sh

4. When finished, tidy up using

`./lab/azvm/pj_azure_vm.sh destroy`
