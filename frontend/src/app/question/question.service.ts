import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Question} from './question.model';

@Injectable()
export class QuestionService {
  baseUrl = 'http://localhost:8081';
  allQuestionsUrl = '/question/all';
  verifyQuestionsUrls = '/question/verify';

  constructor(private http: HttpClient) {
  }

  getQuestions() {
    return this.http.get<Question[]>(this.baseUrl + this.allQuestionsUrl);
  }

  verifyQuestions(questions: Question[]): void {
    this.http.post(this.baseUrl + this.verifyQuestionsUrls, questions);
  }
}
