/*
* Напишите программу, которая считывает два вещественных числа (килограммы и фунты) с консоли и выводит для первого фунты, а для второго килограммы. Знаков после запятой достаточно 3.
* Sample Input:
2.0
4.0
Sample Output:
2 kg = 4.410 lbs
4 lbs = 1.816 kg
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str_a = String::new();
    let mut str_b = String::new();
    io::stdin().read_line(&mut str_a).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str_b).expect(ERR_INPUT_READ);
    let a: f64 = str_a.trim().parse().expect(ERR_INTEGER_READ);
    let b: f64 = str_b.trim().parse().expect(ERR_INTEGER_READ);
    println!("{0:.} kg = {1:.3} lbs", a, a * 2.205);
    println!("{0:.} lbs = {1:.3} kg", b, b * 0.454);
}
