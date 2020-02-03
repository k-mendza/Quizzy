import { Component, OnInit } from '@angular/core';
import {QuestionService} from './question.service';
import {Question} from './question.model';

@Component({
  selector: 'app-question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.scss']
})
export class QuestionComponent implements OnInit {
  questions: Question[] = [];

  constructor(private questionService: QuestionService) { }

  ngOnInit() {
    this.getQuestions();
  }

  getQuestions() {
    this.questionService.getQuestions().subscribe(data => {
      for (const obj of data) {
        this.questions.push(new Question(obj.id, obj.text, obj.answers));
      }
    });
  }

  onSend() {
    this.questionService.verifyQuestions(this.questions);
  }
}
