/*
 *
 */
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str_a = String::new();
    io::stdin().read_line(&mut str_a).expect(ERR_INPUT_READ);
    let a: f64 = str_a.trim().parse().expect(ERR_INTEGER_READ);
    println!("{}", a as i64);
    println!("{:.3}", a - (a as i64) as f64);
}
