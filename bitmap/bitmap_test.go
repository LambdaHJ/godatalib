package bitmap

import "testing"

func TestBitMap_Set(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name   string
		bm *BitMap
		args   args
	}{
		{name: "10bit", bm: MakeBitMap(20), args: args{x:10}},
		{name: "0bit", bm: MakeBitMap(20), args: args{x:0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bm.Set(tt.args.x)
			exists, _ := tt.bm.Exist(tt.args.x)
			if !exists {
				t.Fatalf("args %d, expect %v, got %v", tt.args.x, true, exists)
			}
			if (tt.args.x > 1) {
				exists, _ = tt.bm.Exist(tt.args.x - 1)
				if exists {
					t.Fatalf("args %d, expect %v, got %v", tt.args.x -1, false, exists)
				}
			}
		})
	}
}
func TestBitMap_Clear(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name   string
		bm *BitMap
		args   args
	}{
		{name: "10bit", bm: MakeBitMap(20), args: args{x:10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bm.Set(tt.args.x)
			exists, _ := tt.bm.Exist(tt.args.x)
			if !exists {
				t.Fatalf("args %d, expect %v, got %v", tt.args.x, true, exists)
			}
			tt.bm.Clear(tt.args.x)
			exists, _ = tt.bm.Exist(tt.args.x)
			if exists {
				t.Fatalf("args %d, expect %v, got %v", tt.args.x, false, exists)
			}
		})
	}
}


func TestBitMap_TrySet(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name   string
		bm *BitMap
		args   args
	}{
		{name: "10bit", bm: MakeBitMap(20), args: args{x:10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bm.Set(tt.args.x)
			exists, _ := tt.bm.Exist(tt.args.x)
			if !exists {
				t.Fatalf("args %d, expect %v, got %v", tt.args.x, true, exists)
			}
			success, _ := tt.bm.TrySet(tt.args.x)
			
			if success {
				t.Fatalf("args %d, expect %v, got %v", tt.args.x, false, success)
			}
		})
	}
}

func TestBitMap_CheckOverflow(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name   string
		bm *BitMap
		args   args
	}{
		{name: "30bit", bm: MakeBitMap(20), args: args{x:20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bm.Set(tt.args.x); err != OverFlowError {
				t.Fatalf("args %d, expect %v, got %v", tt.args.x, OverFlowError, err)
			}
			
		})
	}
}