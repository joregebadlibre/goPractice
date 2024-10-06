package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Active    bool
}

type Student struct {
	usuario User
	id      string
}

type Curso struct {
	titulo string
	video  []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	video1 := Video{titulo: "Video 1"}
	video2 := Video{titulo: "Video 2"}

	fmt.Println(video1)
	fmt.Println(video2)

	curso := Curso{titulo: "Curso de Go", video: []Video{video1, video2}}
	fmt.Println(curso)

	videos := []Video{video1, video2}

	curso.video = append(curso.video, videos...)

	video1.curso = curso

	fmt.Println("Los videos del curso: ", curso)
	fmt.Println("Los cursos del video1: ", video1.curso)

}
