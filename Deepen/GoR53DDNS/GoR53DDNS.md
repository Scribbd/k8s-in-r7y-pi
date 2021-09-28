---
---
# go-r53-ddns
Go Route53 DDNS module

Part of running a kubernetes cluster is to get certificates going. To get certificates going you need to be able to answer a challenge. To be able to answer a challenge you need a domain-name with a record pointing  towards the cluster. To have a valid IP you should not have a dynamic IP provided by your ISP. Which I have!

So, what now? I could go for a DDNS provider. Which is nothing more than a DNS provider with a way to easily update the target IP of a record  through a program. If you got a decent ISP and got a decent modem from them, you should be able to use a DDNS function on your modem. I don't have a decent ISP... 

All this writing to just say: I am making a cronjob for k8s to update my IP when it changes and be my own DDNS.  
This I do as an exercise in the following: 
 - Automate building of golang to docker image
 - Check out the AWS module in Ansible
 - Explore Ansible Vault for storing its AWS access key
 - Automate key rotation
 - Automate updating k8s secrets with new key

## IAM users
Two users are needed. One ansible user for deploying