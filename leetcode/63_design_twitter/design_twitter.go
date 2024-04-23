package design_twitter

import (
	"container/heap"
)

type Twitter struct {
	tweets    map[int]*Tweet
	follows   map[int]map[int]bool
	requestId int
}

type Tweet struct {
	tweetId   int
	previous  *Tweet
	timestamp int
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(1)
func Constructor() Twitter {
	return Twitter{
		tweets:    make(map[int]*Tweet),
		follows:   make(map[int]map[int]bool),
		requestId: 0,
	}
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) PostTweet(userId int, tweetId int) {
	twitter.InitUser(userId)
	twitter.requestId += 1
	twitter.tweets[userId] = &Tweet{
		tweetId:   tweetId,
		previous:  twitter.tweets[userId],
		timestamp: twitter.requestId,
	}
}

// Worst time complexity:   O(f log f) where f is the number of followers of the user.
// Average time complexity: O(a log a) where a is the average number of followers a user has.
// Space complexity:        O(f) where f is the number of followers of the user.
func (twitter *Twitter) GetNewsFeed(userId int) []int {
	twitter.InitUser(userId)

	newsFeed, tweetHeap := []int{}, &TweetMaxHeap{}
	heap.Init(tweetHeap)

	for folleweeId, isFollowing := range twitter.follows[userId] {
		if isFollowing {
			twitter.InitUser(folleweeId)
			heap.Push(tweetHeap, twitter.tweets[folleweeId])
		}
	}

	for len(*tweetHeap) > 0 && len(newsFeed) < 10 {
		nextTweet := heap.Pop(tweetHeap).(*Tweet)

		if nextTweet.timestamp == 0 {
			break
		}

		newsFeed = append(newsFeed, nextTweet.tweetId)
		heap.Push(tweetHeap, nextTweet.previous)
	}

	return newsFeed
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) Follow(followerId int, followeeId int) {
	twitter.InitUser(followerId)
	twitter.follows[followerId][followeeId] = true
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) Unfollow(followerId int, followeeId int) {
	twitter.InitUser(followerId)
	twitter.follows[followerId][followeeId] = false
}

// Worst time complexity:   O(1)
// Average time complexity: O(1)
// Space complexity:        O(n)
func (twitter *Twitter) InitUser(userId int) {
	if _, hasTweet := twitter.tweets[userId]; hasTweet {
		return
	}

	twitter.tweets[userId] = &Tweet{}
	twitter.follows[userId] = make(map[int]bool)
	twitter.follows[userId][userId] = true
}

type TweetMaxHeap []*Tweet

func (tweetHeap TweetMaxHeap) Len() int {
	return len(tweetHeap)
}

func (tweetHeap TweetMaxHeap) Less(i, j int) bool {
	return tweetHeap[i].timestamp > tweetHeap[j].timestamp
}

func (tweetHeap TweetMaxHeap) Swap(i, j int) {
	tweetHeap[i], tweetHeap[j] = tweetHeap[j], tweetHeap[i]
}

func (tweetHeap *TweetMaxHeap) Push(x any) {
	*tweetHeap = append(*tweetHeap, x.(*Tweet))
}

func (tweetHeap *TweetMaxHeap) Pop() any {
	current := *tweetHeap
	value := current[len(current)-1]
	*tweetHeap = current[:len(current)-1]

	return value
}
