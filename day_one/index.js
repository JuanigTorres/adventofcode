import { readFileSync } from 'fs';

const DIGITS = {
    'one': 1, 
    'two': 2,
    'three': 3, 
    'four': 4, 
    'five': 5, 
    'six': 6, 
    'seven': 7, 
    'eight': 8, 
    'nine': 9,
}

const tranform = (line) => {
    let min = line.length + 1;
    let first;
    Object
        .entries(DIGITS)
        .forEach(([ number ]) => {
            const ocurrence = line.search(number);
            if (ocurrence !== -1 && ocurrence < min) {
                min = ocurrence;
                first = number;
            } 
        });

    return line
        // Special cases
        .replaceAll(first, DIGITS[first]) 
        // Normal cases
        .replaceAll('one', 1)
        .replaceAll('two', 2)
        .replaceAll('three', 3)
        .replaceAll('four', 4)
        .replaceAll('five', 5)
        .replaceAll('six', 6)
        .replaceAll('seven', 7)
        .replaceAll('eight', 8)
        .replaceAll('nine', 9)
};

const onlyNumbers = (values) => [ ...values ].filter((l) => !isNaN(Number(l)));

const number = (digits) => {
    if (digits.length === 0) return 0;
    if (digits.length === 1) return Number(digits[0] + digits[0]);
    return Number(digits[0] + digits[digits.length - 1])
};

const total = (lines) => lines.reduce((prev, line) => {
    const transformed = tranform(line);
    const digits = onlyNumbers(transformed);
    const n = number(digits);
    return prev + n;
}, 0);

const calibration = total([
    'two1nine',
    'eightwothree',
    'abcone2threexyz',
    'xtwone3four',
    '4nineeightseven2',
    'zoneight234',
    '7pqrstsixteen',
]);

console.log({ calibration });

const stream = readFileSync('./day_one/input').toString(); 
const lines = stream.split('\n');
const result = total(lines);

console.log({ result });