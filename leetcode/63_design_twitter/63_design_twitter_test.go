package design_twitter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/63_design_twitter"
)

var _ = Describe("63DesignTwitter", func() {
	It("get news feed of user with own and follower tweets", func() {
		twitter := Constructor()
		twitter.PostTweet(1, 5)
		Expect(twitter.GetNewsFeed(1)).To(Equal([]int{5}))
		twitter.Follow(1, 2)
		twitter.PostTweet(2, 6)
		Expect(twitter.GetNewsFeed(1)).To(Equal([]int{6, 5}))
		twitter.Unfollow(1, 2)
		Expect(twitter.GetNewsFeed(1)).To(Equal([]int{5}))
	})

	It("should get news feed of user with follower tweets", func() {
		twitter := Constructor()
		twitter.PostTweet(1, 1)
		Expect(twitter.GetNewsFeed(1)).To(Equal([]int{1}))
		twitter.Follow(2, 1)
		Expect(twitter.GetNewsFeed(2)).To(Equal([]int{1}))
		twitter.Unfollow(2, 1)
		Expect(twitter.GetNewsFeed(2)).To(Equal([]int{}))
	})

	It("should get news feed of user with no tweets", func() {
		twitter := Constructor()
		Expect(twitter.GetNewsFeed(1)).To(Equal([]int{}))
		twitter.Follow(1, 5)
		Expect(twitter.GetNewsFeed(1)).To(Equal([]int{}))
	})
})
