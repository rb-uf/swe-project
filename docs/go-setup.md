# Golang env setup
## Installing Go

1. Follow the documentation for downloading and setup [here](https://go.dev/doc/install). If you're on windows make sure to update your path variable so that you can use 'go' from the terminal. To verify its working run `go version` and the output should look like `go version go1.19.5 windows/amd64`
2. (optional) If you want to use vs code, there's plenty of installation guides for it but for go, all you have to do is search for "go" in extensions and grab the top one. For the next steps you can just use the terminal within vs code too
---

## Setting up a Project with various packages (GORM, gin, and mux)

### What Works
- As long as go was installed correctly, it doesn't matter where your project lives, although organization is nice (so make a directory somewhere that makes sense)
- (Note) If you're using vs code and want to launch into a certain working directory, you can just have a terminal open and do "code .", and it will launch a new instance of vs code at the directory you are currently in

### Common Practice
- Navigate to where Go is installed: you can find this by opening a terminal, running `go env GOPATH`, copy the output path and cd to it.
- The go directory should have a few others in it, including bin, pkg, and src. We want to go into src, then 		    github.com. Make a directory (ideally named your git username), and within that is where our project will live. Although as long as your gopath variable is to the go directory mentioned it's fine, and your code can really be anywhere. (this directory is not present by default, you can just add a src directory in the go directory, then a github.com directory in that, and then your username as a directory in that)

### Setup
1. Open a terminal or gitbash instance in the directory you want your project to be in. For me I can just open a gitbash prompt from file explorer, you could also just open one wherever, copy the absolute file path and paste it into a cd command
2. In the terminal we will first run **go mod init**, which will make a default go.mod file which lists the version of go being used, and information on dependencies from external libraries (e.g. GORM, gin, etc.)
3. Next run `go get -u <package>`, where you can substitute \<package> with whatever you want to install. To find what's supposed to go here, use [this](https://pkg.go.dev/) to search for packages. For instance, if I wanted to use gorilla/mux, I would first search for mux, then click on gorilla/mux (this is also the first thing that comes up, most libraries already mentioned are popular enough to be first). 
4. In the page, at the top there's a line that says "Discover Packages \> github.com/gorilla/mux", and generally what you need is to the right of the "\>". After running "go get -u github.com/gorilla/mux", you will notice that the go.mod file has been updated to reflect the new dependency. If you try to compile and run some of the starter code in the documentation, it *should* work. 
5. (Note) If you look at %GOPATH%/pkg you'll notice the packages added are stored here. If you check %GOPATH%/src, some of the packages are installed with some starter code (see verifying... for more)
6. You can repeat this for gin and gorm, but to gloss over it, search for them on [pkg.go.dev](https://pkg.go.dev), grab the path (for gin it's github.com/gin-gonic/gin, for gorm it's gorm.io/gorm)
--- 

## Verifying that it installed correctly
- I copied and merged the starter code in the documentation for go (hello world) and from each package mentioned so far, and just made sure that it compiled and ran as I haven't figured out how to really do anything yet.
- Alternativley you can go to %GOPATH%/src, and look through the directories for programs that came with the packages you installed and try running those. A bunch of them use other external packages so I wouldn't worry about getting them to run but more so just confirming that there is no error regarding the associated package.
---

## Notes/TODO
1. I'll likely continue documenting my way through sprint 1 just to get an idea of what the hell I'm doing and how to actually do something
2. If there are any issues or clarifications with the steps I loosely outlined I'm happy to update and clarify anything, just dm or ping me on discord or email me at emmett.kogan@ufl.edu
---

## References/Links
1. A video I'm currently watching to figure out more about go, I also used it to check against what I figured out from reading documentation as far as setting up a project is concerned: https://www.youtube.com/watch?v=lf_kiH_NPvM
2. Go Documentation: https://go.dev/learn/
3. Go Package Search: https://pkg.go.dev/
4. GORM Documentation: https://gorm.io/
5. Gin Documentation: https://gin-gonic.com/docs/
6. gorilla/mux Documentation (this is just from pkg.go.dev): https://pkg.go.dev/github.com/gorilla/mux#section-readme

