import {Answer} from './answer/answer.model';

export class Question {
  id: number;
  text: string;
  answers: Answer[] = [];

  constructor(id: number = null, text: string = null, answers: Answer[] = null) {
    this.id = id;
    this.text = text;
    this.answers = answers;
  }
}
