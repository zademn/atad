use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

fn hurdleRace(k: i32, height: &[i32]) -> i32 {
    let res = height.iter().max().unwrap() - k;
    if res < 0 {
        0
    } else {
        res
    }
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    // let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let first_multiple_input: Vec<String> = stdin_iterator.next().unwrap().unwrap()
        .split(' ')
        .map(|s| s.to_string())
        .collect();

    let n = first_multiple_input[0].trim().parse::<i32>().unwrap();

    let k = first_multiple_input[1].trim().parse::<i32>().unwrap();

    let height: Vec<i32> = stdin_iterator.next().unwrap().unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<i32>().unwrap())
        .collect();

    let result = hurdleRace(k, &height);
    println!("{}", result);

    // writeln!(&mut fptr, "{}", result).ok();
}
