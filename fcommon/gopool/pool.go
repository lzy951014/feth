/*
 * @Author: Liuzongyun 845666459@qq.com
 * @Date: 2024-11-12 10:46:07
 * @LastEditors: Liuzongyun 845666459@qq.com
 * @LastEditTime: 2024-11-12 10:50:14
 * @FilePath: /feth/fcommon/gopool/pool.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package gopool

import (
	"runtime"
	"time"

	"github.com/panjf2000/ants/v2"
)

var (
	// Init a instance pool when importing ants.
	defaultPool, _   = ants.NewPool(ants.DefaultAntsPoolSize, ants.WithExpiryDuration(10*time.Second))
	minNumberPerTask = 5
)

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})
}

// Submit submits a task to pool.
func Submit(task func()) error {
	return defaultPool.Submit(task)
}

// Running returns the number of the currently running goroutines.
func Running() int {
	return defaultPool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return defaultPool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return defaultPool.Free()
}

// Release Closes the default pool.
func Release() {
	defaultPool.Release()
}

// Reboot reboots the default pool.
func Reboot() {
	defaultPool.Reboot()
}

func Threads(tasks int) int {
	threads := tasks / minNumberPerTask
	if threads > runtime.NumCPU() {
		threads = runtime.NumCPU()
	} else if threads == 0 {
		threads = 1
	}
	return threads
}
