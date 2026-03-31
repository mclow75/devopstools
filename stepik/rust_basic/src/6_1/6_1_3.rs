/*
* Напишите программу, которая считывает вещественное число и выводит (до 3 знаков) его дробную часть, используя оператор %.
Тестовые данные ✅
Sample Input:
17.1235
Sample Output:
0.123
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str = String::new();
    io::stdin().read_line(&mut str).expect(ERR_INPUT_READ);
    let num: f64 = str.trim().parse().expect(ERR_INTEGER_READ);
    println!("{:.3}", num.fract());
}
