package main

func main() {
    
    // Will output error if there two coexisting paths
    fs := NewFS(".")

    finder := wsii.Finder{
        Queue: goconcurrentqueue.NewFIFO(),
        FS: index,
    }
}
