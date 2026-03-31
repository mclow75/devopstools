/*
 * Вам необходимо создать программу, которая считывает вещественное число (масса космических тел в кг) и выводит массу объекта в научной записи (E). Это позволит вам анализировать массу космических объектов в удобной форме и с легкостью их интерпретировать.
 */
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str_a = String::new();
    io::stdin().read_line(&mut str_a).expect(ERR_INPUT_READ);
    let a: f64 = str_a.trim().parse().expect(ERR_INTEGER_READ);
    println!("{:1E}", a);
}
