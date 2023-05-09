package meeting_rooms_2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/37_meeting_rooms_2"
)

var _ = Describe("37MeetingRooms2", func() {
	It("should get minimum meeting rooms", func() {
		Expect(MinMeetingRooms([]*Interval{})).To(Equal(0))
		Expect(MinMeetingRooms([]*Interval{
			{2, 7},
		})).To(Equal(1))
		Expect(MinMeetingRooms([]*Interval{
			{0, 30},
			{5, 10},
			{15, 20},
		})).To(Equal(2))
		Expect(MinMeetingRooms([]*Interval{
			{0, 2},
			{1, 3},
			{3, 5},
			{4, 6},
		})).To(Equal(2))
		Expect(MinMeetingRooms([]*Interval{
			{4, 6},
			{0, 2},
			{3, 5},
			{1, 3},
		})).To(Equal(2))
		Expect(MinMeetingRooms([]*Interval{
			{0, 2},
			{1, 3},
			{4, 6},
			{5, 7},
		})).To(Equal(2))
		Expect(MinMeetingRooms([]*Interval{
			{5, 7},
			{4, 6},
			{1, 3},
			{0, 2},
		})).To(Equal(2))
		Expect(MinMeetingRooms([]*Interval{
			{13, 17},
			{14, 18},
			{15, 20},
			{19, 21},
		})).To(Equal(3))
	})
})
