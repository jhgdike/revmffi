//go:build darwin

package vm

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/../revm -L${SRCDIR}/../revm -lrevmapi -framework CoreFoundation -framework Security
import "C"
