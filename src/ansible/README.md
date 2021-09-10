# README ansible

Ansible playbooks created:
- FreshPi: For freshly installed pis. Sets up some pi specific stuff and sets up the ssh-keys for further ansible use.
	 - Important: assumes that user 'pi' is still on the system
 - First5: first five minutes of setting up a linux box
    - Heavily inspired by [Jeff Geerlings security-role](https://github.com/geerlingguy/ansible-role-security)
    - Important: removes user 'pi'
 - prehorn: sets up an usb sata-ssd peripheral with UASP, and prepares file system for use with longhorn. When possible it will also apply TRIM support.
 - k8s_install: Installs k8s on all nodes
 - longhorn: installs longhorn to use with k8s

I also created a custom docker container. This is its Dockerfile:
![[ansible/Dockerfile]]

Use it with the following command:
`docker run -itd -v ansible_root:/root/ -v /var/run/docker.sock:/var/run/docker.sock -v E:\OneDrive\Documenten\ObsidianVaults\DockerCluster\src\ansible:/ansible --name ansible --rm ansible`

The 'exploded' more readable command:
```bash
docker run -itd \
-v ansible_root:/root/ \
-v /var/run/docker.sock:/var/run/docker.sock \
-v E:\OneDrive\Documenten\ObsidianVaults\DockerCluster\src\ansible:/ansible \
--name ansible \
--rm \
ansible
```

This command will create a docker container that is able to run molecule with docker from the host. All .ssh keys will be stored in a named volume ansible_keys.
