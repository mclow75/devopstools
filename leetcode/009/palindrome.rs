pub fn is_palindrome(x: i32) -> bool {
    if x < 0 {
        return false;
    }

    let mut orig = x;
    let mut rev = 0;

    while orig != 0 {
        let last = orig % 10;
        orig = orig / 10;
        rev = rev * 10 + last;
    }
    return rev == x;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn palindrome_value() {
        assert_eq!(is_palindrome(121), true); // "Это палиндром"
    }
    #[test]
    fn negative() {
        assert_eq!(is_palindrome(-121), false); // "Отрицательное число не палиндром"
    }
    #[test]
    fn leading_zero() {
        assert_eq!(is_palindrome(10), false); // "Лидирующий ноль не палиндром"
    }
}
