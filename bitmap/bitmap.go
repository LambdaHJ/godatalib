package bitmap

import (
	"errors"
	"sync"
)

var (
	OverFlowError = errors.New("value overflow")
)


type BitMap struct {
	sync.RWMutex
	data []byte
	cap int
}

// MakeBitMap
// no error return
func MakeBitMap(cap int) *BitMap {
	s :=  (cap + 7)%8+1
	d := make([]byte, s, s)
	return &BitMap{data: d, cap: cap}
}

// Set set bit.
// if x > cap, OverFlowError will be returned.
func (bm *BitMap) Set(x uint) error {
	if bm.checkOverflow(x) {
		return  OverFlowError
	}
	bm.Lock()
	defer bm.Unlock()
	bm.data[x >> 3] |= 1 << (x&0x07)
	return nil
}

// Exist check bit has been setted.
// if x > cap, OverFlowError will be returned.
func (bm *BitMap) Exist(x uint) (bool, error) {
	if bm.checkOverflow(x) {
		return false,  OverFlowError
	}
	bm.RLock()
	defer bm.RUnlock()
	return bm.data[x >> 3] & (1 << (x&0x07)) != 0, nil
}

// Clear set bit 0.
// if x > cap, OverFlowError will be returned.
func (bm *BitMap) Clear(x uint) error {
	if bm.checkOverflow(x) {
		return OverFlowError
	}
	bm.Lock()
	defer bm.Unlock()
	bm.data[x >> 3] &^= 1 << (x&0x07)
	return nil
}

// TrySet set bit=0 if bit not been setted.
func (bm *BitMap) TrySet(x uint) (bool, error) {
	if bm.checkOverflow(x) {
		return false, OverFlowError
	}
	bm.RLock()
	if (bm.data[x >> 3] & (1 << (x&0x07)) == 0) {
		bm.RUnlock()
		bm.Lock()
		rtn := false
		if (bm.data[x >> 3] & (1 << (x&0x07)) == 0) {
			bm.data[x >> 3] |= 1 << (x&0x07)
			rtn = true
		}
		bm.Unlock()
		return rtn, nil
	}
	bm.RUnlock()
	return false, nil
}

// Cap return bitmap max value.
func (bm *BitMap) Cap() int {
	return bm.cap
}

func (bm *BitMap) checkOverflow(x uint) bool {
	return bm.cap <= int(x)
}