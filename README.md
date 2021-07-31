# 1. lspci 
```
$ cd /sys/bus/pci/devices/0000:04:00.0
$ sudo lspci -v -x -s 04:00.0
04:00.0 VGA compatible controller: NVIDIA Corporation GP107GL [Quadro P1000] (rev a1) (prog-if 00 [VGA controller])
	Subsystem: NVIDIA Corporation GP107GL [Quadro P1000]
	Physical Slot: 0-3
	Flags: bus master, fast devsel, latency 0, IRQ 48
	Memory at fb000000 (32-bit, non-prefetchable) [size=16M]
	Memory at d0000000 (64-bit, prefetchable) [size=256M]
	Memory at e0000000 (64-bit, prefetchable) [size=32M]
	I/O ports at c000 [size=128]
	Expansion ROM at fc000000 [virtual] [disabled] [size=512K]
	Capabilities: [60] Power Management version 3
	Capabilities: [68] MSI: Enable+ Count=1/1 Maskable- 64bit+
	Capabilities: [78] Express Legacy Endpoint, MSI 00
	Capabilities: [100] Virtual Channel
	Capabilities: [250] Latency Tolerance Reporting
	Capabilities: [128] Power Budgeting <?>
	Capabilities: [420] Advanced Error Reporting
	Capabilities: [600] Vendor Specific Information: ID=0001 Rev=1 Len=024 <?>
	Kernel driver in use: nvidia
	Kernel modules: nvidiafb, nouveau, nvidia_drm, nvidia
00: de 10 b1 1c 07 05 10 00 a1 00 00 03 00 00 00 00
10: 00 00 00 fb 0c 00 00 d0 00 00 00 00 0c 00 00 e0
20: 00 00 00 00 01 c0 00 00 00 00 00 00 de 10 bc 11
30: 00 00 00 00 60 00 00 00 00 00 00 00 0b 01 00 00
```

# 2. BAR
Comparing Bar's data with resouce file in /sys/bus/pci/devices/0000:0x:00.0.
```
$ ls -l resource*
-r--r--r-- 1 root root      4096  7月 31 08:03 resource
-rw------- 1 root root  16777216  7月 31 08:32 resource0
-rw------- 1 root root 268435456  7月 31 08:32 resource1
-rw------- 1 root root 268435456  7月 31 08:32 resource1_wc
-rw------- 1 root root  33554432  7月 31 08:32 resource3
-rw------- 1 root root  33554432  7月 31 08:32 resource3_wc
-rw------- 1 root root       128  7月 31 08:32 resource5
```

The following is the result the data is consistent between mapped datas from /dev/mem and resource file.
```
$ sudo go run mmap-pci.go 
----- PCI BAR Space and /sys/bus/pci/devices/0000:0x:00.0/resource -----
err: <nil>
bar: 0x01005a0c000000000000000000000000050000c000000000000000006c0000000000000000000000000000000000000000000000280000000000000028000000
err: <nil>
res: 0x01005a0c000000000000000000000000050000c000000000000000006c0000000000000000000000000000000000000000000000280000000000000028000000
```

# 3. Configration
```
$ od -x config
0000000 10de 1cb1 0507 0010 00a1 0300 0000 0000
0000020 0000 fb00 000c d000 0000 0000 000c e000
0000040 0000 0000 c001 0000 0000 0000 10de 11bc
0000060 0000 0000 0060 0000 0000 0000 010b 0000
0000100
```

# 4. Vendor and Device name
```
$ cat vendor 
0x10de
$ cat device 
0x1cb1
```
