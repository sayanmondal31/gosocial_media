package db

import (
	"context"
	"math/rand"

	"fmt"
	"log"

	"github.com/sayanmondal31/gosocial/internal/store"
)

var usersnames = []string{
	"alice", "bob", "charlie", "david", "emma",
	"frank", "grace", "henry", "isabella", "jack",
	"kate", "leo", "mia", "nathan", "olivia",
	"paul", "quinn", "rose", "sam", "tina",
	"uma", "victor", "wendy", "xavier", "yasmin",
	"zack", "aaron", "bella", "caleb", "diana",
	"ethan", "fiona", "george", "hannah", "ian",
	"julia", "kevin", "lily", "mark", "nora",
	"owen", "piper", "rachel", "steve", "tracy",
	"ursula", "vincent", "willow", "zoe",
}

var titles = []string{
	"go concurrency basics",
	"goroutines explained",
	"channels in practice",
	"go vs nodejs",
	"building rest api in go",
	"go memory management",
	"error handling in go",
	"mutex vs channel",
	"go project structure",
	"grpc with golang",
	"go performance tips",
	"context package guide",
	"go worker pools",
	"rate limiting in go",
	"go clean architecture",
	"testing in golang",
	"go deployment basics",
	"go with docker",
	"kafka with golang",
	"go interview questions",
}

var contents = []string{
	"this article explores simple ideas with practical examples",
	"a quick walkthrough of common mistakes and fixes",
	"learn how small optimizations improve performance",
	"a beginner friendly guide with real world use cases",
	"deep dive into concepts every developer should know",
	"practical tips learned from production systems",
	"understanding tradeoffs in system design decisions",
	"how things work behind the scenes explained simply",
	"less theory more hands on implementation",
	"patterns that scale well in real applications",
	"common pitfalls and how to avoid them",
	"writing clean and maintainable code",
	"designing APIs that are easy to use",
	"handling edge cases the right way",
	"improving reliability with simple techniques",
	"lessons learned from debugging production bugs",
	"making code faster without overengineering",
	"thinking like a backend engineer",
	"building features step by step",
	"small changes that make a big difference",
}

var tags = []string{
	"golang",
	"backend",
	"api",
	"concurrency",
	"performance",
	"microservices",
	"grpc",
	"kafka",
	"docker",
	"kubernetes",
	"cloud",
	"scalability",
	"system-design",
	"distributed-systems",
	"databases",
	"mongodb",
	"redis",
	"devops",
	"testing",
	"architecture",
}

var comments = []string{
	"great explanation, very easy to follow",
	"this cleared up a lot of confusion",
	"nice breakdown with practical examples",
	"exactly what i was looking for",
	"well written and concise",
	"helpful content, thanks for sharing",
	"the examples made it much clearer",
	"looking forward to more posts like this",
	"simple and to the point",
	"this saved me a lot of time",
	"good insights from real experience",
	"easy to understand even for beginners",
	"clear and practical approach",
	"learned something new today",
	"this is very useful in real projects",
	"appreciate the straightforward explanation",
	"well structured and informative",
	"thanks for breaking it down so clearly",
	"great tips, will try this out",
	"solid explanation with real value",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := range num {
		users[i] = &store.User{
			Username: usersnames[i%len(usersnames)] + fmt.Sprintf("%d", i),
			Email:    usersnames[i%len(usersnames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserId:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {

	cmnts := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cmnts[i] = &store.Comment{
			PostId:  posts[rand.Intn(len(posts))].ID,
			UserId:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cmnts
}
