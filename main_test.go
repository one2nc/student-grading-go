package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCSV(t *testing.T) {
	a := assert.New(t)
	students := parseCSV("grades.csv")

	a.Equal(30, len(students), "Student list size should be 30")

	fs := student{"Kaylen", "Johnson", "Duke University", 52, 47, 35, 38}
	a.Equal(fs, students[0], "First student should be Kaylen")

	ls := student{"Solomon", "Hunter", "Boston University", 45, 62, 32, 58}
	a.Equal(ls, students[29], "Last student should be Solomon")
}

func TestCalculateGrade(t *testing.T) {
	a := assert.New(t)

	gradedStudents := calculateGrade(parseCSV("grades.csv"))

	expScore := []float32{43, 59.25, 53, 58.25, 52.25, 50.75, 54.75, 49.25, 64.75, 43.25, 68.5, 57.75, 68.25, 66.75, 45.5, 45.75, 45.5, 58, 56, 60.25, 61, 62.5, 80.5, 53, 30.75, 57.5, 70.75, 48.5, 60.25, 49.25}
	expGrade := []Grade{C, B, B, B, B, B, B, C, B, C, B, B, B, B, C, C, C, B, B, B, B, B, A, B, F, B, A, C, B, C}

	a.Equal(30, len(gradedStudents), "All students should be graded")
	for i, ss := range gradedStudents {
		a.Equal(expScore[i], ss.finalScore, "Student %v expected score %v; but got %v", ss.firstName, expScore[i], ss.finalScore)
		a.Equal(expGrade[i], ss.grade, "Student %v expected grade %v; but got %v", ss.firstName, expGrade[i], ss.grade)
	}
}

func TestFindOverallTopper(t *testing.T) {
	gradedStudents := calculateGrade(parseCSV("grades.csv"))

	got := findOverallTopper(gradedStudents).student
	want := student{"Bernard", "Wilson", "Boston University", 90, 85, 76, 71}

	assert.Equal(t, got, want)
}

func TestFindTopperPerUniversity(t *testing.T) {
	a := assert.New(t)

	tpu := findTopperPerUniversity(calculateGrade(parseCSV("grades.csv")))

	bostonTopper := student{"Bernard", "Wilson", "Boston University", 90, 85, 76, 71}
	dukeTopper := student{"Tamara", "Webb", "Duke University", 73, 62, 90, 58}
	unionTopper := student{"Izayah", "Hunt", "Union College", 29, 78, 41, 85}
	calTopper := student{"Karina", "Shaw", "University of California", 69, 78, 56, 70}
	floTopper := student{"Nathan", "Gordon", "University of Florida", 53, 79, 84, 51}

	a.Equal(bostonTopper, tpu["Boston University"].student, "Boston University topper should be Bernard, but got %v", tpu["Boston University"].firstName)
	a.Equal(dukeTopper, tpu["Duke University"].student, "Duke University topper should be Tamara, but got %v", tpu["Duke University"].firstName)
	a.Equal(unionTopper, tpu["Union College"].student, "Union College topper should be Izayah, but got %v", tpu["Union College"].firstName)
	a.Equal(calTopper, tpu["University of California"].student, "University of California topper should be Karina, but got %v", tpu["University of California"].firstName)
	a.Equal(floTopper, tpu["University of Florida"].student, "University of Florida topper should be Nathan, but got %v", tpu["University of Florida"].firstName)
}
