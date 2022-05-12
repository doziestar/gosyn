package app

import "github.com/doziestar/gosyn/app/slackbot"

// "github.com/doziestar/gosyn/app/bookstore"

// "github.com/doziestar/gosyn/app/movie"
// "github.com/doziestar/gosyn/app/sever"

func Start() {
	// var name string = "gosyn"
	// var version string = "0.0.1"

	// fmt.Printf("%s %s\n", name, version)
	// fmt.Println(&name)

	// name2 := &name // name2 is a pointer to name
	// fmt.Println(name2)

	// // name2 = "gosyn" // cannot assign to name2
	// *name2 = "gosyn2" // name2 is a pointer to name

	// fmt.Println(name)
	// fmt.Println(*name2)
	// sever.Serve()
	// movie.Movie()
	// bookstore.Start()
	slackbot.Init()
}
