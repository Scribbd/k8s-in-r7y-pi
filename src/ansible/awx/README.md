# Readme on AWX

An AWX image needs to build within the docker image as it requires ansible.

[Link to awx docker hub](https://quay.io/repository/ansible/awx?tab=info)

At first I tried building my own image. However, it seems that the install process has been altered and active documentation isn't applicable anymore. I managed to run my own instance by following the next [tutorial](https://computingforgeeks.com/how-to-install-ansible-awx-on-ubuntu-linux/). I used Docker Desktop in Kubernetes mode and used WSL2 for the one command that could be done in windows/Visual studio code. *Update*: On windows this tutorial causes a bootloop in awx.


