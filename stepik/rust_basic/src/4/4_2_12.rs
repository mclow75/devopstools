use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str_a = String::new();
    let mut str_b = String::new();
    io::stdin().read_line(&mut str_a).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str_b).expect(ERR_INPUT_READ);
    let a: f32 = str_a.trim().parse().expect(ERR_INTEGER_READ);
    let b: f32 = str_b.trim().parse().expect(ERR_INTEGER_READ);
    println!("{0:.0} + ({1:.0}) = {2:.0}", a, b, a + b);
    println!("{0:.0} - ({1:.0}) = {2:.0}", a, b, a - b);
    println!("{0:.0} * ({1:.0}) = {2:.0}", a, b, a * b);
    println!("{0:.0} / ({1:.0}) = {2:.3}", a, b, a / b);
    println!("{0:.0} % ({1:.0}) = {2:.3}", a, b, a % b);
}
