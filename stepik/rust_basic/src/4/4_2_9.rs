/*
* Напишите программу, которая считывает вещественное и целое (usize) число , а затем выводит дробное число с точностью до числа знаков после запятой, указанного введённым целым числом.
* Sample Input:
5.123456
5
Sample Output:
5.12346
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str_a = String::new();
    let mut str_b = String::new();
    io::stdin().read_line(&mut str_a).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str_b).expect(ERR_INPUT_READ);
    let a: f32 = str_a.trim().parse().expect(ERR_INTEGER_READ);
    let prec: usize = str_b.trim().parse().expect(ERR_INTEGER_READ);
    println!("{:.prec$}", a);
}
