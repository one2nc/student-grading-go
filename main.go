package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type Student struct {
	FirstName  string
	LastName   string
	University string
	Test1      int
	Test2      int
	Test3      int
	Test4      int
}

// Implementing Stringer interface for Student
func (s Student) String() string {
	return fmt.Sprintf("%s %s (%s): Test1=%d, Test2=%d, Test3=%d, Test4=%d", s.FirstName, s.LastName, s.University, s.Test1, s.Test2, s.Test3, s.Test4)
}

type StudentStat struct {
	Student
	FinalScore float32
	Grade      Grade
}

func parseCSV(filePath string) ([]Student, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var students []Student
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineNumber == 0 {
			lineNumber++
			continue // Skip header row
		}
		fields := strings.Split(line, ",")
		if len(fields) != 7 {
			return nil, fmt.Errorf("invalid CSV format at line %d", lineNumber)
		}

		test1, err := strconv.Atoi(fields[3])
		if err != nil {
			return nil, fmt.Errorf("invalid Test1 score at line %d", lineNumber)
		}
		test2, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, fmt.Errorf("invalid Test2 score at line %d", lineNumber)
		}
		test3, err := strconv.Atoi(fields[5])
		if err != nil {
			return nil, fmt.Errorf("invalid Test3 score at line %d", lineNumber)
		}
		test4, err := strconv.Atoi(fields[6])
		if err != nil {
			return nil, fmt.Errorf("invalid Test4 score at line %d", lineNumber)
		}

		students = append(students, Student{
			FirstName:  fields[0],
			LastName:   fields[1],
			University: fields[2],
			Test1:      test1,
			Test2:      test2,
			Test3:      test3,
			Test4:      test4,
		})
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return students, nil
}

func calculateGrade(students []Student) []StudentStat {
	var stats []StudentStat
	for _, s := range students {
		finalScore := float32(s.Test1+s.Test2+s.Test3+s.Test4) / 4
		var grade Grade
		switch {
		case finalScore < 35:
			grade = F
		case finalScore < 50:
			grade = C
		case finalScore < 70:
			grade = B
		default:
			grade = A
		}
		stats = append(stats, StudentStat{
			Student:    s,
			FinalScore: finalScore,
			Grade:      grade,
		})
	}
	return stats
}

func findOverallTopper(stats []StudentStat) StudentStat {
	var topper StudentStat
	for _, s := range stats {
		if s.FinalScore > topper.FinalScore {
			topper = s
		}
	}
	return topper
}

func findTopperPerUniversity(stats []StudentStat) map[string]StudentStat {
	groupedByUniversity := make(map[string][]StudentStat)
	for _, s := range stats {
		groupedByUniversity[s.University] = append(groupedByUniversity[s.University], s)
	}

	toppers := make(map[string]StudentStat)
	for university, students := range groupedByUniversity {
		toppers[university] = findOverallTopper(students)
	}
	return toppers
}

// Each function has been updated with comments, better error handling, and proper naming conventions following Go idioms.
