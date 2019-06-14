package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "zero",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
		{
			name: "1",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val:  1,
															Next: nil,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val:  1,
														Next: nil,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbers1(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "ex",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "zero",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
		{
			name: "1",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val:  1,
															Next: nil,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val:  1,
														Next: nil,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ex",
			args: args{
				l1: &ListNode{
					Val:  1,
					Next: nil,
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers1(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers1() = %v, want %v", got, tt.want)
			}
		})
	}
}
