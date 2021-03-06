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

You can find the physical address of each device in /proc/iomem.
```
$ sudo cat /proc/iomem 
00000000-00000fff : Reserved
00001000-0009fbff : System RAM
0009fc00-0009ffff : Reserved
000a0000-000bffff : PCI Bus 0000:00
000c0000-000c99ff : Video ROM
000ca000-000cadff : Adapter ROM
000cb000-000cd3ff : Adapter ROM
000f0000-000fffff : Reserved
  000f0000-000fffff : System ROM
00100000-7ffcffff : System RAM
7ffd0000-7fffffff : Reserved
80000000-afffffff : PCI Bus 0000:00
b0000000-bfffffff : PCI MMCONFIG 0000 [bus 00-ff]
  b0000000-bfffffff : Reserved
c0000000-febfffff : PCI Bus 0000:00
  d0000000-efffffff : PCI Bus 0000:04             <--- target of this test
    d0000000-dfffffff : 0000:04:00.0
    e0000000-e1ffffff : 0000:04:00.0
  f0000000-f07fffff : 0000:00:01.0
  f0800000-f09fffff : PCI Bus 0000:09
  f0a00000-f0bfffff : PCI Bus 0000:08
    f0a00000-f0a03fff : 0000:08:00.0
      f0a00000-f0a03fff : virtio-pci-modern
  f0c00000-f0dfffff : PCI Bus 0000:07
    f0c00000-f0c03fff : 0000:07:00.0
      f0c00000-f0c03fff : virtio-pci-modern
  f0e00000-f0ffffff : PCI Bus 0000:06
  f1000000-f11fffff : PCI Bus 0000:05
  f1200000-f13fffff : PCI Bus 0000:03
    f1200000-f1203fff : 0000:03:00.0
      f1200000-f1203fff : virtio-pci-modern
  f1400000-f15fffff : PCI Bus 0000:02
    f1400000-f1403fff : 0000:02:00.0
      f1400000-f1403fff : virtio-pci-modern
  f1600000-f17fffff : PCI Bus 0000:01
    f1600000-f1603fff : 0000:01:00.0
      f1600000-f1603fff : virtio-pci-modern
  f1800000-f1803fff : 0000:00:01.0
    f1800000-f1803fff : virtio-pci-modern
  fb000000-fcffffff : PCI Bus 0000:04
    fb000000-fbffffff : 0000:04:00.0
      fb000000-fbffffff : nvidia
    fc000000-fc07ffff : 0000:04:00.0
  fd000000-fd1fffff : PCI Bus 0000:09
  fd200000-fd3fffff : PCI Bus 0000:08
  fd400000-fd5fffff : PCI Bus 0000:07
  fd600000-fd7fffff : PCI Bus 0000:06
    fd600000-fd603fff : 0000:06:00.0
      fd600000-fd603fff : nvme
  fd800000-fd9fffff : PCI Bus 0000:05
    fd800000-fd803fff : 0000:05:00.0
      fd800000-fd803fff : ICH HD audio
  fda00000-fdbfffff : PCI Bus 0000:03
    fda00000-fda00fff : 0000:03:00.0
  fdc00000-fddfffff : PCI Bus 0000:02
    fdc00000-fdc00fff : 0000:02:00.0
  fde00000-fdffffff : PCI Bus 0000:01
    fde00000-fde7ffff : 0000:01:00.0
    fde80000-fde80fff : 0000:01:00.0
  fe010000-fe013fff : 0000:00:1b.0
    fe010000-fe013fff : ICH HD audio
  fe014000-fe014fff : 0000:00:01.0
  fe015000-fe015fff : 0000:00:02.0
  fe016000-fe016fff : 0000:00:02.1
  fe017000-fe017fff : 0000:00:02.2
  fe018000-fe018fff : 0000:00:02.3
  fe019000-fe019fff : 0000:00:02.4
  fe01a000-fe01afff : 0000:00:02.5
  fe01b000-fe01bfff : 0000:00:02.6
  fe01c000-fe01cfff : 0000:00:02.7
  fe01d000-fe01dfff : 0000:00:03.0
  fe01e000-fe01efff : 0000:00:1d.7
    fe01e000-fe01efff : ehci_hcd
  fe01f000-fe01ffff : 0000:00:1f.2
    fe01f000-fe01ffff : ahci
fec00000-fec003ff : IOAPIC 0
fed1c000-fed1ffff : Reserved
  fed1f410-fed1f414 : iTCO_wdt.1.auto
fee00000-fee00fff : Local APIC
feffc000-feffffff : Reserved
fffc0000-ffffffff : Reserved
100000000-27fffffff : System RAM
  1ba600000-1bb400eb0 : Kernel code
  1bb400eb1-1bbe5763f : Kernel data
  1bc120000-1bc5fffff : Kernel bss
280000000-a7fffffff : PCI Bus 0000:00
```
# 2. BAR
Comparing Bar's data with resouce file in /sys/bus/pci/devices/0000:0x:00.0.
```
$ ls -l resource*
-r--r--r-- 1 root root      4096  7??? 31 08:03 resource
-rw------- 1 root root  16777216  7??? 31 08:32 resource0
-rw------- 1 root root 268435456  7??? 31 08:32 resource1
-rw------- 1 root root 268435456  7??? 31 08:32 resource1_wc
-rw------- 1 root root  33554432  7??? 31 08:32 resource3
-rw------- 1 root root  33554432  7??? 31 08:32 resource3_wc
-rw------- 1 root root       128  7??? 31 08:32 resource5
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
