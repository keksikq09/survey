package main

import (
	"fmt"
)

type Question struct {
	Text    string
	Options []string
}

type Survey struct {
	ID        int
	Title     string
	Questions []Question
}

type Answer struct {
	QuestionID     int
	SelectedOption int
}

type SurveyManager struct {
	survey  Survey
	answers []Answer
}

func (sm *SurveyManager) StartSurvey() {
	for i, question := range sm.survey.Questions {
		fmt.Println(question.Text)
		for j, option := range question.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}
		var answer int
		fmt.Scanln(&answer)
		sm.answers = append(sm.answers, Answer{QuestionID: i, SelectedOption: answer - 1})
	}
}

func (sm *SurveyManager) GetResults() {
	fmt.Println("Survey Results:")
	for _, answer := range sm.answers {
		question := sm.survey.Questions[answer.QuestionID]
		fmt.Printf("Q: %s\nA: %s\n", question.Text, question.Options[answer.SelectedOption])
	}
}

func main() {
	questions := []Question{
		{"What is your favorite color?", []string{"Red", "Blue", "Green", "Yellow"}},
		{"What is your favorite animal?", []string{"Dog", "Cat", "Bird", "Fish"}},
	}
	survey := Survey{ID: 1, Title: "Simple Survey", Questions: questions}
	manager := SurveyManager{survey: survey}
	manager.StartSurvey()
	manager.GetResults()
}
