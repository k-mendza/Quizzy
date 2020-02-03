export class Answer {
  id: number;
  questionId: number;
  isSelected: boolean;
  text: string;


  constructor(id: number, questionId: number, isSelected: boolean, text: string) {
    this.id = id;
    this.questionId = questionId;
    this.isSelected = isSelected;
    this.text = text;
  }
}
