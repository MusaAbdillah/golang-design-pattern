package main 

import "fmt"

type Folder struct {
	components []Component
	name string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching keyword %s on folder %s \n", keyword, f.name)

	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
