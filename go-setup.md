# Golang env setup
## Installing Go

1. Follow the documentation for downloading and setup [here](https://go.dev/doc/install). If you're on windows make sure to update your path variable so that you can use 'go' from the terminal.
		Note: if it is working correctly, you should be able to open a terminal, enter "go version" and the output should look like "go version go1.19.5 windows/amd64"
2. (optional) If you want to use vs code, there's plenty of installation guides for it but for go, all you have to do is search for "go" in extensions and grab the top one. If using vs code, for the next steps you can just use the terminal within vs code
---
## Setting up a Project with various packages (GORM, gin, and mux)
1. To setup a project, first navigate to where Go is installed; for me, this is C:\Users\emmet\go, it might be somewhere different for you.
2. The Go directory should have a few others in it, including bin, pkg, and src. We want to go into src, then github.com. Make a directory (ideally named your git username), and within that is where our project will live. I'm relatively certain this step doesn't matter but also haven't had time to confirm that and for now I'm going to stick to this as it works
3. Now open a terminal instance to this directory, for me I can just open a gitbash prompt from file explorer, you could also just open one wherever, copy the absolute file path and paste it into a cd command
4. In the terminal we will first run "go mod init", which will make a default go.mod file which lists the version of go being used, and information on dependencies from external libraries (e.g. GORM, gin or mux)
5. Next run "go get -u \<library>", where you can substitute \<library> with whatever you want to install. To find what's supposed to go here, use [this](https://pkg.go.dev/) to search for packages. For instance, if I wanted to use gorilla/mux, I would first search for mux, the click on gorilla/mux (this is also the first thing that comes up, most libraries already mentioned are popular enough to be first). In the page, at the top there's a line that says "Discover Packages \> github.com/gorilla/mux", and the link (with the copy button) is what you'll need to place in the get command. After running "go get -u github.com/gorilla/mux", you will notice that the go.mod file has been updated to reflect the new dependency, and if you try to compile or run some of the starter code, it *should* work
6. You can repeat this for gin and gorm, but to gloss over it, search for them on [pkg.go.dev](https://pkg.go.dev), grab the path (for gin it's github.com/gin-gonic/gin, for gorm it's gorm.io/gorm)
--- 
## Verifying that it installed correctly
1. I basically copied and merged the starter code in the documentation for go (hello world), as well as from each package mentioned so far, and just made sure that it compiled and ran as I haven't figured out how to really do anything yet.
---
## Notes/TODO
1. I'll likely continue documenting my way through sprint 1 just to get an idea of what the hell I'm doing and how to actually do something
2. If there are any issues or clarifications with the steps I loosely outlined I'm happy to update and clarify anything, just dm or ping me on discord or email me at emmett.kogan@ufl.edu
---
## References/Links
A video I'm currently watching to figure out more about go, I also used it to check against what I figured out from reading documentation as far as setting up a project is concerned:
&nbsp;&nbsp;&nbsp;&nbsp;https://www.youtube.com/watch?v=lf_kiH_NPvM
Go Documentation:
&nbsp;&nbsp;&nbsp;&nbsp;https://go.dev/learn/
Go Package Search:
&nbsp;&nbsp;&nbsp;&nbsp;https://pkg.go.dev/
GORM Documentation:
&nbsp;&nbsp;&nbsp;&nbsp;https://gorm.io/
Gin Documentation:
&nbsp;&nbsp;&nbsp;&nbsp;https://gin-gonic.com/docs/
gorilla/mux Documentation (this is just from pkg.go.dev):
&nbsp;&nbsp;&nbsp;&nbsp;https://pkg.go.dev/github.com/gorilla/mux#section-readme
