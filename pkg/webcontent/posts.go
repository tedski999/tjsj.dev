package webcontent

// TODO: data structure
type Post struct {
	data string
}

func (content *Content) GetPost(name string) Post {
	return content.posts[name]
}

func (content *Content) loadPosts() error {
	content.posts = make(map[string]Post)
	// TODO: load posts
	content.posts["test"] = Post { "hello" }
	return nil
}
