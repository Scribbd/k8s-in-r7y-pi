# README Golang

A simple excersize in Go and docker / minikube deployment.

## GoHello
Go hello is a simple webserver that replies the path you requested from it. Be carefull, tho, this one is devious.

## GoHello in MiniKube
To use this with Ansible, I need to install minikube in my remote container... Be right back.

What I learned from Alpine linux. It really comes with the bare minimum. Installing a package that requires soft-links apperantly has not the required library installed by default.  
The Dockerfile has been extended with another layer installing Minikube for learning purposes.