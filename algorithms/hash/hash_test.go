package hash

// TODO: revisit when re-attempting to correctly implement hashing algorithms.

//func TestDJB2(t *testing.T) {
//	type args struct {
//		data any
//	}
//	tests := []struct {
//		name string
//		args args
//		want uint32
//	}{
//		{"hash test", args{"test"}, 123},
//		{"hash abc", args{"abc"}, 123},
//		{"hash 123", args{"123"}, 123},
//		{"hash äöü", args{"äöü"}, 123},
//		{"hash i(123)", args{123}, 123},
//		{"hash i(124)", args{124}, 123},
//		{"hash i(125)", args{125}, 123},
//		{"hash i(126)", args{126}, 123},
//		{"hash i(127)", args{127}, 123},
//	}
//	for _, tt := range tests {
//		startT := time.Now()
//		t.Run(tt.name, func(t *testing.T) {
//			if got, _ := DJB2(tt.args.data); got != tt.want {
//				t.Errorf("Murmur2() = %v, want %v", got, tt.want)
//			}
//		})
//		perf.TimeTracker(startT, tt.name)
//		perf.PrintMemUsage(perf.KB, tt.name)
//	}
//}
