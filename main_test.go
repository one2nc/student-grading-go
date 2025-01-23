package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSV(t *testing.T) {
	a := assert.New(t)
	students, err := parseCSV("grades.csv")
	a.NoError(err, "parseCSV should not return an error")

	a.Equal(30, len(students), "Student list size should be 30")

	expectedFirst := Student{"Kaylen", "Johnson", "Duke University", 52, 47, 35, 38}
	a.Equal(expectedFirst, students[0], "First student should be Kaylen")

	expectedLast := Student{"Solomon", "Hunter", "Boston University", 45, 62, 32, 58}
	a.Equal(expectedLast, students[29], "Last student should be Solomon")
}

func TestCalculateGrade(t *testing.T) {
	a := assert.New(t)

	students, err := parseCSV("grades.csv")
	a.NoError(err, "parseCSV should not return an error")

	gradedStudents := calculateGrade(students)

	expScores := []float32{43, 59.25, 53, 58.25, 52.25, 50.75, 54.75, 49.25, 64.75, 43.25, 68.5, 57.75, 68.25, 66.75, 45.5, 45.75, 45.5, 58, 56, 60.25, 61, 62.5, 80.5, 53, 30.75, 57.5, 70.75, 48.5, 60.25, 49.25}
	expGrades := []Grade{C, B, B, B, B, B, B, C, B, C, B, B, B, B, C, C, C, B, B, B, B, B, A, B, F, B, A, C, B, C}

	a.Equal(30, len(gradedStudents), "All students should be graded")
	for i, ss := range gradedStudents {
		a.Equal(expScores[i], ss.FinalScore, "Student %v expected score %v; but got %v", ss.FirstName, expScores[i], ss.FinalScore)
		a.Equal(expGrades[i], ss.Grade, "Student %v expected grade %v; but got %v", ss.FirstName, expGrades[i], ss.Grade)
	}
}

func TestFindOverallTopper(t *testing.T) {
	students, err := parseCSV("grades.csv")
	assert.NoError(t, err, "parseCSV should not return an error")

	gradedStudents := calculateGrade(students)

	got := findOverallTopper(gradedStudents).Student
	want := Student{"Bernard", "Wilson", "Boston University", 90, 85, 76, 71}

	assert.Equal(t, got, want, "Overall topper should be Bernard Wilson")
}

func TestFindTopperPerUniversity(t *testing.T) {
	a := assert.New(t)

	students, err := parseCSV("grades.csv")
	a.NoError(err, "parseCSV should not return an error")

	gradedStudents := calculateGrade(students)

	toppers := findTopperPerUniversity(gradedStudents)

	expectedToppers := map[string]Student{
		"Boston University":        {"Bernard", "Wilson", "Boston University", 90, 85, 76, 71},
		"Duke University":          {"Tamara", "Webb", "Duke University", 73, 62, 90, 58},
		"Union College":            {"Izayah", "Hunt", "Union College", 29, 78, 41, 85},
		"University of California": {"Karina", "Shaw", "University of California", 69, 78, 56, 70},
		"University of Florida":    {"Nathan", "Gordon", "University of Florida", 53, 79, 84, 51},
	}

	for university, expected := range expectedToppers {
		a.Equal(expected, toppers[university].Student, "Topper for %s should be %v", university, expected)
	}
}
