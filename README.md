This is technical task for job interview

Technical task:
Create a utility to process YAML-file and output an /etc/fstab file based on that yaml.
The tasks must run on a Linux system.
Use a programming/scripting language to achieve this task, and not Ansible or such.

Code description:
The code is written in GoLang.
The file fstab actual for RHEL7. 
RHEL documentation: https://www.redhat.com/sysadmin/etc-fstab

Since this is a test task for output using a file in the current directory. If you want to change the output file, change the line in the file main.go
```
result := write("fstab_test", c)
```

There is option root-reserve in fstab.yml. Yoy could use it if you want to change the percentage of reserved blocks on this file system.
Documentation: https://man7.org/linux/man-pages/man8/tune2fs.8.html
Also documentation: https://linuxconfig.org/how-to-tune-linux-extended-ext-filesystems-using-dumpe2fs-and-tune2fs
After runing golang code you will have sh-script which you could run for configure fs

Input data: fstab.yml
Output data: fstab, tune2fs.sh

Requirements:

go	v1.18
gopkg.in/yaml.v3  v3.0.0

To run code use command: go run main.go
To build code use command: go build


Example fstab for rhel os
```
/dev/mapper/rhel-root                         /                       xfs     defaults        0 0
UUID=64351209-b3d4-421d-8900-7d940ca56fea     /boot                   xfs     defaults        0 0
/dev/mapper/rhel-swap                         swap                    swap    defaults        0 0
```

The table itself is a 6 column structure, where each column designates a specific parameter and must be set up in the correct order. The columns of the table are as follows from left to right: 

Device: usually the given name or UUID of the mounted device (sda1/sda2/etc).
Mount Point: designates the directory where the device is/will be mounted. 
File System Type: nothing trick here, shows the type of filesystem in use. 
Options: lists any active mount options. If using multiple options they must be separated by commas. 
Backup Operation: (the first digit) this is a binary system where 1 = dump utility backup of a partition. 0 = no backup. This is an outdated backup method and should NOT be used. 
File System Check Order: (second digit) Here we can see three possible outcomes.  0 means that fsck will not check the filesystem. Numbers higher than this represent the check order. The root filesystem should be set to 1 and other partitions set to 2. 


Alternative ways to refer to partitions:
Network ID
Samba : //server/share
NFS : server:/share
SSHFS : sshfs#user@server:/share
Device : /dev/sdxy (not recommended)
