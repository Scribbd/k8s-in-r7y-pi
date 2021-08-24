Needed packages:
```shell
apt-get install -y sg3-utils lsscsi
```

add trim to external sata-ssd:
```shell
# Get max values "Maximum unmap LBA count:":
max_LBA=sg_vpd -p bl /dev/sda | ((somehow getting the lba count))
blk_size=sg_readcap -l /dev/sda

echo 2147450880 > /sys/block/sda/queue/discard_max_bytes
```
make it stick:
```console
root@raspberrypi:/mnt# lsusb
Bus 002 Device 002: ID 174c:1153 ASMedia Technology Inc. ASM1153 SATA 3Gb/s bridge
Bus 002 Device 001: ID 1d6b:0003 Linux Foundation 3.0 root hub
Bus 001 Device 002: ID 2109:3431 VIA Labs, Inc. Hub
Bus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
```
idVendor: 174c
idProduct: 1153

```shell
echo 'ACTION=="add|change", ATTRS{idVendor}=="1b1c", ATTRS{idProduct}=="1a0e", SUBSYSTEM=="scsi_disk", ATTR{provisioning_mode}="unmap"' > /etc/udev/rules.d/10-trim.rules
```

automate the trimmer:
```shell
sudo systemctl enable fstrim.timer
```

