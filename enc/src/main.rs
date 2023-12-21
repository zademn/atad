#![feature(isqrt)]
use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};
/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

fn encryption(s: &str) -> String {
    let s = s.replace(" ", "");
    let t = s.len().isqrt();
    let row_opt = [t, t, t + 1, t + 1];
    let col_opt = [t, t + 1, t, t + 1];
    let mut n_rows = 0;
    let mut n_cols = 0;
    for (r, c) in row_opt.iter().zip(col_opt.iter()) {
        if r * c >= s.len() {
            n_rows = *r;
            n_cols = *c;
            break;
        }
    }
    println!("{}", n_rows);
    println!("{}", n_cols);
    println!("{}", s[..8].to_owned());
    let mut v = vec![];
    for i in 0..n_rows {
        let row = if i == n_rows - 1 {
            s[i * n_cols..].as_bytes()
        } else {
            s[i * n_cols..(i + 1) * n_cols].as_bytes()
        };
        v.push(row);
    }

    let mut res = vec![];
    for i in 0..n_cols {
        for j in 0..n_rows {
            if i < v[j].len() {
                res.push(v[j][i]);
            }
        }
        res.push(' ' as u8);
    }
    String::from_utf8(res).unwrap()

}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    // let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let s = stdin_iterator.next().unwrap().unwrap();

    println!("{}", s);
    let result = encryption(&s);
    println!("{}", result)
    // writeln!(&mut fptr, "{}", result).ok();
}
