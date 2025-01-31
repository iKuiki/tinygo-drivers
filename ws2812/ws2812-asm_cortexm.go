//go:build cortexm
// +build cortexm

package ws2812

// Warning: autogenerated file. Instead of modifying this file, change
// gen-ws2812.go and run "go generate".

import (
	"device"
	"runtime/interrupt"
)

func (d Device) writeByte16(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	// Timings:
	// T0H:  6 -  8 cycles or 375.0ns - 500.0ns
	// T1H: 17 - 19 cycles or 1062.5ns - 1187.5ns
	// TLD: 19 -    cycles or 1187.5ns -
	mask := interrupt.Disable()
	value := uint32(c) << 24
	device.AsmFull(`
	1: @ send_bit
	  str   {maskSet}, {portSet}     @ [2]   T0H and T0L start here
	  nop
	  nop
	  lsls  {value}, #1              @ [1]
	  bcs.n 2f                       @ [1/3] skip_store
	  str   {maskClear}, {portClear} @ [2]   T0H -> T0L transition
	2: @ skip_store
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  str   {maskClear}, {portClear} @ [2]   T1H -> T1L transition
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  subs  {i}, #1                  @ [1]
	  beq.n 3f                       @ [1/3] end
	  b     1b                       @ [1/3] send_bit
	3: @ end
	`, map[string]interface{}{
		"value":     value,
		"i":         8,
		"maskSet":   maskSet,
		"portSet":   portSet,
		"maskClear": maskClear,
		"portClear": portClear,
	})
	interrupt.Restore(mask)
}

func (d Device) writeByte48(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	// Timings:
	// T0H: 17 - 19 cycles or 354.2ns - 395.8ns
	// T1H: 51 - 53 cycles or 1062.5ns - 1104.2ns
	// TLD: 56 -    cycles or 1166.7ns -
	mask := interrupt.Disable()
	value := uint32(c) << 24
	device.AsmFull(`
	1: @ send_bit
	  str   {maskSet}, {portSet}     @ [2]   T0H and T0L start here
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  lsls  {value}, #1              @ [1]
	  bcs.n 2f                       @ [1/3] skip_store
	  str   {maskClear}, {portClear} @ [2]   T0H -> T0L transition
	2: @ skip_store
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  str   {maskClear}, {portClear} @ [2]   T1H -> T1L transition
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  subs  {i}, #1                  @ [1]
	  beq.n 3f                       @ [1/3] end
	  b     1b                       @ [1/3] send_bit
	3: @ end
	`, map[string]interface{}{
		"value":     value,
		"i":         8,
		"maskSet":   maskSet,
		"portSet":   portSet,
		"maskClear": maskClear,
		"portClear": portClear,
	})
	interrupt.Restore(mask)
}

func (d Device) writeByte64(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	// Timings:
	// T0H: 23 - 25 cycles or 359.4ns - 390.6ns
	// T1H: 68 - 70 cycles or 1062.5ns - 1093.8ns
	// TLD: 74 -    cycles or 1156.2ns -
	mask := interrupt.Disable()
	value := uint32(c) << 24
	device.AsmFull(`
	1: @ send_bit
	  str   {maskSet}, {portSet}     @ [2]   T0H and T0L start here
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  lsls  {value}, #1              @ [1]
	  bcs.n 2f                       @ [1/3] skip_store
	  str   {maskClear}, {portClear} @ [2]   T0H -> T0L transition
	2: @ skip_store
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  str   {maskClear}, {portClear} @ [2]   T1H -> T1L transition
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  subs  {i}, #1                  @ [1]
	  beq.n 3f                       @ [1/3] end
	  b     1b                       @ [1/3] send_bit
	3: @ end
	`, map[string]interface{}{
		"value":     value,
		"i":         8,
		"maskSet":   maskSet,
		"portSet":   portSet,
		"maskClear": maskClear,
		"portClear": portClear,
	})
	interrupt.Restore(mask)
}

func (d Device) writeByte120(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	// Timings:
	// T0H: 42 - 44 cycles or 350.0ns - 366.7ns
	// T1H: 126 - 128 cycles or 1050.0ns - 1066.7ns
	// TLD: 138 -    cycles or 1150.0ns -
	mask := interrupt.Disable()
	value := uint32(c) << 24
	device.AsmFull(`
	1: @ send_bit
	  str   {maskSet}, {portSet}     @ [2]   T0H and T0L start here
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  lsls  {value}, #1              @ [1]
	  bcs.n 2f                       @ [1/3] skip_store
	  str   {maskClear}, {portClear} @ [2]   T0H -> T0L transition
	2: @ skip_store
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  str   {maskClear}, {portClear} @ [2]   T1H -> T1L transition
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  subs  {i}, #1                  @ [1]
	  beq.n 3f                       @ [1/3] end
	  b     1b                       @ [1/3] send_bit
	3: @ end
	`, map[string]interface{}{
		"value":     value,
		"i":         8,
		"maskSet":   maskSet,
		"portSet":   portSet,
		"maskClear": maskClear,
		"portClear": portClear,
	})
	interrupt.Restore(mask)
}

func (d Device) writeByte168(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	// Timings:
	// T0H: 59 - 61 cycles or 351.2ns - 363.1ns
	// T1H: 177 - 179 cycles or 1053.6ns - 1065.5ns
	// TLD: 194 -    cycles or 1154.8ns -
	mask := interrupt.Disable()
	value := uint32(c) << 24
	device.AsmFull(`
	1: @ send_bit
	  str   {maskSet}, {portSet}     @ [2]   T0H and T0L start here
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  lsls  {value}, #1              @ [1]
	  bcs.n 2f                       @ [1/3] skip_store
	  str   {maskClear}, {portClear} @ [2]   T0H -> T0L transition
	2: @ skip_store
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  str   {maskClear}, {portClear} @ [2]   T1H -> T1L transition
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  nop
	  subs  {i}, #1                  @ [1]
	  beq.n 3f                       @ [1/3] end
	  b     1b                       @ [1/3] send_bit
	3: @ end
	`, map[string]interface{}{
		"value":     value,
		"i":         8,
		"maskSet":   maskSet,
		"portSet":   portSet,
		"maskClear": maskClear,
		"portClear": portClear,
	})
	interrupt.Restore(mask)
}
