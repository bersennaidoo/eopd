package gentoken

import "sync/atomic"

var ops uint64

// GenerateTokenNumber Generate token number for patient.
func GenerateTokenNumber(start uint64) uint64 {

	if start > 0 {
		ops = start
		return ops
	}
	atomic.AddUint64(&ops, 1)

	return ops
}
