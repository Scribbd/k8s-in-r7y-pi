---
tags: [exploration, ansible]
---

# Configuring the nodes automagicaly 
I was struck with a question when thinking about flashing new SD cards for  It is a luxury problem, I know. However, this document is to explore possible ways to automate this process.

Unlike the whole 'what to use for Kubernetes'-paralysis, I am pretty much set on Ansible. The only challenge for me is that I intend to use my Windows PC as the master server for this one. However, it should be possible to use it on [WSL](http://blog.rolpdog.com/2020/03/why-no-ansible-controller-for-windows.html).

Well, I will be using WSL in some way as Desktop Docker is using this.

## Installation
As mentioned before I am going to use Docker(WSL). So installation is done with a Dockerfile.

Link to current version: ![[ansible/Dockerfile]]

## Configuration
I am following [playlist by Jeff Geerling](https://www.youtube.com/playlist?list=PL2_OBreMn7FqZkvMYt6ATmgC0KAGGJNAN) tutorial and he suggested some good configurations to speed up ansible:
```bash
ANSIBLE_PIPELINING=True #This allows ansible to reuse the same ssh-connection
ANSIBLE_GATHERING=smart #This caches new machines automagicaly
ANSIBLE_RETRY_FILES_ENABLED=False #This prevents the creation of retry files
```

## File structure
I will be trying to adhere to the directory structure advised [here](https://docs.ansible.com/ansible/latest/user_guide/sample_setup.html#sample-directory-layout)

Ansible seems to use a set structure for playbook management. I could be mistaken.

## Still exploring?
No. This is going way beyond just exploration. And trying to stay in this one file gets a bit cramped. So for further reading and serious setting up an ansible playbook for my pi's, look here: [[Ansible]].

# Links
WSL blogpost: http://blog.rolpdog.com/2020/03/why-no-ansible-controller-for-windows.html
