---
tags: [ansible, playbook]
aliases: [ansible playbooks]
---
# Playbook Index
This is an index of playbooks in order of operation. To get a working K3s cluster going, run these playbooks in order of appearance. From a newly flashed Raspian image, to k3s master / worker node.

There will probably be a master playbook in the future that will use roles and the like. For now just run the playbooks in order.

Run them as follows:
 - Your workdir has to be ./src as the ansible.cnf has connecting information in it.
 - Use `ssh-add` to add your configured pubkey to your agent. Not doing this will cause timeout (UNREACHABLE) issues when running on multiplen nodes
 - `ansible-playbook -i inventory [playbook-path]`


Handy commands for ansible:
 - `ansible -i inventory all -a '/sbin/reboot' --become -u ansible-controller` This will return failed messages, however, they do succeed in rebooting the sytem.

## Inventory file
This is my current inventory file for future reference:
```ini
[master_nodes]
pi1.cluster.lan ansible_host=192.168.178.150

[worker_nodes]
pi2.cluster.lan ansible_host=192.168.178.160
pi3.cluster.lan ansible_host=192.168.178.161

[cluster:children]
master_nodes
worker_nodes

[cluster:vars]
ansible_python_interpreter=/usr/bin/python3
master_ip=192.168.178.150
k3s_version=v1.21.4+k3s1
```

## fresh_pi
This is a playbook for the freshly installed Raspbian image. It has a specific Raspbian specific task for disabling the wifi and Bluetooth on boot. After this it will do some SSH specific tasks for later use. And creates and configure users that only uses the pubkey authentication method.

Note: Run this inside the freshpi folder as it will not work with the ansible.cnf file in the root src folder.

Summary of tasks:
 - Disable wifi and Bluetooth
 - Change SSH port
 - Setup automated ansible and user access with associated pubkeys for later use

## First 5
This playbook follows the general guide on hardening linux distro's. This is more of a general playbook. However, this is made specifically for Raspbian in mind and has no automagic on other distro's.

Summary of tasks:
 - Remove the default `pi` user
 - Setup hostnames and install `/etc/hosts` template
 - Install and configure `unattended-updates`
 - Install and configure `fail2ban`

## Prehorn
This playbook is very specific for Raspbian and Raspberry Pi's and my hardware configuration. It will setup an external volume for use with [[Longhorn]].
It will do some low-level stuff that might brick your USB-to-Sata adapter, or anything that isn't that on the /dev/sda path.

This playbook is normally configured to only run on the worker nodes. This is due to my [[01_Docker in Cluster Project#Available hardware]]. See [[Longhorn#How will it fit in my cluster]] for more info.

Summary of tasks:
 - Creates GPT table and primary partition
 - Formats primary partition with ext4
 - Adds fstab and mounts drive to `/var/lib/longhorn` as the default path for longhorn
 - Configures TRIM on the system for the volume, when drive supports it. [^funnystory]

[^funnystory]: So I got a funny story to tell about this. I spend a whole day trying to figure out why one adapter did support TRIM and why the other didn't. Well, it appeared it hadn't anything to do with the adapters I used. They supported TRIM just fine. However, it were the drives I bought which were the culprit of my troubles. Which made the webshop that sold them to me the source of my annoyance. You see, I ordered two of the same (re-used) drives. But I got two different ones. A X300 and a X300s... And yes, they are different as the s-variant supports TRIM and the other doesn't. -_- sigh.

## K8s_install
Installs K3s on the nodes. This playbook is mostly stolen/inspired from [rancher's ansible playbook](https://github.com/k3s-io/k3s-ansible). The original playbook will install k3s in a smarter and distro-agnostic way. I just de-roled it, and just grabbed the tasks that applies to my cluster hardware. This to get more a grasp on what the playbook does.

Summary of tasks:
 - Setup some Raspian specific configurations
 - Sets up the master node with access keys and retrieves node token
 - Sets up worker nodes with node token

## K8s_services
This playbook is for installing several services on the cluster. Services to be added are:
 - Kubernetes Dashboard
	 - Provides a way to manage the K8s cluster through a web console
 - Longhorn
	 - Persistent distributed storage
 - Ansible AWX
	 - Ansible Tower, but the community kind. It will make it possible to automate the K3s cluster from the inside

Longhorn is a solution for Persistent Volume Storage that runs on and for Kubernetes.

> From here on out k8s will be used in ansible. Note to myself, here is a link to the docs on it: https://docs.ansible.com/ansible/latest/collections/kubernetes/core/index.html#plugins-in-kubernetes-core

Summary of tasks:
 - Installing dependencies needed by Longhorn. Of note:
   - open-scsi
   - nfsv4
 - Installing Longhorn through kubectl
 - Installing Kubernetes Dashboard through kubectl
 - Installing Ansible AWX through AWX operator