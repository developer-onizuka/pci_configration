package main

import (
	"os"
	"fmt"
	"syscall"
	//"unsafe"
	//"time"
)

func main() {
	fmt.Printf("----- PCI BAR Space and /sys/bus/pci/devices/0000:0x:00.0/resource -----\n")
	f, _ := os.OpenFile("/dev/mem", os.O_RDWR, 0666)
        defer f.Close()
	bar, err := syscall.Mmap(int(f.Fd()), int64(0xd0000000), 64, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	fmt.Printf("err: %v\n", err)
	defer syscall.Munmap(bar)
	fmt.Printf("bar: %#x\n", bar)

	g, _ := os.OpenFile("/sys/bus/pci/devices/0000:04:00.0/resource1", os.O_RDWR, 0666)
        defer g.Close()
	res, err := syscall.Mmap(int(g.Fd()), 0, 64, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	fmt.Printf("err: %v\n", err)
	defer syscall.Munmap(res)
	fmt.Printf("res: %#x\n", res)
}
