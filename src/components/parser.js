const logger = require('./logger');
const { parserError } = require('./error-types');

class Parser {
  constructor(input) {
    this.input = input;
    this.index = 0;
  }

  parse() {
    try {
      return this.parseExpression();
    } catch (error) {
      if (error instanceof parserError) {
        throw error;
      } else {
        throw new parserError('Unexpected parsing error', error);
      }
    }
  }

  parseExpression() {
    const term = this.parseTerm();
    return this.parseExpressionPrime(term);
  }

  parseExpressionPrime(term) {
    if (this.index >= this.input.length) {
      return term;
    }

    const currentChar = this.input[this.index];
    if (currentChar === '+') {
      this.index++;
      const nextTerm = this.parseTerm();
      return { type: 'addition', left: term, right: nextTerm };
    } else if (currentChar === '-') {
      this.index++;
      const nextTerm = this.parseTerm();
      return { type: 'subtraction', left: term, right: nextTerm };
    } else {
      return term;
    }
  }

  parseTerm() {
    const factor = this.parseFactor();
    return this.parseTermPrime(factor);
  }

  parseTermPrime(factor) {
    if (this.index >= this.input.length) {
      return factor;
    }

    const currentChar = this.input[this.index];
    if (currentChar === '*') {
      this.index++;
      const nextFactor = this.parseFactor();
      return { type: 'multiplication', left: factor, right: nextFactor };
    } else if (currentChar === '/') {
      this.index++;
      const nextFactor = this.parseFactor();
      return { type: 'division', left: factor, right: nextFactor };
    } else {
      return factor;
    }
  }

  parseFactor() {
    if (this.index >= this.input.length) {
      throw new parserError('Unexpected end of input');
    }

    const currentChar = this.input[this.index];
    if (currentChar === '(') {
      this.index++;
      const expression = this.parseExpression();
      if (this.index >= this.input.length || this.input[this.index] !== ')') {
        throw new parserError('Missing closing parenthesis');
      }
      this.index++;
      return expression;
    } else if (!isNaN(parseInt(currentChar))) {
      let number = '';
      while (this.index < this.input.length && !isNaN(parseInt(this.input[this.index]))) {
        number += this.input[this.index];
        this.index++;
      }
      return { type: 'number', value: parseInt(number) };
    } else {
      throw new parserError(`Unexpected character '${currentChar}'`);
    }
  }
}

module.exports = Parser;