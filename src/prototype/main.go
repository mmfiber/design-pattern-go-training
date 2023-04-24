package prototype

import "fmt"

type Inode interface {
	print(string)
	Clone() Inode
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Printf("%s(pointer address: %p)\n", indentation+f.name, &f)
}

func (f *File) Clone() Inode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Printf("%s(pointer address: %p)\n", indentation+f.name, &f)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	cloneFolder := folder2.Clone()

	fmt.Println("\nPrinting hierarchy for orginal Folder")
	folder2.print("  ")

	fmt.Println("\nPrinting hierarchy for cloned Folder")
	cloneFolder.print("  ")
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "prototype pattern"
}

func (e Executer) Do() {
	main()
}
