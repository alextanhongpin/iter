package iter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/alextanhongpin/iter"
)

func TestReduce(t *testing.T) {
	got := iter.Reduce([]int{1, 2, 3}, 0, func(acc, n int) int {
		return acc + n
	})
	if exp := 6; exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}

func TestReduceIndex(t *testing.T) {
	chars := []string{"a", "b", "c"}
	got := iter.ReduceIndex(chars, "", func(acc string, i int) string {
		return acc + chars[i]
	})
	if exp := "abc"; exp != got {
		t.Errorf("expected %s, got %s", exp, got)
	}
}

type post struct {
	id     int
	photos []*photo
}

type photo struct {
	id     int
	postID int
	tags   []tag
}

type tag struct {
	id      int
	photoID int
	name    string
}

func TestLazyLoading(t *testing.T) {
	start := time.Now()
	posts := []post{{id: 1}, {id: 2}, {id: 3}}

	fetchPhotos := func(p post) []*photo {
		photos := make([]*photo, 3)
		for i := 0; i < 3; i++ {
			photos[i] = &photo{id: p.id * i, postID: p.id}
		}
		time.Sleep(1 * time.Second)
		return photos
	}

	fetchTags := func(p *photo) []tag {
		tags := make([]tag, 5)
		for i := 0; i < 5; i++ {
			tags[i] = tag{id: p.id * i, photoID: p.id}
		}
		time.Sleep(1 * time.Second)
		return tags
	}

	// Wait for all photos to fetch.
	iter.GoEachIndex(posts, func(i int) {
		photos := fetchPhotos(posts[i])
		posts[i].photos = photos
	})

	// Wait for all tags to fetch.
	iter.GoEachIndex(posts, func(i int) {
		iter.GoEachIndex(posts[i].photos, func(j int) {
			posts[i].photos[j].tags = fetchTags(posts[i].photos[j])
		})
	})

	t.Log(time.Since(start))
	for i := range posts {
		t.Logf("%d: post id: %+v\n", i+1, posts[i].id)
		t.Log(len(posts[i].photos), "photos")

		for j := range posts[i].photos {
			t.Logf("\t%+v\n", posts[i].photos[j])
		}
		fmt.Println()
	}
}
