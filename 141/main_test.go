package linkedlistcycle

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	nfour := &ListNode{Val: -4}
	zero := &ListNode{Val: 0, Next: nfour}
	two := &ListNode{Val: 2, Next: zero}
	three := &ListNode{Val: 3, Next: two}
	nfour.Next = two

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				head: three,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
